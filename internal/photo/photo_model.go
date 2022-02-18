package photo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Photo struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title" validate:"required"`
	Url          string             `json:"url" bson:"url" validate:"required"`
	ThumbnailUrl string             `json:"thumbnailUrl" bson:"thumbnailUrl" validate:"required"`
	Owner        primitive.ObjectID `json:"owner", bson:"owner"`
	Album        primitive.ObjectID `json:"album" bson:"album"`
}
