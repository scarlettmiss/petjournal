package application

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	petService "github.com/scarlettmiss/bestPal/application/services/petService"
	treatmentService "github.com/scarlettmiss/bestPal/application/services/treatmentService"
	userService "github.com/scarlettmiss/bestPal/application/services/userService"
	"github.com/scarlettmiss/bestPal/utils"
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

func (a *Application) CreateUser(u user.User) (string, error) {
	u, err := a.userService.CreateUser(u)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(u.Id, u.UserType)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Application) CheckEmail(email string) error {
	_, ok := a.userService.UserByEmail(email)
	if !ok {
		return nil
	}
	return user.ErrMailExists
}

func (a *Application) UpdateUser(u user.User) (user.User, error) {
	return a.userService.UpdateUser(u)
}

func (a *Application) Users() map[uuid.UUID]user.User {
	return a.userService.Users()
}

func (a *Application) User(id uuid.UUID) (user.User, error) {
	return a.userService.User(id)
}

func (a *Application) Authenticate(email string, password string) (string, error) {
	u, err := a.userService.Authenticate(email, password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(u.Id, u.UserType)
	if err != nil {
		return "", err
	}

	return token, nil
}
