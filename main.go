package main

import (

	// "errors"
	// "go/parser"

	"log"
	"os"
	"totality-assignment/mod/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// errors "github.com/haproxytech/config-parser/v4/errors"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	controllers.ConnectToDB()
	router := gin.Default()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	router.POST("/signup", controllers.SignupUser)
	router.POST("/login", controllers.LoginUser)
	router.POST("/notes", controllers.CreateNote)
	router.GET("/notes", controllers.GetNotes)
	router.DELETE("/notes", controllers.DeleteNote)

	router.Run("localhost:8080")
}

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}
