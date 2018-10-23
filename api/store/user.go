package store

import (
	"wloggr/api/models"
)

// NewUser ...
func (p *DB) NewUser(u *models.User) error {
	q := `INSERT INTO users
		  (userId, email, hash, fullName)
		  VALUES (?, ?, ?, ?)`
	_, err := p.db.Exec(q, u.UserID, u.Email, u.Hash, u.FullName)
	return err
}

// UpdateUserFullName ...
func (p *DB) UpdateUserFullName(userID, fullName string) error {
	q := `UPDATE users SET fullName = (?)
		  WHERE userId = (?)`
	_, err := p.db.Exec(q, fullName, userID)
	return err
}

// GetUserByID ...
func (p *DB) GetUserByID(userID string) (*models.User, error) {
	q := `SELECT userId, email, hash, fullName
		  FROM users WHERE userId = (?)`
	u := models.User{}
	err := p.db.Get(&u, q, userID)
	return &u, err
}

// GetUserByEmail ...
func (p *DB) GetUserByEmail(email string) (*models.User, error) {
	q := `SELECT userId, email, hash, fullName
		  FROM users WHERE email = (?)`
	u := models.User{}
	err := p.db.Get(&u, q, email)
	return &u, err
}

// GetUsers ...
func (p *DB) GetUsers() ([]*models.User, error) {
	q := `SELECT userId, email, hash, fullName
		  FROM users`
	rows, err := p.db.Queryx(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
