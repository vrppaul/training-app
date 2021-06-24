package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/vrppaul/training-app/db/crud"
	"github.com/vrppaul/training-app/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CRUDDB    *crud.CRUDDB
	exercises []*model.Exercise
}
