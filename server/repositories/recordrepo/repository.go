package recordrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/record"
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
	Name           string      `bson:"name,omitempty"`
	Date           time.Time   `bson:"date"`
	Lot            string      `bson:"lot,omitempty"`
	Result         string      `bson:"result,omitempty"`
	Description    string      `bson:"description,omitempty"`
	Notes          string      `bson:"notes,omitempty"`
	AdministeredBy uuid.UUID   `bson:"administered_by,omitempty"`
	VerifiedBy     uuid.UUID   `bson:"verified_by,omitempty"`
	GroupId        uuid.UUID   `bson:"group_id,omitempty"`
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
		GroupId:        r.GroupId,
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
		GroupId:        dbRecord.GroupId,
	}
}

type Repository interface {
	CreateRecord(record record.Record) (record.Record, error)
	CreateRecords(records []record.Record) ([]record.Record, error)
	Record(id uuid.UUID) (record.Record, error)
	Records(includeDel bool) ([]record.Record, error)
	UpdateRecord(record record.Record) (record.Record, error)
	DeleteRecord(id uuid.UUID) error
}

type repository struct {
	mux        sync.Mutex
	recordsCol *mongo.Collection
}

func New(collection *mongo.Collection) Repository {
	return &repository{
		recordsCol: collection,
	}
}

func (r *repository) CreateRecord(rec record.Record) (record.Record, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	id, err := uuid.NewRandom()
	if err != nil {
		return record.Nil, err
	}
	rec.Id = id

	now := time.Now()
	rec.CreatedAt = now
	rec.UpdatedAt = now

	rec.Deleted = false

	dbRec, err := bson.Marshal(ConvertToRecordDBModel(rec))
	if err != nil {
		return record.Nil, err
	}

	_, err = r.recordsCol.InsertOne(context.Background(), dbRec)
	if err != nil {
		return record.Nil, err
	}

	return rec, nil
}

func (r *repository) CreateRecords(recs []record.Record) ([]record.Record, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	groupId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	dbItems := make([]interface{}, len(recs))
	now := time.Now()
	hasError := false
	for i, rec := range recs {
		id, err := uuid.NewRandom()
		if err != nil {
			hasError = true
		}
		rec.Id = id
		rec.GroupId = groupId
		rec.CreatedAt = now
		rec.UpdatedAt = now

		rec.Deleted = false

		var dbRec []byte
		dbRec, err = bson.Marshal(ConvertToRecordDBModel(rec))
		if err != nil {
			hasError = true
		}
		recs[i] = rec
		dbItems[i] = dbRec
	}

	if hasError {
		hasError = true
		return nil, err
	}

	_, err = r.recordsCol.InsertMany(context.Background(), dbItems)
	if err != nil {
		return nil, err
	}

	return recs, nil
}

func (r *repository) Record(id uuid.UUID) (record.Record, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	retrievedRecord, err := r.recordInternal(id)

	return ConvertToRecordDomainModel(retrievedRecord), err

}

func (r *repository) recordInternal(id uuid.UUID) (RecordDBModel, error) {
	var retrievedRecord RecordDBModel

	filter := bson.M{"_id": id}

	err := r.recordsCol.FindOne(context.Background(), filter).Decode(&retrievedRecord)
	if err != nil {
		return RecordDBModel{}, record.ErrNotFound
	}

	return retrievedRecord, nil

}

func (r *repository) Records(includeDel bool) ([]record.Record, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	var records []record.Record

	var filter bson.M

	if includeDel {
		// Define an empty filter to retrieve all users
		filter = bson.M{}
	} else {
		filter = bson.M{"deleted": false}
	}

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

func (r *repository) UpdateRecord(rec record.Record) (record.Record, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	updatedRec, err := r.updateRecordInternal(ConvertToRecordDBModel(rec))
	if err != nil {
		return record.Nil, err
	}

	return ConvertToRecordDomainModel(updatedRec), nil
}

func (r *repository) updateRecordInternal(rec RecordDBModel) (RecordDBModel, error) {
	// Define the filter to identify the document to update
	filter := bson.M{"_id": rec.Id}

	rec.UpdatedAt = time.Now()

	// Define the update document using the '$set' operator
	replacement, err := bson.Marshal(rec)
	if err != nil {
		return RecordDBModel{}, err
	}

	// Perform the update operation
	_, err = r.recordsCol.ReplaceOne(context.Background(), filter, replacement)
	if err != nil {
		return RecordDBModel{}, err
	}

	return rec, nil
}

func (r *repository) DeleteRecord(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	retrievedRecord, err := r.recordInternal(id)
	if err != nil {
		return err
	}

	retrievedRecord.Deleted = true

	_, err = r.updateRecordInternal(retrievedRecord)

	return err
}
