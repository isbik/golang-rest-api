package album

import (
	"context"
	"main/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateAlbumById(id string, album *Album) error {
	var albumCollection *mongo.Collection = database.MI.DB.Collection("album")

	albumId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = albumCollection.UpdateOne(context.Background(), bson.M{"_id": albumId}, bson.M{
		"$set": bson.D{{
			"title", album.Title,
		}}},
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAlbumById(id string) error {
	var albumCollection *mongo.Collection = database.MI.DB.Collection("album")

	albumId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = albumCollection.DeleteOne(context.Background(), bson.M{"_id": albumId})

	if err != nil {
		return err
	}

	return nil
}

func InsertAlbums(id string, albums *[]Album) error {
	var albumCollection *mongo.Collection = database.MI.DB.Collection("album")

	var docs []interface{}

	for _, album := range *albums {
		docs = append(docs, album)
	}

	_, err := albumCollection.InsertMany(context.Background(), docs)

	if err != nil {
		return err
	}

	return nil
}
