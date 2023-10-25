package recordrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/record"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

type RecordDBModel struct {
	Id             uuid.UUID   `bson:"_id"`
	CreatedAt      time.Time   `bson:"createdAt"`
	UpdatedAt      time.Time   `bson:"updatedAt"`
	Deleted        bool        `bson:"deleted"`
	PetId          uuid.UUID   `bson:"petId"`
	RecordType     record.Type `bson:"recordType"`
	Name           string      `bson:"name"`
	Date           time.Time   `bson:"date"`
	Lot            string      `bson:"lot,omitempty"`
	Result         string      `bson:"result,omitempty"`
	Description    string      `bson:"description,omitempty"`
	Notes          string      `bson:"notes,omitempty"`
	AdministeredBy uuid.UUID   `bson:"administered_by,omitempty"`
	VerifiedBy     uuid.UUID   `bson:"verified_by,omitempty"`
	NextDate       time.Time   `bson:"next_date,omitempty"`
}

func ConvertToRecordDBModel(r record.Record) RecordDBModel {
	return RecordDBModel{
		Id:             r.Id,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
		Deleted:        r.Deleted,
		PetId:          r.PetId,
		RecordType:     r.RecordType,
		Name:           r.Name,
		Date:           r.Date,
		Lot:            r.Lot,
		Result:         r.Result,
		Description:    r.Description,
		Notes:          r.Notes,
		AdministeredBy: r.AdministeredBy,
		VerifiedBy:     r.VerifiedBy,
		NextDate:       r.NextDate,
	}
}

func ConvertToRecordDomainModel(dbRecord RecordDBModel) record.Record {
	return record.Record{
		Id:             dbRecord.Id,
		CreatedAt:      dbRecord.CreatedAt,
		UpdatedAt:      dbRecord.UpdatedAt,
		Deleted:        dbRecord.Deleted,
		PetId:          dbRecord.PetId,
		RecordType:     dbRecord.RecordType,
		Name:           dbRecord.Name,
		Date:           dbRecord.Date,
		Lot:            dbRecord.Lot,
		Result:         dbRecord.Result,
		Description:    dbRecord.Description,
		Notes:          dbRecord.Notes,
		AdministeredBy: dbRecord.AdministeredBy,
		VerifiedBy:     dbRecord.VerifiedBy,
		NextDate:       dbRecord.NextDate,
	}
}

type Repository struct {
	mux        sync.Mutex
	recordsCol *mongo.Collection
}

func New(collection *mongo.Collection) *Repository {
	return &Repository{
		recordsCol: collection,
	}
}

func (r *Repository) CreateRecord(rec record.Record) (record.Record, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	id, err := uuid.NewUUID()
	if err != nil {
		return record.Nil, err
	}
	rec.Id = id

	now := time.Now()
	rec.CreatedAt = now
	rec.UpdatedAt = now

	rec.Deleted = false

	dbUser, err := bson.Marshal(ConvertToRecordDBModel(rec))
	if err != nil {
		return record.Nil, err
	}

	_, err = r.recordsCol.InsertOne(context.Background(), dbUser)
	if err != nil {
		return record.Nil, err
	}

	return rec, nil
}

func (r *Repository) Record(id uuid.UUID) (record.Record, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	var retrievedRecord RecordDBModel

	filter := bson.M{"_id": id}

	err := r.recordsCol.FindOne(context.Background(), filter).Decode(&retrievedRecord)
	if err != nil {
		return record.Nil, user.ErrNotFound
	}

	return ConvertToRecordDomainModel(retrievedRecord), nil

}

func (r *Repository) Records() ([]record.Record, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	var records []record.Record

	// Define an empty filter to retrieve all users
	filter := bson.M{}

	ctx := context.Background()
	// Perform the find operation
	cursor, err := r.recordsCol.Find(ctx, filter)
	if err != nil {
		return records, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode the users
	for cursor.Next(ctx) {
		var u RecordDBModel
		err = cursor.Decode(&u)

		if err != nil {
			return records, err
		}

		records = append(records, ConvertToRecordDomainModel(u))
	}

	// Check for any errors during cursor iteration
	err = cursor.Err()
	if err != nil {
		return records, err
	}

	return records, nil
}

func (r *Repository) UpdateRecord(rec record.Record) (record.Record, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	// Define the filter to identify the document to update
	filter := bson.M{"_id": rec.Id}

	// Define the update document using the '$set' operator
	replacement, err := bson.Marshal(ConvertToRecordDBModel(rec))
	if err != nil {
		return record.Nil, err
	}

	// Perform the update operation
	_, err = r.recordsCol.ReplaceOne(context.Background(), filter, replacement)
	if err != nil {
		return record.Nil, err
	}

	return rec, nil
}

func (r *Repository) DeleteRecord(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	filter := bson.M{"_id": id}

	result := r.recordsCol.FindOneAndDelete(context.Background(), filter)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
