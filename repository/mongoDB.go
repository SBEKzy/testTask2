package repository

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func ConnectDB() *mongo.Database {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("DB_HOST")))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	db := mongoClient.Database(viper.GetString("DB_NAME"))
	return db
}
