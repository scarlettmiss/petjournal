package application

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	"github.com/scarlettmiss/petJournal/application/domain/record"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	jwtService "github.com/scarlettmiss/petJournal/application/services/jwtService"
	petService "github.com/scarlettmiss/petJournal/application/services/petService"
	recordService "github.com/scarlettmiss/petJournal/application/services/recordService"
	userService "github.com/scarlettmiss/petJournal/application/services/userService"
	"github.com/scarlettmiss/petJournal/utils"
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

func (a *Application) CreateUser(u user.User) (user.User, error) {
	err := a.CheckEmail(u.Email, u.Id, true)
	if err != nil {
		return user.Nil, err
	}

	return a.userService.CreateUser(u)
}

func (a *Application) UserToken(u user.User) (string, error) {
	return jwtService.GenerateJWT(u.Id, u.UserType)
}

func (a *Application) CheckEmail(email string, id uuid.UUID, includeDel bool) error {
	if !utils.IsEmailValid(email) {
		return user.ErrNoValidMail
	}

	u, ok := a.userService.UserByEmail(email, includeDel)

	if !ok {
		return nil
	}

	if u.Id == id {
		return nil
	}

	return user.ErrMailExists
}

func (a *Application) UpdateUser(u user.User, includeDel bool) (user.User, error) {
	err := a.CheckEmail(u.Email, u.Id, includeDel)
	if err != nil {
		return u, err
	}
	return a.userService.UpdateUser(u)
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
	u, err := a.userService.UserByType(id, t, includeDel)
	if err != nil {
		return user.Nil, err
	}
	return u, nil
}

func (a *Application) DeleteUser(id uuid.UUID) error {
	return a.userService.DeleteUser(id)
}

func (a *Application) Authenticate(email string, password string) (user.User, error) {
	u, err := a.userService.Authenticate(email, password)
	if err != nil {
		return user.Nil, err
	}

	return u, nil
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

func (a *Application) PetByOwner(uId uuid.UUID, id uuid.UUID, includeDel bool) (pet.Pet, error) {
	return a.petService.PetByOwner(uId, id, includeDel)
}

func (a *Application) DeletePet(id uuid.UUID) error {
	return a.petService.DeletePet(id)
}

func (a *Application) RemoveVet(id uuid.UUID) error {
	return a.petService.RemoveVet(id)
}

func (a *Application) CreatePet(p pet.Pet) (pet.Pet, error) {
	return a.petService.CreatePet(p)
}

func (a *Application) UpdatePet(p pet.Pet) (pet.Pet, error) {
	return a.petService.UpdatePet(p)
}

func (a *Application) CreateRecord(t record.Record) (record.Record, error) {
	return a.recordService.CreateRecord(t)
}

func (a *Application) RecordsByUser(uId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
	pets, err := a.PetsByUser(uId, includeDel)
	if err != nil {
		return nil, err
	}
	return a.recordService.PetsRecords(lo.Keys[uuid.UUID, pet.Pet](pets), includeDel)
}

func (a *Application) RecordsByPet(pId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
	return a.recordService.PetRecords(pId, includeDel)
}

func (a *Application) RecordByPet(pId uuid.UUID, tId uuid.UUID, includeDel bool) (record.Record, error) {
	return a.recordService.PetRecord(pId, tId, includeDel)
}

func (a *Application) UpdateRecord(t record.Record) (record.Record, error) {
	return a.recordService.UpdateRecord(t)
}

func (a *Application) DeleteRecord(id uuid.UUID) error {
	return a.recordService.DeleteRecord(id)
}
