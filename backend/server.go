package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/vrppaul/training-app/db"
	"github.com/vrppaul/training-app/db/crud"
	"github.com/vrppaul/training-app/graph"
	"github.com/vrppaul/training-app/graph/generated"
)

func main() {
	config := GetConfig()

	client, disconnect := db.GetClient(config.MongoDbUri, config.MongoUsername, config.MongoPassword)
	defer disconnect()
	database := crud.GetCRUDDB(config.MongoDbName, client)

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://0.0.0.0", "http://localhost"},
		AllowCredentials: true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{CRUDDB: database}}))

	router.Route("/api", func(router chi.Router) {
		router.Handle("/playground", playground.Handler("GraphQL playground", "/api"))
		router.Handle("/", srv)
	})

	uri := config.Host + ":" + config.Port
	log.Printf("application is served ad %s", uri)
	log.Fatal(http.ListenAndServe(uri, router))
}
