package main

import "os"

type Config struct {
	Host          string
	Port          string
	MongoDbUri    string
	MongoDbName   string
	MongoUsername string
	MongoPassword string
}

const defaultHost = "0.0.0.0"
const defaultPort = "8080"
const defaultMongoDbHost = "localhost"
const defaultMongoDbPort = "27017"
const defaultMongoDbName = "training"

func GetConfig() *Config {
	host := os.Getenv("HOST")
	if host == "" {
		host = defaultHost
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoDbHost := os.Getenv("MONGO_HOST")
	mongoDbPort := os.Getenv("MONGO_PORT")
	if mongoDbHost == "" {
		mongoDbHost = defaultMongoDbHost
	}
	if mongoDbPort == "" {
		mongoDbPort = defaultMongoDbPort
	}
	mongoDbUri := "mongodb://" + mongoDbHost + mongoDbPort

	mongoDbName := os.Getenv("MONGO_DB_NAME")
	if mongoDbName == "" {
		mongoDbName = defaultMongoDbName
	}

	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	return &Config{
		Host:          host,
		Port:          port,
		MongoDbUri:    mongoDbUri,
		MongoDbName:   mongoDbName,
		MongoUsername: mongoUsername,
		MongoPassword: mongoPassword,
	}
}
