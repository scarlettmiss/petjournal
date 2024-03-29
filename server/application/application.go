package application

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	"github.com/scarlettmiss/petJournal/application/domain/record"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	petService "github.com/scarlettmiss/petJournal/application/services/petService"
	recordService "github.com/scarlettmiss/petJournal/application/services/recordService"
	userService "github.com/scarlettmiss/petJournal/application/services/userService"
	"github.com/scarlettmiss/petJournal/utils"
	authUtils "github.com/scarlettmiss/petJournal/utils/authorization"
	jwtUtils "github.com/scarlettmiss/petJournal/utils/jwt"
	"time"
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

type PetCreateOptions struct {
	OwnerId     uuid.UUID
	VetId       uuid.UUID
	Avatar      string
	Name        string
	DateOfBirth time.Time
	Gender      string
	BreedName   string
	Colors      []string
	Description string
	Pedigree    string
	Microchip   string
	Metas       map[string]string
}

type PetUpdateOptions struct {
	Id          uuid.UUID
	Avatar      string
	Name        string
	DateOfBirth time.Time
	Gender      string
	BreedName   string
	Colors      []string
	Description string
	Pedigree    string
	Microchip   string
	OwnerId     uuid.UUID
	VetId       uuid.UUID
	Metas       map[string]string
}

type RecordCreateOptions struct {
	PetId          uuid.UUID
	RecordType     string
	Name           string
	Date           time.Time
	Lot            string
	Result         string
	Description    string
	Notes          string
	AdministeredBy uuid.UUID
	VerifiedBy     uuid.UUID
}

type RecordsCreateOptions struct {
	PetId          uuid.UUID
	RecordType     string
	Name           string
	Date           time.Time
	Lot            string
	Result         string
	Description    string
	Notes          string
	AdministeredBy uuid.UUID
	VerifiedBy     uuid.UUID
	NextDate       time.Time
}

type RecordUpdateOptions struct {
	Id             uuid.UUID
	RecordType     string
	Name           string
	Date           time.Time
	Lot            string
	Result         string
	Description    string
	Notes          string
	NextDate       time.Time
	VerifiedBy     uuid.UUID
	AdministeredBy uuid.UUID
}

type LoginOptions struct {
	Email    string
	Password string
}

type UserCreateOptions struct {
	UserType string
	Email    string
	Password string
	Name     string
	Surname  string
	Phone    string
	Address  string
	City     string
	State    string
	Country  string
	Zip      string
}

type UserUpdateOptions struct {
	Id      uuid.UUID
	Email   string
	Name    string
	Surname string
	Phone   string
	Address string
	City    string
	State   string
	Country string
	Zip     string
}

func New(opts Options) *Application {
	app := Application{petService: opts.PetService, userService: opts.UserService, recordService: opts.RecordService}

	return &app
}

func (a *Application) CreateUser(opts UserCreateOptions) (user.User, error) {
	u := user.User{}

	typ, err := user.ParseType(opts.UserType)
	if err != nil {
		return user.Nil, user.ErrNoValidType
	}

	err = a.CheckEmail(opts.Email, u.Id, true)
	if err != nil {
		return user.Nil, err
	}

	err = utils.IsPasswordValid(opts.Password)
	if err != nil {
		return user.Nil, err
	}

	hashed, err := authUtils.HashPassword(opts.Password)
	if err != nil {
		return user.Nil, err
	}

	if utils.TextIsEmpty(opts.Name) {
		return user.Nil, user.ErrNoValidName
	}

	if utils.TextIsEmpty(opts.Surname) {
		return user.Nil, user.ErrNoValidSurname
	}

	u.UserType = typ
	u.Email = opts.Email
	u.PasswordHash = hashed
	u.Name = opts.Name
	u.Surname = opts.Surname
	u.Phone = opts.Phone
	u.Address = opts.Address
	u.City = opts.City
	u.State = opts.State
	u.Country = opts.Country
	u.Zip = opts.Zip

	return a.userService.CreateUser(u)
}

