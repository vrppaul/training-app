package main

import "os"

type Config struct {
	Port        string
	MongoUri    string
	MongoDbName string
}

const defaultPort = "8080"
const defaultMongoDbUri = "mongodb://localhost"
const defaultMongoDbName = "training"

func GetConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		mongoUri = defaultMongoDbUri
	}

	mongoDbName := os.Getenv("MONGO_DB_NAME")
	if mongoDbName == "" {
		mongoDbName = defaultMongoDbName
	}

	return &Config{
		Port:        port,
		MongoUri:    mongoUri,
		MongoDbName: mongoDbName,
	}
}
