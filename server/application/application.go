package application

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	"github.com/scarlettmiss/petJournal/application/domain/record"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	"github.com/scarlettmiss/petJournal/application/services"
	petService "github.com/scarlettmiss/petJournal/application/services/petService"
	recordService "github.com/scarlettmiss/petJournal/application/services/recordService"
	userService "github.com/scarlettmiss/petJournal/application/services/userService"
	"github.com/scarlettmiss/petJournal/repositories/petrepo"
	"github.com/scarlettmiss/petJournal/repositories/recordrepo"
	"github.com/scarlettmiss/petJournal/repositories/userrepo"
)

/*
what the actor can do.
application talks with all the services
*/
type application struct {
	petService    petService.Service
	userService   userService.Service
	recordService recordService.Service
}

type Options struct {
	PetRepo    petrepo.Repository
	UserRepo   userrepo.Repository
	RecordRepo recordrepo.Repository
}

type Application interface {
	CreateUser(opts services.UserCreateOptions) (user.User, string, error)
	UpdateUser(opts services.UserUpdateOptions, includeDel bool) (user.User, error)
	Users(includeDel bool) ([]user.User, error)
	UsersByType(t user.Type, includeDel bool) ([]user.User, error)
	User(id uuid.UUID) (user.User, error)
	UserByType(id uuid.UUID, t user.Type, includeDel bool) (user.User, error)
	DeleteUser(id uuid.UUID) error
	Authenticate(opts services.LoginOptions) (user.User, string, error)
	PetsByUser(uId uuid.UUID, includeDel bool) (map[uuid.UUID]pet.Pet, error)
	Pet(id uuid.UUID) (pet.Pet, error)
	PetByUser(uId uuid.UUID, id uuid.UUID, includeDel bool) (pet.Pet, error)
	DeletePet(uId uuid.UUID, id uuid.UUID) error
	CreatePet(opts services.PetCreateOptions) (pet.Pet, error)
	UpdatePet(opts services.PetUpdateOptions) (pet.Pet, error)
	CreateRecord(opts services.RecordCreateOptions) (record.Record, error)
	CreateRecords(opts services.RecordsCreateOptions) (map[uuid.UUID]record.Record, error)
	RecordsByUser(uId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error)
	RecordsByUserPet(uId uuid.UUID, pId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error)
	RecordByUserPet(uId uuid.UUID, pId uuid.UUID, tId uuid.UUID, includeDel bool) (record.Record, error)
	UpdateRecord(opts services.RecordUpdateOptions) (record.Record, error)
	DeleteRecordUserPet(uId uuid.UUID, pId uuid.UUID, id uuid.UUID) error
}

func New(opts Options) (Application, error) {
	//init services
	ps, err := petService.New(opts.PetRepo)
	if err != nil {
		return nil, err
	}
	us, err := userService.New(opts.UserRepo)
	if err != nil {
		return nil, err
	}
	rs, err := recordService.New(opts.RecordRepo)
	if err != nil {
		return nil, err
	}

	app := application{petService: ps, userService: us, recordService: rs}

	return &app, nil
}

func (a *application) CreateUser(opts services.UserCreateOptions) (user.User, string, error) {
	return a.userService.CreateUser(opts)
}

func (a *application) UpdateUser(opts services.UserUpdateOptions, includeDel bool) (user.User, error) {
	return a.userService.UpdateUser(opts, includeDel)
}

func (a *application) Users(includeDel bool) ([]user.User, error) {
	return a.userService.Users(includeDel)
}

func (a *application) UsersByType(t user.Type, includeDel bool) ([]user.User, error) {
	return a.userService.UsersByType(t, includeDel)
}

func (a *application) User(id uuid.UUID) (user.User, error) {
	return a.userService.User(id)
}

func (a *application) UserByType(id uuid.UUID, t user.Type, includeDel bool) (user.User, error) {
	return a.userService.UserByType(id, t, includeDel)
}

func (a *application) DeleteUser(id uuid.UUID) error {
	return a.userService.DeleteUser(id)
}

