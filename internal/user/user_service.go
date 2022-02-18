package user

import (
	"context"
	"main/pkg/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *User) error {
	var userCollection *mongo.Collection = database.MI.DB.Collection("user")

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().Format(time.RFC3339)

	_, err := userCollection.InsertOne(context.Background(), user)

	if err != nil {
		return err
	}

	return nil
}

func CountUsersByEmail(email string) (int64, error) {
	var userCollection *mongo.Collection = database.MI.DB.Collection("user")

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": email})
	defer cancel()

	if err != nil {
		return 0, err
	}

	return count, nil
}

func FindUserById(id string, user *User) error {
	var userCollection *mongo.Collection = database.MI.DB.Collection("user")

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	err = userCollection.FindOne(context.Background(), bson.M{"_id": userId}).Decode(user)

	if err != nil {
		return err
	}

	return nil
}

func FindUserByEmail(email string, user *User) error {
	var userCollection *mongo.Collection = database.MI.DB.Collection("user")

	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(user)

	if err != nil {
		return err
	}

	return nil
}

func FindAllUsers() ([]User, error) {
	var userCollection *mongo.Collection = database.MI.DB.Collection("user")

	var users []User

	cursor, err := userCollection.Find(context.Background(), bson.D{{}})

	if err != nil {
		return users, err
	}

	if err = cursor.All(context.TODO(), &users); err != nil {
		return users, err
	}

	return users, nil
}
