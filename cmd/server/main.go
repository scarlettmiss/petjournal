package main

import (
	"github.com/scarlettmiss/bestPal/api"
	"github.com/scarlettmiss/bestPal/api/config"
	"github.com/scarlettmiss/bestPal/application"
	petService "github.com/scarlettmiss/bestPal/application/services/petService"
	treatmentService "github.com/scarlettmiss/bestPal/application/services/treatmentService"
	userService "github.com/scarlettmiss/bestPal/application/services/userService"
	"github.com/scarlettmiss/bestPal/repositories/petrepo"
	"github.com/scarlettmiss/bestPal/repositories/treatmentrepo"
	"github.com/scarlettmiss/bestPal/repositories/userrepo"
	"os"
	"os/signal"
)

func main() {
	//init db
	//init repos
	petRepo := petrepo.New()
	userRepo := userrepo.New()
	treatmentRepo := treatmentrepo.New()
	//init services
	ps, err := petService.New(petRepo)
	if err != nil {
		panic(err)
	}
	us, err := userService.New(userRepo)
	if err != nil {
		panic(err)
	}
	ts, err := treatmentService.New(treatmentRepo)
	if err != nil {
		panic(err)
	}

	//pass services to application
	opts := application.Options{PetService: ps, UserService: us, TreatmentService: ts}
	app := application.New(opts)

	if err != nil {
		panic(err)
	}

	restServer := api.New(app)

	go func() { // Start listening and serving requests
		err = restServer.Run(config.Host + ":" + config.Port)

		if err != nil {
			panic(err)
		}
	}()
	//ctrl + c to stop server
	waitForInterrupt := make(chan os.Signal, 1)
	signal.Notify(waitForInterrupt, os.Interrupt, os.Kill)

	<-waitForInterrupt
}