func (a *application) Authenticate(opts services.LoginOptions) (user.User, string, error) {
	return a.userService.Authenticate(opts.Email, opts.Password)
}

func (a *application) PetsByUser(uId uuid.UUID, includeDel bool) (map[uuid.UUID]pet.Pet, error) {
	return a.petService.PetsByUser(uId, includeDel)
}

func (a *application) Pet(id uuid.UUID) (pet.Pet, error) {
	return a.petService.Pet(id)
}

func (a *application) PetByUser(uId uuid.UUID, id uuid.UUID, includeDel bool) (pet.Pet, error) {
	return a.petService.PetByUser(uId, id, includeDel)
}

func (a *application) DeletePet(uId uuid.UUID, id uuid.UUID) error {
	return a.petService.DeletePet(uId, id)
}

func (a *application) CreatePet(opts services.PetCreateOptions) (pet.Pet, error) {
	_, err := a.User(opts.OwnerId)
	if err != nil {
		return pet.Nil, err
	}

	if opts.VetId != uuid.Nil {
		_, err = a.UserByType(opts.VetId, user.Vet, false)
		if err != nil {
			return pet.Nil, err
		}
	}

	return a.petService.CreatePet(opts)
}

func (a *application) UpdatePet(opts services.PetUpdateOptions) (pet.Pet, error) {
	return a.petService.UpdatePet(opts)
}

func (a *application) CreateRecord(opts services.RecordCreateOptions) (record.Record, error) {
	_, err := a.PetByUser(opts.AdministeredBy, opts.PetId, false)
	if err != nil {
		return record.Nil, err
	}

	if opts.VerifiedBy != uuid.Nil {
		_, err = a.UserByType(opts.VerifiedBy, user.Vet, true)
		if err != nil {
			switch err {
			case user.ErrNotFound:
				return record.Nil, record.ErrNotValidVerifier
			default:
				return record.Nil, err
			}
		}
	}

	return a.recordService.CreateRecord(opts)
}

func (a *application) CreateRecords(opts services.RecordsCreateOptions) (map[uuid.UUID]record.Record, error) {
	_, err := a.PetByUser(opts.AdministeredBy, opts.PetId, false)
	if err != nil {
		return nil, err
	}

	if opts.VerifiedBy != uuid.Nil {
		_, err := a.UserByType(opts.VerifiedBy, user.Vet, true)
		if err != nil {
			switch err {
			case user.ErrNotFound:
				return nil, record.ErrNotValidVerifier
			default:
				return nil, err
			}
		}
	}
	return a.recordService.CreateRecords(opts)
}

func (a *application) RecordsByUser(uId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
	pets, err := a.PetsByUser(uId, includeDel)
	if err != nil {
		return nil, err
	}
	return a.recordService.PetsRecords(lo.Keys[uuid.UUID, pet.Pet](pets), includeDel)
}

func (a *application) RecordsByUserPet(uId uuid.UUID, pId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
	_, err := a.PetByUser(uId, pId, false)
	if err != nil {
		return nil, err
	}
	return a.recordService.PetRecords(pId, includeDel)
}

func (a *application) RecordByUserPet(uId uuid.UUID, pId uuid.UUID, tId uuid.UUID, includeDel bool) (record.Record, error) {
	_, err := a.PetByUser(uId, pId, false)
	if err != nil {
		return record.Nil, err
	}
	return a.recordService.PetRecord(pId, tId, includeDel)
}

func (a *application) UpdateRecord(opts services.RecordUpdateOptions) (record.Record, error) {
	if opts.VerifiedBy != uuid.Nil {
		_, err := a.UserByType(opts.VerifiedBy, user.Vet, true)
		if err != nil {
			switch err {
			case user.ErrNotFound:
				return record.Nil, record.ErrNotValidVerifier
			default:
				return record.Nil, err
			}
		}
	}

	return a.recordService.UpdateRecord(opts)
}

func (a *application) DeleteRecordUserPet(uId uuid.UUID, pId uuid.UUID, id uuid.UUID) error {
	_, err := a.RecordByUserPet(uId, pId, id, false)
	if err != nil {
		return err
	}

	return a.recordService.DeleteRecord(id)
}
