package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/scarlettmiss/bestPal/application"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"github.com/scarlettmiss/bestPal/cmd/server/types"
	"github.com/scarlettmiss/bestPal/middlewares"
	"github.com/scarlettmiss/bestPal/utils"
	"net/http"
)

type API struct {
	*gin.Engine
	app *application.Application
}

func New(application *application.Application) *API {
	api := &API{
		Engine: gin.Default(),
		app:    application,
	}

	api.NoRoute(func(ctx *gin.Context) { ctx.Status(http.StatusNotFound) })

	api.POST("/api/auth/register", api.register)
	api.POST("/api/auth/login", api.login)

	protected := api.Group("/").Use(middlewares.Auth())
	protected.GET("/api/users", api.users)

	return api
}

func (api *API) register(c *gin.Context) {
	var requestBody types.Account

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	typ, err := user.ParseType(requestBody.UserType)
	if err != nil {
		fmt.Println(err)
	}
	u := user.User{}
	u.UserType = typ
	u.Email = requestBody.Email
	hashed, err := utils.HashPassword(requestBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	u.PasswordHash = hashed
	u.Name = requestBody.Name
	u.Surname = requestBody.Surname
	u.Phone = requestBody.Phone
	u.Address = requestBody.Address
	u.City = requestBody.City
	u.State = requestBody.State
	u.Country = requestBody.Country
	u.Zip = requestBody.Zip

	u, err = api.app.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, u)
	fmt.Println(u)

}

func (api *API) login(c *gin.Context) {
	var requestBody types.LoginRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	token, err := api.app.Authenticate(requestBody.Email, requestBody.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (api *API) users(c *gin.Context) {
	c.JSON(http.StatusOK, api.app.Users())
}
