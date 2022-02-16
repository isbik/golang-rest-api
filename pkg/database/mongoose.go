package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DatabaseName string
	Host         string
	Port         string
	Username     string
	Password     string
	Ssl          string
}

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func MongoDB() {

	var cfg = Config{
		Username: viper.GetString("database.user"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Ssl:      viper.GetString("database.ssl"),
	}

	ctx := context.Background()

	clientUri := "mongodb://" + cfg.Host + ":" + cfg.Port + "/?ssl=" + cfg.Ssl + "&w=majority"
	fmt.Println(clientUri)
	clientOptions := options.Client().ApplyURI(clientUri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(viper.GetString("database.name")),
	}
}
