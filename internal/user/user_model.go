package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Login     string             `json:"login" validate:"required,min=2,max=100"`
	Email     string             `json:"email" validate:"required,email"`
	Password  string             `json:"password"`
	CreatedAt string             `json:"created_at"`
}
