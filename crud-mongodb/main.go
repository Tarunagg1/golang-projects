package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Data struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
}

type Data2 struct {
	ID    string `json:"id,omitempty" bson:"id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

var Mgr Manager

type Manager struct {
	// MongoDB connection and other fields can be added here
	Insert     func(interface{}) error
	GetAll     func() ([]Data, error)
	DeleteData func(primitive.ObjectID) error
	updateData    func(primitive.ObjectID) (Data, error)
}

func connectDb() {
	url := "mongodb://localhost:27017"

	client, err := mongo.Connect(options.Client().ApplyURI(url))

	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// err = client.connect(ctx)

	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)
		return
	}

	fmt.Println("Connected to MongoDB!")

	Mgr = &manager{
		connection: client,
		ctx:        ctx,
		cancel:     cancel,
	}
}

// func closeDb() {
// 	cancel.context.Cancel
// }

func init() {
	connectDb()
}

func main() {
	r := gin.Default()

	r.POST("/data", InsertData)

	r.GET("/data", getAll)

	r.DELETE("/data", deleteData)

	r.PUT("/data", UpdateData)

	r.Run(":9090")
}

func InsertData(c *gin.Context) {
	var data Data
	err := c.BindJSON(&data)

	if err != nil {
		fmt.Println("Error binding JSON:", err)
		return
	}

	Mgr.Insert(data)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}

func getAll(c *gin.Context) {
	data, err := Mgr.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func deleteData(c *gin.Context) {

}

func UpdateData(c *gin.Context) {
	var data Data2
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := primitive.ObjectIDFromHex(data.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	existingData, err := Mgr.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	existingData.Name = data.Name
	existingData.Email = data.Email

	err = Mgr.(id, existingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": existingData})
}
