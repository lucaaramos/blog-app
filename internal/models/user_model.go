//put package

package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	UserName  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	Role      string    `json:"role" bson:"role"`
	Posts     []Post    `json:"posts"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"update_at" bson:"update_at"`
	Status    bool      `json:"status" bson:"status"`
}

func NewUser(name, email, username, password, role string, posts []Post) *User {
	return &User{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     email,
		UserName:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      role,
		Status:    true,
	}
}
