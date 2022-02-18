package main

import (
	"log"
	"main/internal/auth"
	"main/internal/photo"
	"main/internal/user"
	"main/pkg/config"
	"main/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.AutomaticEnv()

	err := config.ReadConf()

	if err != nil {
		log.Fatalf("Read config: %v", err)
	}

	database.MongoDB()

	gin.SetMode("release")

	router := gin.Default()

	router.Use(gin.Logger())

	auth.AuthRoutes(router)
	user.UserRoutes(router)
	photo.PhotoRoutes(router)

	router.Run(":" + viper.GetString("server.port"))

}
