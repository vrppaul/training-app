package crud

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vrppaul/training-app/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (cd *CRUDDB) GetExercises() ([]*model.Exercise, error) {
	var exercises []*model.Exercise

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := cd.ExercisesCollection.Find(ctx, bson.D{})
	if err != nil {
		return exercises, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var exercise *model.Exercise
		err := cur.Decode(&exercise)
		if err != nil {
			return exercises, err
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

func (cd *CRUDDB) InsertExercises() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := cd.ExercisesCollection.InsertOne(ctx, bson.D{{"name", "some"}, {"description", "exercise"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
