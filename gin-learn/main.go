package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type api struct {
	ID    primitive.ObjectID `bson:"_id, omitempty"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}

type manager struct {
	Connection *mongo.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
}

// var data api

var DbMgr manager

type Manager interface {
	Insert(interface{}) error
	GetAll() ([]api, error)
	DeleteOne(primitive.ObjectID) error
	UpdateData(api) error
}

var Mgr Manager

func connectDB() {
	url := "localhost:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", url)))

	if err != nil {
		fmt.Println("Error while connnection: ", err)
		return
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	Mgr = &manager{
		Connection: client,
		Ctx:        ctx,
		Cancel:     cancel,
	}

}

func CloseDB(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func (mgr *manager) Insert(data interface{}) error {
	orgCollection := mgr.Connection.Database("gouser").Collection("user")
	result, err := orgCollection.InsertOne(context.TODO(), data)
	fmt.Println(result.InsertedID)
	return err
}

func (mgr *manager) DeleteOne(id primitive.ObjectID) error {
	collection := mgr.Connection.Database("gouser").Collection("user")
	filter := primitive.D{{Key: "_id", Value: id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}

func (mgr *manager) GetAll() ([]api, error) {
	var users []api
	collection := mgr.Connection.Database("gouser").Collection("user")
	cursor, err := collection.Find(context.TODO(), primitive.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user api
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (mgr *manager) UpdateData(user api) error {
	collection := mgr.Connection.Database("gouser").Collection("user")
	filter := primitive.D{{Key: "_id", Value: user.ID}}
	update := primitive.D{{Key: "$set", Value: user}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func main() {
	connectDB()
	r := gin.Default()

	r.GET("/todos", GetTodoList)
	r.GET("/user", GetUser)
	r.POST("/user", CreateUser)

	r.Run()
}

func GetUser(c *gin.Context) {
	data, err := Mgr.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"userData": data,
	})
}

func CreateUser(c *gin.Context) {
	var u api
	err := c.BindJSON(&u)

	Mgr.Insert(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data inserted",
	})
}

func GetTodoList(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
		return
	}

	var target []map[string]interface{}

	err = json.Unmarshal(data, &target)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, target)
}
