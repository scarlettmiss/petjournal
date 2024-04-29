package main_test

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	"github.com/scarlettmiss/petJournal/application/services"
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
testing suit for the User actions.
*/
func TestUser(t *testing.T) {
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
	if err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	db := client.Database("petjournal-test")

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
	assert.Nil(t, err)

	createOptions := services.UserCreateOptions{}
	_, _, err = app.CreateUser(createOptions)
	assert.EqualError(t, err, user.ErrNoValidType.Error())

	createOptions.UserType = "test"
	_, _, err = app.CreateUser(createOptions)
	assert.EqualError(t, err, user.ErrNoValidType.Error())

	createOptions.UserType = "vet"
	_, _, err = app.CreateUser(createOptions)
	assert.EqualError(t, err, user.ErrNoValidMail.Error())

	createOptions.Email = "mail@mail"
	_, _, err = app.CreateUser(createOptions)
	assert.EqualError(t, err, user.ErrNoValidMail.Error())

	createOptions.Email = "mail@mail.com"
	_, _, err = app.CreateUser(createOptions)
	assert.EqualError(t, err, user.ErrPasswordLength.Error())

	createOptions.Password = "12345678aA!"
	_, _, err = app.CreateUser(createOptions)
	assert.EqualError(t, err, user.ErrNoValidName.Error())

	createOptions.Name = "testName"
	_, _, err = app.CreateUser(createOptions)
	assert.EqualError(t, err, user.ErrNoValidSurname.Error())

	createOptions.Surname = "testSurname"
	u, token, err := app.CreateUser(createOptions)
	assert.Nil(t, err)
	assert.NotEqual(t, token, "")

	_, _, err = app.CreateUser(createOptions)
	assert.EqualError(t, err, user.ErrMailExists.Error())

	updateOptions := services.UserUpdateOptions{}
	_, err = app.UpdateUser(updateOptions, false)
	assert.EqualError(t, err, user.ErrNotFound.Error())

	updateOptions.Id = u.Id
	_, err = app.UpdateUser(updateOptions, false)
	assert.EqualError(t, err, user.ErrNoValidMail.Error())

	updateOptions.Email = "mail@mail.com"
	_, err = app.UpdateUser(updateOptions, false)
	assert.EqualError(t, err, user.ErrNoValidName.Error())

	updateOptions.Name = "testName"
	_, err = app.UpdateUser(updateOptions, false)
	assert.EqualError(t, err, user.ErrNoValidSurname.Error())

	updateOptions.Surname = "testSurname2"
	u, err = app.UpdateUser(updateOptions, false)
	assert.Nil(t, err)

	u, err = app.User(u.Id)
	assert.Nil(t, err)

	_, err = app.User(uuid.Nil)
	assert.EqualError(t, err, user.ErrNotFound.Error())

	err = app.DeleteUser(u.Id)
	assert.Nil(t, err)

	u, err = app.User(u.Id)
	assert.Nil(t, err)

	u, err = app.User(u.Id)
	assert.Nil(t, err)

	u, err = app.User(u.Id)
	assert.Nil(t, err)

	_, err = app.Users(true)
	assert.Nil(t, err)

	db.Drop(ctx)
}
