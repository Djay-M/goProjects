package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbConnectString = "mongodb://localhost:27017/"
const dbName = "netflix"
const colName = "watchlist"

var Collections *mongo.Collection

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// create client option
	clientOption := options.Client().ApplyURI(dbConnectString)

	// connect to DB and create a client
	client, dbErr := mongo.Connect(context.TODO(), clientOption)

	handleErrors(dbErr)
	fmt.Println("DB connection successful")
	Collections = client.Database(dbName).Collection(colName)
	fmt.Println("Collections are ready !!")
}