func (a *Application) UserToken(u user.User) (string, error) {
	if u.Deleted {
		return "", user.ErrUserDeleted
	}
	return jwtUtils.GenerateJWT(u.Id, u.UserType)
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

func (a *Application) UpdateUser(opts UserUpdateOptions, includeDel bool) (user.User, error) {
	u, err := a.User(opts.Id)
	if err != nil {
		return user.Nil, user.ErrNotFound
	}

	err = a.CheckEmail(opts.Email, u.Id, includeDel)
	if err != nil {
		return user.Nil, err
	}

	if utils.TextIsEmpty(opts.Name) {
		return user.Nil, user.ErrNoValidName
	}

	if utils.TextIsEmpty(opts.Surname) {
		return user.Nil, user.ErrNoValidSurname
	}

	u.Email = opts.Email
	u.Name = opts.Name
	u.Surname = opts.Surname
	u.Phone = opts.Phone
	u.Address = opts.Address
	u.City = opts.City
	u.State = opts.State
	u.Country = opts.Country
	u.Zip = opts.Zip

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

func (a *Application) Authenticate(opts LoginOptions) (user.User, error) {
	u, err := a.userService.Authenticate(opts.Email, opts.Password)
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

func (a *Application) petByOwner(uId uuid.UUID, id uuid.UUID, includeDel bool) (pet.Pet, error) {
	return a.petService.PetByOwner(uId, id, includeDel)
}

func (a *Application) DeletePet(uId uuid.UUID, id uuid.UUID) error {
	_, err := a.petByOwner(uId, id, false)

	if err != nil {
		// if it's not the owner then check if it's the pet vet
		// in this case we don't want to delete the pet but we want to remove the vet
		if err == pet.ErrNotFound {
			return a.removeVet(uId, id)
		}
		return err
	}

	return a.petService.DeletePet(id)
}

func (a *Application) removeVet(uId uuid.UUID, id uuid.UUID) error {
	_, err := a.UserByType(uId, user.Vet, false)
	if err != nil {
		return err
	}
	_, err = a.PetByUser(uId, id, false)
	if err != nil {
		return err
	}
	return a.petService.RemoveVet(id)
}

func (a *Application) CreatePet(opts PetCreateOptions) (pet.Pet, error) {
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

	if utils.TextIsEmpty(opts.Name) {
		return pet.Nil, pet.ErrNoValidName
	}

	if opts.DateOfBirth.IsZero() {
		return pet.Nil, pet.ErrNoValidBirthDate
	}

	gender, err := pet.ParseGender(opts.Gender)
	if err != nil {
		return pet.Nil, err
	}

	if utils.TextIsEmpty(opts.BreedName) {
		return pet.Nil, pet.ErrNoValidBreedname
	}

	p := pet.Pet{}
	p.Name = opts.Name
	p.DateOfBirth = opts.DateOfBirth
	p.Gender = gender
	p.BreedName = opts.BreedName
	p.Colors = opts.Colors
	p.Description = opts.Description
	p.Pedigree = opts.Pedigree
	p.Microchip = opts.Microchip
	p.OwnerId = opts.OwnerId
	p.VetId = opts.VetId
	p.Metas = opts.Metas
	p.Avatar = opts.Avatar
	return a.petService.CreatePet(p)
}

func (a *Application) UpdatePet(opts PetUpdateOptions) (pet.Pet, error) {
	p, err := a.PetByUser(opts.OwnerId, opts.Id, false)
	if err != nil {
		return pet.Nil, err
	}

	if utils.TextIsEmpty(opts.Name) {
		return pet.Nil, pet.ErrNoValidName
	}

	if utils.TextIsEmpty(opts.BreedName) {
		return pet.Nil, pet.ErrNoValidBreedname
	}

	if opts.DateOfBirth.IsZero() {
		return pet.Nil, pet.ErrNoValidBirthDate
	}

	gender, err := pet.ParseGender(opts.Gender)
	if err != nil {
		return pet.Nil, err
	}

	p.Name = opts.Name
	p.DateOfBirth = opts.DateOfBirth
	p.Gender = gender
	p.BreedName = opts.BreedName
	p.Colors = opts.Colors
	p.Description = opts.Description
	p.Pedigree = opts.Pedigree
	p.Microchip = opts.Microchip
	p.VetId = opts.VetId
	p.Metas = opts.Metas
	p.Avatar = opts.Avatar

	return a.petService.UpdatePet(p)
}

func (a *Application) CreateRecord(opts RecordCreateOptions) (record.Record, error) {
	r := record.Nil

	_, err := a.PetByUser(opts.AdministeredBy, opts.PetId, false)
	if err != nil {
		return r, err
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

	typ, err := record.ParseType(opts.RecordType)
	if err != nil {
		return r, record.ErrNotValidType
	}

	if typ == record.Weight || typ == record.Temperature {
		if utils.TextIsEmpty(opts.Result) {
			return record.Nil, record.ErrNotValidResult
		}
		if opts.Date.After(time.Now()) {
			return record.Nil, record.ErrNotValidDate
		}
	} else {
		if utils.TextIsEmpty(opts.Name) {
			return record.Nil, record.ErrNotValidName
		}
	}

	if opts.Date.IsZero() {
		return r, record.ErrNotValidDate
	}

	r.PetId = opts.PetId
	r.RecordType = typ
	r.Name = opts.Name
	r.Date = opts.Date
	r.Lot = opts.Lot
	r.Result = opts.Result
	r.Description = opts.Description
	r.Notes = opts.Notes

	if !opts.Date.After(time.Now()) {
		r.AdministeredBy = opts.AdministeredBy
		r.VerifiedBy = opts.VerifiedBy
	}

	return a.recordService.CreateRecord(r)
}

func (a *Application) CreateRecords(opts RecordsCreateOptions) (map[uuid.UUID]record.Record, error) {
	r := record.Nil

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

	typ, err := record.ParseType(opts.RecordType)
	if err != nil {
		return nil, record.ErrNotValidType
	}

	if typ == record.Weight || typ == record.Temperature {
		if utils.TextIsEmpty(opts.Result) {
			return nil, record.ErrNotValidResult
		}
		if opts.Date.After(time.Now()) {
			return nil, record.ErrNotValidDate
		}
	} else {
		if utils.TextIsEmpty(opts.Name) {
			return nil, record.ErrNotValidName
		}
	}

	if opts.Date.IsZero() {
		return nil, record.ErrNotValidDate
	}

	if opts.NextDate.IsZero() {
		return nil, record.ErrNotValidDate
	}

	recs := make([]record.Record, 2)
	r.PetId = opts.PetId
	r.RecordType = typ
	r.Name = opts.Name
	r.Date = opts.Date
	r.Lot = opts.Lot
	r.Result = opts.Result
	r.Description = opts.Description
	r.Notes = opts.Notes

	if !opts.Date.After(time.Now()) {
		r.AdministeredBy = opts.AdministeredBy
		r.VerifiedBy = opts.VerifiedBy
	}
	recs[0] = r

	nextRecord := record.Record{}
	nextRecord.PetId = opts.PetId
	nextRecord.RecordType = typ
	nextRecord.Name = opts.Name
	nextRecord.Date = opts.NextDate
	recs[1] = nextRecord

	return a.recordService.CreateRecords(recs)
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

func (a *Application) UpdateRecord(opts RecordUpdateOptions) (record.Record, error) {
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

	typ, err := record.ParseType(opts.RecordType)
	if err != nil {
		return record.Nil, record.ErrNotValidType
	}

	if typ == record.Weight || typ == record.Temperature {
		if utils.TextIsEmpty(opts.Result) {
			return record.Nil, record.ErrNotValidResult
		}
		if opts.Date.After(time.Now()) {
			return record.Nil, record.ErrNotValidDate
		}
	} else {
		if utils.TextIsEmpty(opts.Name) {
			return record.Nil, record.ErrNotValidName
		}
	}

	if opts.Date.IsZero() {
		return record.Nil, record.ErrNotValidDate
	}

	r, err := a.recordService.Record(opts.Id)
	if err != nil {
		return record.Nil, err
	}

	r.RecordType = typ
	r.Name = opts.Name
	r.Date = opts.Date
	r.Lot = opts.Lot
	r.Result = opts.Result
	r.Description = opts.Description
	r.Notes = opts.Notes
	r.VerifiedBy = opts.VerifiedBy
	if r.AdministeredBy == uuid.Nil {
		r.AdministeredBy = opts.AdministeredBy
	}
	if opts.Date.After(time.Now()) {
		r.AdministeredBy = uuid.Nil
		r.VerifiedBy = uuid.Nil
	}

	return a.recordService.UpdateRecord(r)
}

func (a *Application) DeleteRecordUserPet(uId uuid.UUID, pId uuid.UUID, id uuid.UUID) error {
	_, err := a.RecordByUserPet(uId, pId, id, false)
	if err != nil {
		return err
	}

	return a.recordService.DeleteRecord(id)
}
