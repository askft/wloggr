package models

type User struct {
	UserID   string `json:"userId"   db:"userId"`
	Email    string `json:"email"    db:"email"`
	Hash     string `json:"hash"     db:"hash"`
	FullName string `json:"fullName" db:"fullName"`
	// CreatedDate string `json:"createdDate"`
	// UpdatedDate string `json:"updatedDate"`
}

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
