package main_test

import (
	"context"
	"fmt"
	"github.com/scarlettmiss/petJournal/application"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	authService "github.com/scarlettmiss/petJournal/application/services/authService"
	petService "github.com/scarlettmiss/petJournal/application/services/petService"
	recordService "github.com/scarlettmiss/petJournal/application/services/recordService"
	userService "github.com/scarlettmiss/petJournal/application/services/userService"
	"github.com/scarlettmiss/petJournal/repositories/petrepo"
	"github.com/scarlettmiss/petJournal/repositories/recordrepo"
	"github.com/scarlettmiss/petJournal/repositories/userrepo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

/*
*
testing suit for the application.
*/
func TestApplicationCreation(t *testing.T) {
	//init db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.Nil(t, err)
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	assert.Nil(t, err)

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	db := client.Database("petjournal-test")

	//init repos
	petsCollection := db.Collection("pets")
	petRepo := petrepo.New(petsCollection)

	usersCollection := db.Collection("users")
	userRepo := userrepo.New(usersCollection)

	recordsCollection := db.Collection("treatments")
	recordRepo := recordrepo.New(recordsCollection)
	//init services
	ps, err := petService.New(petRepo)
	assert.Nil(t, err)
	us, err := userService.New(userRepo)
	assert.Nil(t, err)
	rs, err := recordService.New(recordRepo)
	assert.Nil(t, err)

	//pass services to application
	opts := application.Options{PetService: ps, UserService: us, RecordService: rs}
	app := application.New(opts)

	assert.NotNil(t, &app)
}

/*
*
testing suit for the application.
*/
func TestUserCreation(t *testing.T) {
	//init db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.Nil(t, err)
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	assert.Nil(t, err)

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	db := client.Database("petjournal-test")

	//init repos
	petsCollection := db.Collection("pets")
	petRepo := petrepo.New(petsCollection)

	usersCollection := db.Collection("users")
	userRepo := userrepo.New(usersCollection)

	recordsCollection := db.Collection("records")
	recordRepo := recordrepo.New(recordsCollection)
	//init services
	ps, err := petService.New(petRepo)
	assert.Nil(t, err)
	us, err := userService.New(userRepo)
	assert.Nil(t, err)
	rs, err := recordService.New(recordRepo)
	assert.Nil(t, err)

	//pass services to application
	opts := application.Options{PetService: ps, UserService: us, RecordService: rs}
	app := application.New(opts)

	u := user.Nil
	u.UserType = user.Vet

	_, err = app.CreateUser(u)
	assert.NotNil(t, err)

	u.Name = "name"
	_, err = app.CreateUser(u)
	assert.NotNil(t, err)

	u.Surname = "surname"
	_, err = app.CreateUser(u)
	assert.NotNil(t, err)

	u.Email = "surname@mail.com"
	_, err = app.CreateUser(u)
	assert.NotNil(t, err)

	pass, err := authService.HashPassword("password")
	assert.Nil(t, err)
	u.PasswordHash = pass
	_, err = app.CreateUser(u)
	assert.Nil(t, err)

}
