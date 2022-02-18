package photo

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"main/internal/album"
	photo "main/internal/photo/dto"
	"main/pkg/database"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindUserPhotos(userId string, page int64, limit int64, photos *[]Photo) error {
	var photoCollection = database.MI.DB.Collection("photo")

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	options := options.Find()
	options.SetSkip(page * limit)
	options.SetLimit(limit)

	cursor, err := photoCollection.Find(context.Background(), bson.M{"owner": id}, options)
	if err != nil {
		return err
	}

	if err = cursor.All(context.TODO(), photos); err != nil {
		return err
	}

	return nil

}

func DeletePhotoById(id string) error {
	var photoCollection *mongo.Collection = database.MI.DB.Collection("photo")

	photoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = photoCollection.DeleteOne(context.Background(), bson.M{"_id": photoId})
	if err != nil {
		return err
	}

	return nil
}

func FindPhotoById(id string, photo *Photo) error {
	var photoCollection *mongo.Collection = database.MI.DB.Collection("photo")

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = photoCollection.FindOne(context.Background(), bson.M{"_id": userId}).Decode(photo)
	if err != nil {
		return err
	}

	return nil
}

func LoadPhotosFromJsonPlaceholder() ([]photo.PhotoLoadDto, error) {

	var photos []photo.PhotoLoadDto

	response, err := http.Get("http://jsonplaceholder.typicode.com/photos")
	if err != nil {
		return photos, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return photos, err
	}

	err = json.Unmarshal(body, &photos)
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func InsertPhotos(userId string, photos []photo.PhotoLoadDto) error {
	var photoCollection *mongo.Collection = database.MI.DB.Collection("photo")

	ownerId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	albumsMap, err := GetAlbumsFromPhotos(ownerId, photos)
	if err != nil {
		return err
	}

	var photoDocs []interface{}

	for _, photo := range photos {
		photoDocs = append(photoDocs, Photo{
			ID:           primitive.NewObjectID(),
			Title:        photo.Title,
			Url:          photo.Url,
			ThumbnailUrl: photo.ThumbnailUrl,
			Owner:        ownerId,
			Album:        albumsMap[photo.AlbumId].ID,
		})
	}

	var albums []album.Album

	for _, album := range albumsMap {
		albums = append(albums, album)
	}

	err = album.InsertAlbums(userId, &albums)
	if err != nil {
		return err
	}

	// TODO add transaction or albums will be empty
	_, err = photoCollection.InsertMany(context.Background(), photoDocs)
	if err != nil {
		return err
	}

	return nil

}

func GetAlbumsFromPhotos(ownerId primitive.ObjectID, photos []photo.PhotoLoadDto) (map[int]album.Album, error) {
	uniqueAlbums := make(map[int]album.Album)

	for _, photo := range photos {
		if _, ok := uniqueAlbums[photo.AlbumId]; !ok {
			uniqueAlbums[photo.AlbumId] = album.Album{
				Owner: ownerId,
				Title: string(rune(photo.AlbumId)),
				ID:    primitive.NewObjectID(),
			}
		}
	}

	return uniqueAlbums, nil
}
