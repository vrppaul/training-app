package main

import "os"

type Config struct {
	Host          string
	Port          string
	MongoUri      string
	MongoDbName   string
	MongoUsername string
	MongoPassword string
}

const defaultHost = "0.0.0.0"
const defaultPort = "8080"
const defaultMongoDbUri = "mongodb://localhost"
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

	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		mongoUri = defaultMongoDbUri
	}

	mongoDbName := os.Getenv("MONGO_DB_NAME")
	if mongoDbName == "" {
		mongoDbName = defaultMongoDbName
	}

	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	return &Config{
		Host:          host,
		Port:          port,
		MongoUri:      mongoUri,
		MongoDbName:   mongoDbName,
		MongoUsername: mongoUsername,
		MongoPassword: mongoPassword,
	}
}
