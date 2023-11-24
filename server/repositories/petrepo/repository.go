package petrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

type PetDBModel struct {
	Id          uuid.UUID         `bson:"_id"`
	CreatedAt   time.Time         `bson:"created_at"`
	UpdatedAt   time.Time         `bson:"updated_at"`
	Deleted     bool              `bson:"deleted"`
	Name        string            `bson:"name"`
	DateOfBirth time.Time         `bson:"date_of_birth,omitempty"`
	Gender      string            `bson:"gender,omitempty"`
	BreedName   string            `bson:"breed_name,omitempty"`
	Colors      []string          `bson:"colors,omitempty"`
	Description string            `bson:"description,omitempty"`
	Pedigree    string            `bson:"pedigree,omitempty"`
	Microchip   string            `bson:"microchip,omitempty"`
	OwnerID     uuid.UUID         `bson:"owner_id"`
	VetID       uuid.UUID         `bson:"vet_id,omitempty"`
	Metas       map[string]string `bson:"metas,omitempty"`
	Avatar      string            `bson:"avatar,omitempty"`
}

func ConvertToPetDBModel(pet pet.Pet) PetDBModel {
	return PetDBModel{
		Id:          pet.Id,
		CreatedAt:   pet.CreatedAt,
		UpdatedAt:   pet.UpdatedAt,
		Deleted:     pet.Deleted,
		Name:        pet.Name,
		DateOfBirth: pet.DateOfBirth,
		Gender:      string(pet.Gender),
		BreedName:   pet.BreedName,
		Colors:      pet.Colors,
		Description: pet.Description,
		Pedigree:    pet.Pedigree,
		Microchip:   pet.Microchip,
		OwnerID:     pet.OwnerId,
		VetID:       pet.VetId,
		Metas:       pet.Metas,
		Avatar:      pet.Avatar,
	}
}

func ConvertToPetDomainModel(dbPet PetDBModel) pet.Pet {
	return pet.Pet{
		Id:          dbPet.Id,
		CreatedAt:   dbPet.CreatedAt,
		UpdatedAt:   dbPet.UpdatedAt,
		Deleted:     dbPet.Deleted,
		Name:        dbPet.Name,
		DateOfBirth: dbPet.DateOfBirth,
		Gender:      pet.Gender(dbPet.Gender),
		BreedName:   dbPet.BreedName,
		Colors:      dbPet.Colors,
		Description: dbPet.Description,
		Pedigree:    dbPet.Pedigree,
		Microchip:   dbPet.Microchip,
		OwnerId:     dbPet.OwnerID,
		VetId:       dbPet.VetID,
		Metas:       dbPet.Metas,
		Avatar:      dbPet.Avatar,
	}
}

type Repository struct {
	mux  sync.Mutex
	pets *mongo.Collection
}

func New(collection *mongo.Collection) *Repository {
	return &Repository{
		pets: collection,
	}
}

func (r *Repository) CreatePet(p pet.Pet) (pet.Pet, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	id, err := uuid.NewUUID()
	if err != nil {
		return pet.Nil, err
	}
	p.Id = id

	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now

	p.Deleted = false

	dbPet, err := bson.Marshal(ConvertToPetDBModel(p))
	if err != nil {
		return pet.Nil, err
	}

	_, err = r.pets.InsertOne(context.Background(), dbPet)
	if err != nil {
		return pet.Nil, err
	}

	return p, nil
}

func (r *Repository) Pet(id uuid.UUID) (pet.Pet, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	retrievedPet, err := r.petInternal(id)
	return ConvertToPetDomainModel(retrievedPet), err
}

func (r *Repository) petInternal(id uuid.UUID) (PetDBModel, error) {
	var retrievedPet PetDBModel

	filter := bson.M{"_id": id}

	err := r.pets.FindOne(context.Background(), filter).Decode(&retrievedPet)
	if err != nil {
		return PetDBModel{}, pet.ErrNotFound
	}
	return retrievedPet, nil
}

func (r *Repository) Pets(includeDel bool) ([]pet.Pet, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	var pets []pet.Pet

	// Define an empty filter to retrieve all pets
	var filter bson.M

	if includeDel {
		filter = bson.M{}
	} else {
		filter = bson.M{"deleted": false}
	}

	ctx := context.Background()
	// Perform the find operation
	cursor, err := r.pets.Find(ctx, filter)
	if err != nil {
		return pets, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode the users
	for cursor.Next(ctx) {
		var p PetDBModel
		err = cursor.Decode(&p)

		if err != nil {
			return pets, err
		}

		pets = append(pets, ConvertToPetDomainModel(p))
	}

	// Check for any errors during cursor iteration
	err = cursor.Err()
	if err != nil {
		return pets, err
	}

	return pets, nil
}

func (r *Repository) UpdatePet(p pet.Pet) (pet.Pet, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	updatedPet, err := r.updatePetInternal(ConvertToPetDBModel(p))
	if err != nil {
		return pet.Nil, err
	}

	return ConvertToPetDomainModel(updatedPet), nil
}

func (r *Repository) updatePetInternal(p PetDBModel) (PetDBModel, error) {
	// Define the filter to identify the document to update
	filter := bson.M{"_id": p.Id}

	// Define the update document using the '$set' operator
	p.UpdatedAt = time.Now()
	replacement, err := bson.Marshal(p)
	if err != nil {
		return PetDBModel{}, err
	}

	// Perform the update operation
	_, err = r.pets.ReplaceOne(context.Background(), filter, replacement)
	if err != nil {
		return PetDBModel{}, err
	}

	return p, nil
}

func (r *Repository) DeletePet(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	retrievedPet, err := r.petInternal(id)
	if err != nil {
		return err
	}

	retrievedPet.Deleted = true

	_, err = r.updatePetInternal(retrievedPet)

	return err
}
