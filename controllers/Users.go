package controllers

import (
	"context"
	"encoding/json"

	"log"
	"net/http"
	"totality-assignment/mod/db/entities"

	// entities "gitlab.com/shiiivaaansh/totality-corp/db"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	// "go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var SECRET_KEY = []byte("gosecretkey")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

func ConnectToDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://shivansh:shivansh@cluster0.eg46rjw.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil

}

func GetCollection(DbName string, CollectionName string) (*mongo.Collection, error) {

	client, _ := ConnectToDB()

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}

func SignupUser(c *gin.Context) {
	collection, err := GetCollection("accuknox", "users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	var user = new(entities.User)
	if err := c.BindJSON(user); err != nil {

		c.JSON(http.StatusOK, "ERRORRR from JSON!!!!!!")
		return
	}
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	// response, _ := json.Marshal(res)
	c.JSON(http.StatusOK, nil)
}

func LoginUser(c *gin.Context) {

	var user = new(entities.UserLogin)
	if err := c.BindJSON(user); err != nil {

		c.JSON(http.StatusOK, "ERRORRR from JSON!!!!!!")
		return
	}
	collection, err := GetCollection("accuknox", "users")
	if err != nil {
		c.JSON(http.StatusBadRequest, "error1")
		return
	}
	err = collection.FindOne(context.Background(), bson.M{"email": user.Email, "password": user.Password}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error2")
		return
	}
	jwtToken, err := GenerateJWT()
	if err != nil {
		c.JSON(http.StatusBadRequest, "error3")
		return
	}
	// filter := bson.D{{"email": user.Email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "sid", Value: jwtToken}}}}
	_, err = collection.UpdateOne(context.Background(), bson.M{"email": user.Email}, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error3")
		return
	}
	c.JSON(http.StatusOK, jwtToken)

}

func CreateNote(c *gin.Context) {
	collection, err := GetCollection("accuknox", "notes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	var note = new(entities.Note)
	if err := c.BindJSON(note); err != nil {

		c.JSON(http.StatusOK, "ERRORRR from JSON!!!!!!")
		return
	}
	res, err := collection.InsertOne(context.Background(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	response, _ := json.Marshal(res)
	c.JSON(http.StatusOK, response)
}

func GetNotes(c *gin.Context) {
	collection, err := GetCollection("accuknox", "notes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	var notes = new(entities.SearchNote)
	if err := c.BindJSON(notes); err != nil {

		c.JSON(http.StatusOK, "ERRORRR from JSON!!!!!!")
		return
	}
	var note = new(entities.Note)
	err = collection.FindOne(context.Background(), bson.M{"sid": notes.SID}).Decode(&note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error2")
		return
	}
	c.JSON(http.StatusOK, note)

}

func DeleteNote(c *gin.Context) {
	collection, err := GetCollection("accuknox", "notes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	var del = new(entities.DeleteNote)
	if err := c.BindJSON(del); err != nil {

		c.JSON(http.StatusOK, "ERRORRR from JSON!!!!!!")
		return
	}

	_, err = collection.DeleteMany(context.TODO(), bson.M{"sid": del.SID})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, "Deleted")
}
