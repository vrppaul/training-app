package crud

import "go.mongodb.org/mongo-driver/mongo"

func GetCRUDDB(dbName string, client *mongo.Client) *CRUDDB {
	database := client.Database(dbName)
	return &CRUDDB{
		ExercisesCollection: database.Collection("exercises"),
	}
}

type CRUDDB struct {
	ExercisesCollection *mongo.Collection
}
