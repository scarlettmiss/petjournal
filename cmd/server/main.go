package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scarlettmiss/bestPal/application"
	"github.com/scarlettmiss/bestPal/application/repositories/baseRepo"
	service "github.com/scarlettmiss/bestPal/application/services/baseService"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "ui/home.html")
}

func main() {
	//init db
	//init repos
	repo := baseRepo.New()
	//init services
	baseService, err := service.New(repo)
	if err != nil {
		panic(err)
	}
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()
	//pass services to application
	_, err = application.New(application.Options{baseService})
	if err != nil {
		panic(err)
	}

	//serve home
	router.GET("/", func(ctx *gin.Context) {
		serveHome(ctx.Writer, ctx.Request)
	})
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	// Start listening and serving requests
	err = router.Run(":8080")

	if err != nil {
		panic(err)
	}

	//ctrl + c to stop server
	waitForInterrupt := make(chan os.Signal, 1)
	signal.Notify(waitForInterrupt, os.Interrupt, os.Kill)

	<-waitForInterrupt
}
