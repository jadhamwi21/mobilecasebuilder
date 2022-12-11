package entity

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int    `json:"id" bson:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Role      string `json:"role" bson:"role"`
	Orders    []int  `json:"orders" bson:"orders"`
}

func (user *User) Creating(ctx context.Context) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
