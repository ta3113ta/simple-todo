package todoCRUD

import (
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id        string `json:"id"`
	Title     string `json:"title" binding:"required"`
	Note      string `json:"note"`
	Completed bool   `json:"completed"`
}

type TodoUpdate struct {
	Title     string `json:"title"`
	Note      string `json:"note"`
	Completed bool   `json:"completed"`
}

func CreateTodo(todo *Todo) {
	client, ctx, cancel := GetConnection()
	defer cancel()

	collection := client.Database("projects").Collection("todos")

	todo.Id = uuid.New().String()

	if _, err := collection.InsertOne(ctx, todo); err != nil {
		log.Printf("Can not insert data, Error: %v", err)
	}
}

func FindAllTodo() []primitive.M {
	client, ctx, cancel := GetConnection()
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
	client, ctx, cancel := GetConnection()
	defer cancel()

	filter := bson.M{"id": bson.M{"$eq": id}}

	var todo primitive.M
	collection := client.Database("projects").Collection("todos")
	err := collection.FindOne(ctx, filter).Decode(&todo)

	if err != nil {
		log.Fatal(err)
	}

	return todo
}

func UpdateTodo(id string, updateTodoDto TodoUpdate) {
	client, ctx, cancel := GetConnection()
	defer cancel()

	collection := client.Database("projects").Collection("todos")
	filter := bson.M{"id": bson.M{"$eq": id}}
	todo := bson.D{primitive.E{Key: "$set", Value: updateTodoDto}}

	result := collection.FindOneAndUpdate(ctx, filter, todo)
	if err := result.Err(); err != nil {
		log.Println(err)
	}
}

func DeleteTodo(id string) {
	client, ctx, cancel := GetConnection()
	defer cancel()

	collection := client.Database("projects").Collection("todos")
	filter := bson.M{"id": bson.M{"$eq": id}}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
}
