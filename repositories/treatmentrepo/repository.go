package treatmentrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/treatment"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

type TreatmentDBModel struct {
	Id             uuid.UUID      `bson:"_id"`
	CreatedAt      time.Time      `bson:"createdAt"`
	UpdatedAt      time.Time      `bson:"updatedAt"`
	Deleted        bool           `bson:"deleted"`
	PetId          uuid.UUID      `bson:"petId"`
	TreatmentType  treatment.Type `bson:"treatmentType"`
	Name           string         `bson:"name"`
	Date           time.Time      `bson:"date"`
	Lot            string         `bson:"lot,omitempty"`
	Result         string         `bson:"result,omitempty"`
	Description    string         `bson:"description,omitempty"`
	Notes          string         `bson:"notes,omitempty"`
	AdministeredBy uuid.UUID      `bson:"administered_by,omitempty"`
	VerifiedBy     uuid.UUID      `bson:"verified_by,omitempty"`
	RecurringRule  string         `bson:"recurring_rule,omitempty"`
}

func ConvertToTreatmentDBModel(treatment treatment.Treatment) TreatmentDBModel {
	return TreatmentDBModel{
		Id:             treatment.Id,
		CreatedAt:      treatment.CreatedAt,
		UpdatedAt:      treatment.UpdatedAt,
		Deleted:        treatment.Deleted,
		PetId:          treatment.PetId,
		TreatmentType:  treatment.TreatmentType,
		Name:           treatment.Name,
		Date:           treatment.Date,
		Lot:            treatment.Lot,
		Result:         treatment.Result,
		Description:    treatment.Description,
		Notes:          treatment.Notes,
		AdministeredBy: treatment.AdministeredBy,
		VerifiedBy:     treatment.VerifiedBy,
		RecurringRule:  treatment.RecurringRule,
	}
}

func ConvertToTreatmentDomainModel(dbTreatment TreatmentDBModel) treatment.Treatment {
	return treatment.Treatment{
		Id:             dbTreatment.Id,
		CreatedAt:      dbTreatment.CreatedAt,
		UpdatedAt:      dbTreatment.UpdatedAt,
		Deleted:        dbTreatment.Deleted,
		PetId:          dbTreatment.PetId,
		TreatmentType:  dbTreatment.TreatmentType,
		Name:           dbTreatment.Name,
		Date:           dbTreatment.Date,
		Lot:            dbTreatment.Lot,
		Result:         dbTreatment.Result,
		Description:    dbTreatment.Description,
		Notes:          dbTreatment.Notes,
		AdministeredBy: dbTreatment.AdministeredBy,
		VerifiedBy:     dbTreatment.VerifiedBy,
		RecurringRule:  dbTreatment.RecurringRule,
	}
}

type Repository struct {
	mux        sync.Mutex
	treatments *mongo.Collection
}

func New(collection *mongo.Collection) *Repository {
	return &Repository{
		treatments: collection,
	}
}

func (r *Repository) CreateTreatment(t treatment.Treatment) (treatment.Treatment, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	id, err := uuid.NewUUID()
	if err != nil {
		return treatment.Nil, err
	}
	t.Id = id

	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	t.Deleted = false

	dbUser, err := bson.Marshal(ConvertToTreatmentDBModel(t))
	if err != nil {
		return treatment.Nil, err
	}

	_, err = r.treatments.InsertOne(context.Background(), dbUser)
	if err != nil {
		return treatment.Nil, err
	}

	return t, nil
}

func (r *Repository) Treatment(id uuid.UUID) (treatment.Treatment, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	var retrievedTreatment TreatmentDBModel

	filter := bson.M{"_id": id}

	err := r.treatments.FindOne(context.Background(), filter).Decode(&retrievedTreatment)
	if err != nil {
		return treatment.Nil, user.ErrNotFound
	}

	return ConvertToTreatmentDomainModel(retrievedTreatment), nil

}

func (r *Repository) Treatments() ([]treatment.Treatment, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	var treatments []treatment.Treatment

	// Define an empty filter to retrieve all users
	filter := bson.M{}

	ctx := context.Background()
	// Perform the find operation
	cursor, err := r.treatments.Find(ctx, filter)
	if err != nil {
		return treatments, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode the users
	for cursor.Next(ctx) {
		var u TreatmentDBModel
		err = cursor.Decode(&u)

		if err != nil {
			return treatments, err
		}

		treatments = append(treatments, ConvertToTreatmentDomainModel(u))
	}

	// Check for any errors during cursor iteration
	err = cursor.Err()
	if err != nil {
		return treatments, err
	}

	return treatments, nil
}

func (r *Repository) UpdateTreatment(t treatment.Treatment) (treatment.Treatment, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	// Define the filter to identify the document to update
	filter := bson.M{"_id": t.Id}

	// Define the update document using the '$set' operator
	replacement, err := bson.Marshal(ConvertToTreatmentDBModel(t))
	if err != nil {
		return treatment.Nil, err
	}

	// Perform the update operation
	_, err = r.treatments.ReplaceOne(context.Background(), filter, replacement)
	if err != nil {
		return treatment.Nil, err
	}

	return t, nil
}

func (r *Repository) DeleteTreatment(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	filter := bson.M{"_id": id}

	result := r.treatments.FindOneAndDelete(context.Background(), filter)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
