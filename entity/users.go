package entity

import (
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Number    string     `json:"number"`
	Roles     string     `json:"roles"`
	Saldo     int64      `json:"saldo"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// Admin New User
func NewUser(name, email, number, password, roles string, saldo int64) *User {
	return &User{
		Name:      name,
		Email:     email,
		Number:    number,
		Roles:     roles,
		Saldo:     saldo,
		Password:  password,
		CreatedAt: time.Now(),
	}
}

// Admin Update User
func UpdateUser(id int64, name, email, number, roles, password string, saldo int64) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Number:    number,
		Roles:     roles,
		Password:  password,
		Saldo:     saldo,
		UpdatedAt: time.Now(),
	}
}

// Public Register
func Register(email, password, roles, number string) *User {
	return &User{
		Email:    email,
		Password: password,
		Roles:    roles,
		Number:   number,
	}
}

// User update by self
func UpdateUserSelf(id int64, name, email, number, password, roles string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Number:    number,
		Password:  password,
		Roles:     roles,
		UpdatedAt: time.Now(),
	}
}
