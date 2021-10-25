package todoCRUD

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CreateTodo(todo *Todo) {
	client, ctx, cancel := GetConnection()
	defer cancel()

	collection := client.Database("projects").Collection("todos")

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
	client, ctx, cancel := GetConnection()
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
	client, ctx, cancel := GetConnection()
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
