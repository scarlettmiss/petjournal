package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/bestPal/application"
	"github.com/scarlettmiss/bestPal/application/domain/pet"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	pet2 "github.com/scarlettmiss/bestPal/cmd/server/types/pet"
	user2 "github.com/scarlettmiss/bestPal/cmd/server/types/user"
	"github.com/scarlettmiss/bestPal/converters/petConverter"
	"github.com/scarlettmiss/bestPal/converters/userConverter"
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

	userApi := api.Group("/").Use(middlewares.Auth())
	userApi.GET("/api/users", api.users)
	userApi.GET("/api/user", api.user)
	userApi.GET("/api/user/:id", api.user)
	userApi.PATCH("/api/user", api.updateUser)
	userApi.DELETE("/api/user", api.deleteUser)

	petApi := api.Group("/").Use(middlewares.Auth())
	petApi.POST("/api/pet")
	petApi.GET("/api/pets", api.pets)
	petApi.GET("/api/pet/:id", api.pet)
	petApi.PATCH("/api/pet")
	petApi.DELETE("/api/pet/:id", api.deleteUser)

	return api
}

func (api *API) register(c *gin.Context) {
	var requestBody user2.UserCreateRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	u, err := userConverter.UserCreateRequestToUser(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	u, err = api.app.CreateUser(u)
	if err != nil {
		if err == user.ErrMailExists {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	token, err := api.app.UserToken(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": userConverter.UserToResponse(u), "token": token})
	fmt.Println(u)

}

func (api *API) login(c *gin.Context) {
	var requestBody user2.LoginRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	u, err := api.app.Authenticate(requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token, err := api.app.UserToken(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": userConverter.UserToResponse(u), "token": token})

}

func (api *API) users(c *gin.Context) {
	users := api.app.Users()

	usersResp := lo.MapValues(users, func(u user.User, _ uuid.UUID) user2.UserResponse {
		return userConverter.UserToResponse(u)
	})

	c.JSON(http.StatusOK, usersResp)
}

func (api *API) user(c *gin.Context) {
	id := c.Param("id")

	var ok bool
	if id == "" {
		idParam, _ := c.Get("UserId")
		// Convert UserId to string if it's not already.
		id, ok = idParam.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	//parse
	uId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	u, err := api.app.User(uId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, userConverter.UserToResponse(u))
}

func (api *API) deleteUser(c *gin.Context) {
	idParam, _ := c.Get("UserId")
	// Convert UserId to string if it's not already.
	id, ok := idParam.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	//parse
	uId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	err = api.app.DeleteUser(uId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (api *API) updateUser(c *gin.Context) {
	id, _ := c.Get("UserId")
	// Convert UserId to string if it's not already.
	idString, ok := id.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	uId, err := uuid.Parse(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	var requestBody user2.UserUpdateRequest
	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	u, err := api.app.User(uId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	u = userConverter.UserUpdateRequestToUser(requestBody, u)

	u, err = api.app.UpdateUser(u)
	if err != nil {
		if err == user.ErrMailExists {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	c.JSON(http.StatusOK, userConverter.UserToResponse(u))
}

func (api *API) pets(c *gin.Context) {
	pets := api.app.Pets()

	petsResp := lo.MapValues(pets, func(p pet.Pet, _ uuid.UUID) pet2.PetResponse {
		return petConverter.PetToResponse(p)
	})

	c.JSON(http.StatusOK, petsResp)
}

func (api *API) pet(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	//parse
	pId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	//TODO check if pet belongs to user
	p, err := api.app.Pet(pId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, petConverter.PetToResponse(p))
}

func (api *API) deletePet(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	//parse
	pId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	//TODO check if pet belongs to user
	err = api.app.DeletePet(pId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pet deleted"})
}
