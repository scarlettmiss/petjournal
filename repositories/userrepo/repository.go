package userrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

type UserDBModel struct {
	Id           uuid.UUID `bson:"_id"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
	Deleted      bool      `bson:"deleted"`
	UserType     user.Type `bson:"user_type"`
	Email        string    `bson:"email"`
	PasswordHash string    `bson:"password_hash"`
	Name         string    `bson:"name"`
	Surname      string    `bson:"surname"`
	Phone        string    `bson:"phone,omitempty"`
	Address      string    `bson:"address,omitempty"`
	City         string    `bson:"city,omitempty"`
	State        string    `bson:"state,omitempty"`
	Country      string    `bson:"country,omitempty"`
	Zip          string    `bson:"zip,omitempty"`
	VetId        uuid.UUID `bson:"vet_id"`
}

func ConvertToUserDBModel(user user.User) UserDBModel {
	return UserDBModel{
		Id:           user.Id,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Deleted:      user.Deleted,
		UserType:     user.UserType,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Name:         user.Name,
		Surname:      user.Surname,
		Phone:        user.Phone,
		Address:      user.Address,
		City:         user.City,
		State:        user.State,
		Country:      user.Country,
		Zip:          user.Zip,
		VetId:        user.VetId,
	}
}

func ConvertToUserDomainModel(dbUser UserDBModel) user.User {
	return user.User{
		Id:           dbUser.Id,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
		Deleted:      dbUser.Deleted,
		UserType:     dbUser.UserType,
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		Name:         dbUser.Name,
		Surname:      dbUser.Surname,
		Phone:        dbUser.Phone,
		Address:      dbUser.Address,
		City:         dbUser.City,
		State:        dbUser.State,
		Country:      dbUser.Country,
		Zip:          dbUser.Zip,
		VetId:        dbUser.VetId,
	}
}

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

	dbUser, err := bson.Marshal(ConvertToUserDBModel(u))
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

	var retrievedUser UserDBModel

	filter := bson.M{"_id": id}

	err := r.users.FindOne(context.Background(), filter).Decode(&retrievedUser)
	if err != nil {
		return user.Nil, user.ErrNotFound
	}

	return ConvertToUserDomainModel(retrievedUser), nil
}

func (r *Repository) Users() ([]user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	var users []user.User

	// Define an empty filter to retrieve all users
	filter := bson.M{}

	ctx := context.Background()
	// Perform the find operation
	cursor, err := r.users.Find(ctx, filter)
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode the users
	for cursor.Next(ctx) {
		var u UserDBModel
		err = cursor.Decode(&u)

		if err != nil {
			return users, err
		}

		users = append(users, ConvertToUserDomainModel(u))
	}

	// Check for any errors during cursor iteration
	err = cursor.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *Repository) UpdateUser(u user.User) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	// Define the filter to identify the document to update
	filter := bson.M{"_id": u.Id}

	// Define the update document using the '$set' operator
	replacement, err := bson.Marshal(ConvertToUserDBModel(u))
	if err != nil {
		return user.Nil, err
	}

	// Perform the update operation
	_, err = r.users.ReplaceOne(context.Background(), filter, replacement)
	if err != nil {
		return user.Nil, err
	}

	return u, nil
}

func (r *Repository) DeleteUser(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	filter := bson.M{"_id": id}

	result := r.users.FindOneAndDelete(context.Background(), filter)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
