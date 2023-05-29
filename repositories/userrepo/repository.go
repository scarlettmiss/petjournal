package userrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

type Repository struct {
	mux   sync.Mutex
	users *mongo.Collection
}

func New(collection *mongo.Collection) *Repository {
	return &Repository{
		users: collection,
	}
}

func (r *Repository) CreateUser(u user.User) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	id, err := uuid.NewUUID()
	if err != nil {
		return user.Nil, err
	}
	u.Id = id

	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now

	u.Deleted = false

	dbUser, err := bson.Marshal(u)
	if err != nil {
		return user.Nil, err
	}
	_, err = r.users.InsertOne(context.Background(), dbUser)
	if err != nil {
		return user.Nil, err
	}

	return u, nil
}

func (r *Repository) User(id uuid.UUID) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	var retrievedUser user.User
	err := r.users.FindOne(context.Background(), bson.M{"id": id}).Decode(&retrievedUser)
	if err != nil {
		log.Fatal(err)
		return user.Nil, user.ErrNotFound
	}
	return retrievedUser, nil
}

func (r *Repository) Users() ([]user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	// Define an empty filter to retrieve all users
	filter := bson.M{}

	// Perform the find operation
	cursor, err := r.users.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode the users
	var users []user.User
	for cursor.Next(context.Background()) {
		var u user.User
		if err := cursor.Decode(&u); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	// Check for any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return users, err
	}

	return users, nil
}

func (r *Repository) UpdateUser(u user.User) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	// Define the filter to identify the document to update
	filter := bson.M{"id": u.Id}

	// Define the update document using the '$set' operator
	dbUser, err := bson.Marshal(u)
	if err != nil {
		return user.Nil, err
	}

	var updateDoc bson.M
	err = bson.Unmarshal(dbUser, &updateDoc)
	if err != nil {
		log.Fatal(err)
	}

	update := bson.M{"$set": updateDoc}

	// Specify options for the update operation (optional)
	opts := options.Update().SetUpsert(false)

	// Perform the update operation
	_, err = r.users.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
		return user.Nil, err
	}

	return u, nil
}

func (r *Repository) DeleteUser(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	result := r.users.FindOneAndDelete(context.Background(), bson.M{"id": id})
	if result.Err() != nil {
		log.Fatal(result.Err())
		return result.Err()
	}
	return nil
}
