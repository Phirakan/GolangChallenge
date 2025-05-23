package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"-"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}

type UserRegister struct {
	Name    		string `json:"name" binding:"required,min=2,max=100"`
	Email    		string `json:"email" binding:"required,email"`
	Password 		string `json:"password" binding:"required,min=6"`
	
}

type UserResponse struct {
	ID        bson.ObjectID `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	CreatedAt time.Time     `json:"created_at"`
}

type UserLogin struct {
	Email    	string `json:"email" binding:"required,email"`
	Password 	string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Name  string `json:"name,omitempty" binding:"omitempty,min=2,max=100"`
	Email string `json:"email,omitempty" binding:"omitempty,email"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	
	}
}