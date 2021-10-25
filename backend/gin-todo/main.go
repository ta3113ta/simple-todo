package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	Title     string `json:"title" binding:"required"`
	Note      string `json:"note"`
	Completed bool   `json:"completed"`
}

type TodoUpdate struct {
	Title     string `json:"title"`
	Note      string `json:"note"`
	Completed bool   `json:"completed"`
}

const (
	connectTimeout = 5
	connectionURI  = "mongodb://localhost:2717/projects"
)

func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))

	if err != nil {
		log.Fatalf("Failed to create client : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	if err = client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	return client, ctx, cancel
}

func CreateTodo(todo *Todo) {
	client, ctx, cancel := getConnection()
	defer cancel()

	collection := client.Database("projects").Collection("todos")

	if _, err := collection.InsertOne(ctx, todo); err != nil {
		log.Printf("Can not insert data, Error: %v", err)
	}
}

func FindAllTodo() []primitive.M {
	client, ctx, cancel := getConnection()
	defer cancel()

	collection := client.Database("projects").Collection("todos")
	curr, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var todos []primitive.M

	for curr.Next(ctx) {
		var todo bson.M
		err := curr.Decode(&todo)

		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	if err := curr.Err(); err != nil {
		log.Fatal(err)
	}

	return todos
}

func FindOneById(id string) primitive.M {
	client, ctx, cancel := getConnection()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	var todo primitive.M
	collection := client.Database("projects").Collection("todos")
	err = collection.FindOne(ctx, filter).Decode(&todo)

	if err != nil {
		log.Fatal(err)
	}

	return todo
}

func UpdateTodo(id string, updateTodoDto TodoUpdate) {
	client, ctx, cancel := getConnection()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("projects").Collection("todos")
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	todo := bson.D{primitive.E{Key: "$set", Value: updateTodoDto}}
	
	result := collection.FindOneAndUpdate(ctx, filter, todo)
	if err := result.Err(); err != nil {
		log.Println(err)
	}
}

func DeleteTodo(id string) {
	client, ctx, cancel := getConnection()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("projects").Collection("todos")
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
}

func findAll(c *gin.Context) {
	todos := FindAllTodo()
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": todos})
}

func findOne(c *gin.Context) {
	id := c.Param("id")
	todo := FindOneById(id)

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": todo})
}

func create(c *gin.Context) {
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CreateTodo(&todo)
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "created"})
}

func update(c *gin.Context) {
	var todoUpdate TodoUpdate
	id := c.Param("id")

	if err := c.ShouldBindJSON(&todoUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateTodo(id, todoUpdate)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "updated"})
}

func delete(c *gin.Context) {
	id := c.Param("id")
	DeleteTodo(id)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "deleted"})
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())

	router.GET("/todo", findAll)
	router.GET("/todo/:id", findOne)
	router.POST("/todo", create)
	router.PATCH("/todo/:id", update)
	router.DELETE("/todo/:id", delete)

	router.Run(":5000")
}
