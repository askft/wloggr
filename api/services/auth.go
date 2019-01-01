/*
	services.go

	Signin, Signup, GetToken (JWT)
*/

package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/askft/wloggr/api/models"
	"github.com/askft/wloggr/api/store"
	"github.com/askft/wloggr/api/util"

	validator "github.com/asaskevich/govalidator"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type StatusError struct {
	Code int
	Msg  string
	error
}

var errNull = errors.New("")

// Signin ...
func Signin(c *models.Signin) (*models.User, *StatusError) {
	u, err := store.Store.GetUserByEmail(strings.ToLower(c.Email))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &StatusError{
				http.StatusNotFound,
				"user does not exist",
				err,
			}
		}
		fmt.Println(err.Error())
		return nil, &StatusError{
			http.StatusInternalServerError,
			err.Error(),
			errNull,
		}
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(u.Hash),
		[]byte(c.Password))
	if err != nil {
		return nil, &StatusError{
			http.StatusUnauthorized,
			"incorrect password",
			err,
		}
	}
	return u, nil
}

// Signup ...
func Signup(c *models.Signup) *StatusError {
	if !validator.IsEmail(c.Email) {
		return &StatusError{
			http.StatusBadRequest,
			"invalid email address (bad format)",
			errNull,
		}
	}
	if !validator.IsASCII(c.Password) {
		return &StatusError{
			http.StatusBadRequest,
			"invalid password (non-ascii)",
			errNull,
		}
	}
	if !validator.IsByteLength(c.Password, 8, 128) {
		return &StatusError{
			http.StatusBadRequest,
			"invalid password (must be between 8 and 128 bytes)",
			errNull,
		}
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(c.Password),
		bcrypt.DefaultCost)
	if err != nil {
		return &StatusError{
			http.StatusInternalServerError,
			"could not generate password",
			err,
		}
	}
	fmt.Printf("Hash of %s: %s\n", c.Password, hash)

	u, err := uuid.NewV4()
	if err != nil {
		return &StatusError{
			http.StatusInternalServerError,
			"could not generate UUID",
			err,
		}
	}
	userID := uuid.Must(u, err).String()

	err = store.Store.NewUser(&models.User{
		UserID: userID,
		Email:  strings.ToLower(c.Email),
		Hash:   string(hash),
	})
	if err != nil {
		return &StatusError{
			http.StatusInternalServerError,
			err.Error(), // TODO
			err,
		}
	}
	return nil
}

/*
	JWT specific services ---------------------------------
*/

// GetToken ...
func GetToken(userID string) (string, *StatusError) {
	claims := jwt.StandardClaims{
		Subject:   strings.ToLower(userID),
		Issuer:    util.Config.DomainName,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signature, err := token.SignedString([]byte(util.Config.JwtSigningKey))
	if err != nil {
		log.Printf("err: %+v\n", err)
		return "", &StatusError{
			http.StatusInternalServerError,
			err.Error(),
			err,
		}
	}
	log.Printf("Issued token: %v\n", signature)
	return signature, nil
}
