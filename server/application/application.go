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
)

/*
what the actor can do.
application talks with all the services
*/
type Application struct {
	petService    petService.Service
	userService   userService.Service
	recordService recordService.Service
}

type Options struct {
	PetService    petService.Service
	UserService   userService.Service
	RecordService recordService.Service
}

func New(opts Options) *Application {
	app := Application{petService: opts.PetService, userService: opts.UserService, recordService: opts.RecordService}

	return &app
}

func (a *Application) CreateUser(opts services.UserCreateOptions) (user.User, string, error) {
	return a.userService.CreateUser(opts)
}

func (a *Application) UpdateUser(opts services.UserUpdateOptions, includeDel bool) (user.User, error) {
	return a.userService.UpdateUser(opts, includeDel)
}

func (a *Application) Users(includeDel bool) ([]user.User, error) {
	return a.userService.Users(includeDel)
}

func (a *Application) UsersByType(t user.Type, includeDel bool) ([]user.User, error) {
	return a.userService.UsersByType(t, includeDel)
}

func (a *Application) User(id uuid.UUID) (user.User, error) {
	return a.userService.User(id)
}

func (a *Application) UserByType(id uuid.UUID, t user.Type, includeDel bool) (user.User, error) {
	return a.userService.UserByType(id, t, includeDel)
}

func (a *Application) DeleteUser(id uuid.UUID) error {
	return a.userService.DeleteUser(id)
}

func (a *Application) Authenticate(opts services.LoginOptions) (user.User, string, error) {
	return a.userService.Authenticate(opts.Email, opts.Password)
}

func (a *Application) PetsByUser(uId uuid.UUID, includeDel bool) (map[uuid.UUID]pet.Pet, error) {
	return a.petService.PetsByUser(uId, includeDel)
}

func (a *Application) Pet(id uuid.UUID) (pet.Pet, error) {
	return a.petService.Pet(id)
}

func (a *Application) PetByUser(uId uuid.UUID, id uuid.UUID, includeDel bool) (pet.Pet, error) {
	return a.petService.PetByUser(uId, id, includeDel)
}

func (a *Application) DeletePet(uId uuid.UUID, id uuid.UUID) error {
	return a.petService.DeletePet(uId, id)
}

func (a *Application) CreatePet(opts services.PetCreateOptions) (pet.Pet, error) {
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

func (a *Application) UpdatePet(opts services.PetUpdateOptions) (pet.Pet, error) {
	return a.petService.UpdatePet(opts)
}

func (a *Application) CreateRecord(opts services.RecordCreateOptions) (record.Record, error) {
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

func (a *Application) CreateRecords(opts services.RecordsCreateOptions) (map[uuid.UUID]record.Record, error) {
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

func (a *Application) RecordsByUser(uId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
	pets, err := a.PetsByUser(uId, includeDel)
	if err != nil {
		return nil, err
	}
	return a.recordService.PetsRecords(lo.Keys[uuid.UUID, pet.Pet](pets), includeDel)
}

func (a *Application) RecordsByUserPet(uId uuid.UUID, pId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
	_, err := a.PetByUser(uId, pId, false)
	if err != nil {
		return nil, err
	}
	return a.recordService.PetRecords(pId, includeDel)
}

func (a *Application) RecordByUserPet(uId uuid.UUID, pId uuid.UUID, tId uuid.UUID, includeDel bool) (record.Record, error) {
	_, err := a.PetByUser(uId, pId, false)
	if err != nil {
		return record.Nil, err
	}
	return a.recordService.PetRecord(pId, tId, includeDel)
}

func (a *Application) UpdateRecord(opts services.RecordUpdateOptions) (record.Record, error) {
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

func (a *Application) DeleteRecordUserPet(uId uuid.UUID, pId uuid.UUID, id uuid.UUID) error {
	_, err := a.RecordByUserPet(uId, pId, id, false)
	if err != nil {
		return err
	}

	return a.recordService.DeleteRecord(id)
}
