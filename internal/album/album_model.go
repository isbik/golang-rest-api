package album

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	ID    primitive.ObjectID `json:"id" bson:"_id" validate:"required"`
	Title string             `json:"title" validate:"required"`
	Owner primitive.ObjectID `json:"owner", bson:"user,omitempty" validate:"required"`
}
