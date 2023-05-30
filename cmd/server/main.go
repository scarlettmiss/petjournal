package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/scarlettmiss/bestPal/api"
	"github.com/scarlettmiss/bestPal/api/config"
	"github.com/scarlettmiss/bestPal/application"
	petService "github.com/scarlettmiss/bestPal/application/services/petService"
	treatmentService "github.com/scarlettmiss/bestPal/application/services/treatmentService"
	userService "github.com/scarlettmiss/bestPal/application/services/userService"
	"github.com/scarlettmiss/bestPal/repositories/petrepo"
	"github.com/scarlettmiss/bestPal/repositories/treatmentrepo"
	"github.com/scarlettmiss/bestPal/repositories/userrepo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	//init db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URL")))
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

	treatmentsCollection := db.Collection("treatments")
	treatmentRepo := treatmentrepo.New(treatmentsCollection)
	//init services
	ps, err := petService.New(petRepo)
	if err != nil {
		panic(err)
	}
	us, err := userService.New(userRepo)
	if err != nil {
		panic(err)
	}
	ts, err := treatmentService.New(treatmentRepo)
	if err != nil {
		panic(err)
	}

	//pass services to application
	opts := application.Options{PetService: ps, UserService: us, TreatmentService: ts}
	app := application.New(opts)

	if err != nil {
		panic(err)
	}

	restServer := api.New(app)

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
