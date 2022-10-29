package application

import service "github.com/scarlettmiss/bestPal/application/services/baseService"

/*
*
what the user can do.
application talks with all the services
*/
type Application struct {
	service *service.Service
}

type Options struct {
	Service *service.Service
}

func New(opts Options) (*Application, error) {
	app := Application{service: opts.Service}

	return &app, nil
}
