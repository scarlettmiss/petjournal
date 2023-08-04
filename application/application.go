package application

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/bestPal/application/domain/pet"
	"github.com/scarlettmiss/bestPal/application/domain/treatment"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	jwtService "github.com/scarlettmiss/bestPal/application/services/jwtService"
	petService "github.com/scarlettmiss/bestPal/application/services/petService"
	treatmentService "github.com/scarlettmiss/bestPal/application/services/treatmentService"
	userService "github.com/scarlettmiss/bestPal/application/services/userService"
)

/*
what the actor can do.
application talks with all the services
*/
type Application struct {
	petService       petService.Service
	userService      userService.Service
	treatmentService treatmentService.Service
}

type Options struct {
	PetService       petService.Service
	UserService      userService.Service
	TreatmentService treatmentService.Service
}

func New(opts Options) *Application {
	app := Application{petService: opts.PetService, userService: opts.UserService, treatmentService: opts.TreatmentService}

	return &app
}

func (a *Application) CreateUser(u user.User) (user.User, error) {
	err := a.CheckEmail(u.Email, u.Id)
	if err != nil {
		return user.Nil, err
	}

	return a.userService.CreateUser(u)
}

func (a *Application) UserToken(u user.User) (string, error) {
	return jwtService.GenerateJWT(u.Id, u.UserType)
}

func (a *Application) CheckEmail(email string, id uuid.UUID) error {
	u, ok := a.userService.UserByEmail(email)

	if !ok {
		return nil
	}

	if u.Id == id {
		return nil
	}

	return user.ErrMailExists
}

func (a *Application) UpdateUser(u user.User) (user.User, error) {
	err := a.CheckEmail(u.Email, u.Id)
	if err != nil {
		return user.Nil, err
	}
	return a.userService.UpdateUser(u)
}

func (a *Application) Users() ([]user.User, error) {
	return a.userService.Users()
}

func (a *Application) UsersByType(t user.Type) ([]user.User, error) {
	return a.userService.UsersByType(t)
}

func (a *Application) User(id uuid.UUID) (user.User, error) {
	return a.userService.User(id)
}

func (a *Application) UserByType(id uuid.UUID, t user.Type) (user.User, error) {
	u, err := a.userService.UserByType(id, t)
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

func (a *Application) PetsByUser(uId uuid.UUID) (map[uuid.UUID]pet.Pet, error) {
	return a.petService.PetsByUser(uId)
}

func (a *Application) Pet(id uuid.UUID) (pet.Pet, error) {
	return a.petService.Pet(id)
}

func (a *Application) PetByUser(uId uuid.UUID, id uuid.UUID) (pet.Pet, error) {
	return a.petService.PetByUser(uId, id)
}

func (a *Application) DeletePet(id uuid.UUID) error {
	return a.petService.DeletePet(id)
}

func (a *Application) CreatePet(p pet.Pet) (pet.Pet, error) {
	return a.petService.CreatePet(p)
}

func (a *Application) UpdatePet(p pet.Pet) (pet.Pet, error) {
	return a.petService.UpdatePet(p)
}

func (a *Application) CreateTreatment(t treatment.Treatment) (treatment.Treatment, error) {
	return a.treatmentService.CreateTreatment(t)
}

func (a *Application) TreatmentsByUser(uId uuid.UUID) (map[uuid.UUID]treatment.Treatment, error) {
	pets, err := a.PetsByUser(uId)
	if err != nil {
		return nil, err
	}
	return a.treatmentService.PetsTreatments(lo.Keys[uuid.UUID, pet.Pet](pets))
}

func (a *Application) TreatmentsByPet(pId uuid.UUID) (map[uuid.UUID]treatment.Treatment, error) {
	return a.treatmentService.PetTreatments(pId)
}

func (a *Application) TreatmentByPet(pId uuid.UUID, tId uuid.UUID) (treatment.Treatment, error) {
	return a.treatmentService.PetTreatment(pId, tId)
}

func (a *Application) UpdateTreatment(t treatment.Treatment) (treatment.Treatment, error) {
	return a.treatmentService.UpdateTreatment(t)
}

func (a *Application) DeleteTreatment(id uuid.UUID) error {
	return a.treatmentService.DeleteTreatment(id)
}
