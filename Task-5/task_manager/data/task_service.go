package data

import (
	"context"
	"log"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection

func InitMongoDB(uri string, dbName string, collectionName string) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("MongoDB connection error:", err)
    }

    taskCollection = client.Database(dbName).Collection(collectionName)
}

// CRUD Functions

func CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
    return taskCollection.InsertOne(context.TODO(), task)
}

func GetAllTasks() ([]models.Task, error) {
    var tasks []models.Task
    cursor, err := taskCollection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var task models.Task
        if err := cursor.Decode(&task); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

func GetTaskByID(id string) (*models.Task, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var task models.Task
    err = taskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
    return &task, err
}

func UpdateTask(id string, task models.Task) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = taskCollection.UpdateOne(
        context.TODO(),
        bson.M{"_id": objID},
        bson.M{"$set": task},
    )
    return err
}

func DeleteTask(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = taskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    return err
}
