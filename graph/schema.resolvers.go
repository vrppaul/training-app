package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vrppaul/training-app/graph/generated"
	"github.com/vrppaul/training-app/graph/model"
)

func (r *mutationResolver) CreateExercise(ctx context.Context, input model.NewExercise) (*model.Exercise, error) {
	return r.CRUDDB.InsertExercise(&input)
}

func (r *queryResolver) Exercise(ctx context.Context, id string) (*model.Exercise, error) {
	return r.CRUDDB.GetExerciseById(id)
}

func (r *queryResolver) Exercises(ctx context.Context) ([]*model.Exercise, error) {
	return r.CRUDDB.GetExercises()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
