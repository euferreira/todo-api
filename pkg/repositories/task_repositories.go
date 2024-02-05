package repositories

import (
	"log"

	"github.com/euferreira/pkg/configurations"
	"github.com/euferreira/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	database *mongo.Client
}

func NewRepository() *Repository {
	return &Repository{
		database: configurations.Client,
	}
}

func (r *Repository) GetTasks() []entities.Task {
	ctx, cancel := getContext()
	defer cancel()

	log.Println("Getting tasks from database")

	database := r.database.Database("tasks")
	collection := database.Collection("tasks")
	cursor, _ := collection.Find(ctx, bson.D{})

	if cursor == nil {
		log.Println("Cursor is empty")
		return nil
	}

	var tasks []entities.Task
	for cursor.Next(ctx) {
		var task entities.Task
		err := cursor.Decode(&task)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func (r *Repository) CreateTask(task entities.Task) {
	ctx, cancel := getContext()
	defer cancel()

	r.database.Database("tasks").Collection("tasks").InsertOne(ctx, task)
}

func (r *Repository) UpdateTask(taskID string, updateTask entities.Task) {
	objID, _ := primitive.ObjectIDFromHex(taskID)
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":       updateTask.Title,
			"description": updateTask.Description,
			"isCompleted": updateTask.IsCompleted,
			"status":      updateTask.Status,
			"active":      updateTask.Active,
		},
	}

	ctx, cancel := getContext()

	_, err := r.database.Database("tasks").Collection("tasks").UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}

	defer cancel()
}

func (r *Repository) DeleteTask(id string) {
	ctx, cancel := getContext()
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := r.database.Database("tasks").Collection("tasks").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		panic(err)
	}
}

func (r *Repository) GetTask(id string) *entities.Task {
	ctx, cancel := getContext()
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	var task entities.Task

	err := r.database.Database("tasks").Collection("tasks").FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return &entities.Task{}
	}

	return &task
}
