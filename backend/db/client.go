package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DisconnectFunc func()

func GetClient(uri string, username string, password string) (*mongo.Client, DisconnectFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// credentials := options.Credential{
	// 	Username: username,
	// 	Password: password,
	// }
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri)) //.SetAuth(credentials))
	return client, func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
}
