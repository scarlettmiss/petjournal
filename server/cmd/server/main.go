package main

import (
	"context"
	"fmt"
	"github.com/scarlettmiss/petJournal/api"
	"github.com/scarlettmiss/petJournal/api/config"
	"github.com/scarlettmiss/petJournal/application"
	"github.com/scarlettmiss/petJournal/repositories/petrepo"
	"github.com/scarlettmiss/petJournal/repositories/recordrepo"
	"github.com/scarlettmiss/petJournal/repositories/userrepo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"time"
)

import "embed"

//go:embed all:public
var ui embed.FS

func main() {
	uri := os.Getenv("DB_URL")
	if uri == "" {
		log.Fatal("You must set your 'DB_URL' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	//init db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	db := client.Database(os.Getenv("DB_NAME"))

	//init repos
	petsCollection := db.Collection("pets")
	petRepo := petrepo.New(petsCollection)

	usersCollection := db.Collection("users")
	userRepo := userrepo.New(usersCollection)

	recordsCollection := db.Collection("records")
	recordRepo := recordrepo.New(recordsCollection)

	//pass services to application
	opts := application.Options{PetRepo: petRepo, UserRepo: userRepo, RecordRepo: recordRepo}
	app, err := application.New(opts)
	if err != nil {
		panic(err)
	}

	restServer := api.New(app, ui)

	go func() { // Start listening and serving requests
		err = restServer.Run(config.Host + ":" + config.Port)

		if err != nil {
			panic(err)
		}
	}()
	//ctrl + c to stop server
	waitForInterrupt := make(chan os.Signal, 1)
	signal.Notify(waitForInterrupt, os.Interrupt, os.Kill)

	<-waitForInterrupt
}
