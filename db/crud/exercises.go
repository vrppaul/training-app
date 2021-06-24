package crud

import (
	"context"
	"time"

	"github.com/vrppaul/training-app/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (cd *CRUDDB) GetExerciseById(ID string) (*model.Exercise, error) {
	var exercise *model.Exercise
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return exercise, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := cd.ExercisesCollection.FindOne(ctx, bson.M{"_id": objectId})
	res.Decode(&exercise)
	return exercise, nil
}

func (cd *CRUDDB) InsertExercise(input *model.NewExercise) (*model.Exercise, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := cd.ExercisesCollection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}
	return &model.Exercise{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		Name:        input.Name,
		Description: input.Description,
	}, nil
}
