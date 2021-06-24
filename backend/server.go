package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vrppaul/training-app/db"
	"github.com/vrppaul/training-app/db/crud"
	"github.com/vrppaul/training-app/graph"
	"github.com/vrppaul/training-app/graph/generated"
)

func main() {
	config := GetConfig()

	client, disconnect := db.GetClient(config.MongoUri)
	defer disconnect()
	database := crud.GetCRUDDB(config.MongoDbName, client)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{CRUDDB: database}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
