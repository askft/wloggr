package route

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/askft/wloggr/api/util"

	"github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
)

var (
	errInvalidToken = errors.New("invalid token")
)

// VerifyToken is middleware that requires authentication through a
// JWT (JSON Web Token). Binds the JWT claims subject to the `ckUserID`
// context key. The user ID can be retrieved inside a handler function
// as, with `r` as an `http.Request`, `userID := ctxString(r, ckUserID)`.
func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, err := verify(w, r)
			if err != nil {
				log.Println(err)
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}

func verify(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	signingKeyFn := func(token *jwt.Token) (interface{}, error) {
		return []byte(util.Config.JwtSigningKey), nil
	}
	var claims jwt.StandardClaims

	token, err := jwtreq.ParseFromRequest(
		r,
		jwtreq.AuthorizationHeaderExtractor,
		signingKeyFn,
		jwtreq.WithClaims(&claims))

	// ErrNoTokenInRequest is the only error returned from ParseFromRequest.
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	if !token.Valid {
		http.Error(w, errInvalidToken.Error(), http.StatusUnauthorized)
		return nil, errInvalidToken
	}
	return context.WithValue(r.Context(), ckUserID, claims.Subject), nil
}
