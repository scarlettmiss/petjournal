package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	typesPet "github.com/scarlettmiss/bestPal/cmd/server/types/pet"
	typesTreatment "github.com/scarlettmiss/bestPal/cmd/server/types/treatment"
	typesUser "github.com/scarlettmiss/bestPal/cmd/server/types/user"
	"github.com/scarlettmiss/bestPal/converters/petConverter"
	"github.com/scarlettmiss/bestPal/converters/treatmentConverter"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	api.Use(cors.New(config))

	api.NoRoute(func(ctx *gin.Context) { ctx.Status(http.StatusNotFound) })

	api.POST("/api/auth/register", api.register)
	api.POST("/api/auth/login", api.login)
	api.GET("/api/petPage/:petId", api.petSimplified)
	api.GET("/api/vets", api.vets)

	userApi := api.Group("/").Use(middlewares.Auth())
	userApi.GET("/api/users", api.users)
	userApi.GET("/api/user", api.user)
	userApi.GET("/api/user/:id", api.user)
	userApi.PATCH("/api/user", api.updateUser)
	userApi.DELETE("/api/user", api.deleteUser)

	petApi := api.Group("/").Use(middlewares.Auth())
	petApi.POST("/api/pet", api.createPet)
	petApi.GET("/api/pets", api.pets)
	petApi.GET("/api/pet/:petId", api.pet)
	petApi.PATCH("/api/pet/:petId", api.updatePet)
	petApi.DELETE("/api/pet/:petId", api.deletePet)

	treatmentApi := api.Group("/").Use(middlewares.Auth())
	treatmentApi.POST("/api/pet/:petId/treatment", api.createTreatment)
	treatmentApi.GET("/api/pet/:petId/treatments", api.treatmentsByPet)
	treatmentApi.GET("/api/pet/:petId/treatment/:treatmentId", api.treatmentByPet)
	treatmentApi.PATCH("/api/pet/:petId/treatment/:treatmentId", api.updateTreatment)
	treatmentApi.DELETE("/api/pet/:petId/treatment/:treatmentId", api.deleteTreatment)

	return api
}

func (api *API) register(c *gin.Context) {
	var requestBody typesUser.UserCreateRequest

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
}

func (api *API) users(c *gin.Context) {
	users, err := api.app.Users()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	usersResp := make([]typesUser.UserResponse, 0, len(users))
	for _, u := range users {
		userResponse := userConverter.UserToResponse(u)
		usersResp = append(usersResp, userResponse)
	}

	c.JSON(http.StatusOK, usersResp)
}

func (api *API) vets(c *gin.Context) {
	users, err := api.app.UsersByType(user.Vet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	usersResp := make([]typesUser.UserResponse, 0, len(users))
	for _, u := range users {
		userResponse := userConverter.UserToResponse(u)
		usersResp = append(usersResp, userResponse)
	}

	c.JSON(http.StatusOK, usersResp)
}

func (api *API) user(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		id = c.GetString("UserId")
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

func (api *API) updateUser(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	var requestBody typesUser.UserUpdateRequest
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

func (api *API) deleteUser(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
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

func (api *API) login(c *gin.Context) {
	var requestBody typesUser.LoginRequest

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

	c.JSON(http.StatusOK, gin.H{"user": userConverter.UserToResponse(u), "token": token})

}

func (api *API) createPet(c *gin.Context) {
	var requestBody typesPet.PetCreateRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	owner, err := api.app.User(uId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	vet := user.Nil
	vId := uuid.Nil
	if requestBody.VetId != "" {
		vId, err = uuid.Parse(requestBody.VetId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		vet, err = api.app.UserByType(vId, user.Vet)
		if err != nil {
			if err == user.ErrNotFound {
				c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
				return
			}
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
	}

	p, err := petConverter.PetCreateRequestToPet(requestBody, owner.Id, vId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	p, err = api.app.CreatePet(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"pet": petConverter.PetToResponse(p, owner, vet)})
}

func (api *API) pets(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	pets, err := api.app.PetsByUser(uId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	petsResp := make([]typesPet.PetResponse, 0, len(pets))
	var hasError bool
	for _, p := range pets {
		owner, err := api.app.User(p.OwnerId)
		if err != nil {
			hasError = true
		}
		vet := user.Nil
		if p.VetId != uuid.Nil {
			vet, err = api.app.UserByType(p.VetId, user.Vet)
			if err != nil {
				hasError = true
			}
		}

		petResponse := petConverter.PetToResponse(p, owner, vet)
		petsResp = append(petsResp, petResponse)
	}

	if hasError {
		if err == user.ErrNotFound {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, petsResp)
}

func (api *API) petSimplified(c *gin.Context) {
	id := c.Param("petId")

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

	p, err := api.app.Pet(pId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	owner, err := api.app.User(p.OwnerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	vet := user.Nil
	if p.VetId != uuid.Nil {
		vet, err = api.app.User(p.VetId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
	}

	c.JSON(http.StatusOK, petConverter.PetToSimplifiedResponse(p, owner, vet))
}

func (api *API) pet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	id := c.Param("petId")

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
	p, err := api.app.PetByUser(uId, pId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	owner, err := api.app.User(p.OwnerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	vet := user.Nil
	if p.VetId != uuid.Nil {
		vet, err = api.app.User(p.VetId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
	}

	c.JSON(http.StatusOK, petConverter.PetToResponse(p, owner, vet))
}

func (api *API) updatePet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	id := c.Param("petId")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	var requestBody typesPet.PetUpdateRequest

	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	pId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	p, err := api.app.PetByUser(uId, pId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	vet := user.Nil
	vId := uuid.Nil
	if requestBody.VetId != "" {
		vId, err = uuid.Parse(requestBody.VetId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		vet, err = api.app.UserByType(vId, user.Vet)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}
	}

	p, err = petConverter.PetUpdateRequestToPet(requestBody, p, vId)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	p, err = api.app.UpdatePet(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	owner, err := api.app.User(p.OwnerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"pet": petConverter.PetToResponse(p, owner, vet)})
}

func (api *API) deletePet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	id := c.Param("petId")
	//parse
	pId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	_, err = api.app.PetByUser(uId, pId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	err = api.app.DeletePet(pId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pet deleted"})
}

func (api *API) createTreatment(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	u, err := api.app.User(uId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")

	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	_, err = api.app.PetByUser(uId, petId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	var requestBody typesTreatment.TreatmentCreateRequest

	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	t, err := treatmentConverter.TreatmentCreateRequestToTreatment(requestBody, petId, uId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var verifier = user.Nil
	if u.UserType == user.Vet {
		t.VerifiedBy = u.Id
		verifier = u
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	t, err = api.app.CreateTreatment(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"treatment": treatmentConverter.TreatmentToResponse(t, u, verifier)})
}

func (api *API) treatmentsByPet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")

	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	_, err = api.app.PetByUser(uId, petId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	treatments, err := api.app.TreatmentsByPet(petId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	var hasError bool
	treatmentsResp := make([]typesTreatment.TreatmentResponse, 0, len(treatments))
	for _, t := range treatments {
		administer, err := api.app.User(t.AdministeredBy)
		if err != nil {
			hasError = true
		}

		verifier := user.Nil
		if t.VerifiedBy != uuid.Nil {
			verifier, err = api.app.User(t.VerifiedBy)
			if err != nil {
				hasError = true
			}
		}

		treatmentResponse := treatmentConverter.TreatmentToResponse(t, administer, verifier)
		treatmentsResp = append(treatmentsResp, treatmentResponse)
	}

	if hasError {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, treatmentsResp)
}

func (api *API) treatmentByPet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")

	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	tId := c.Param("treatmentId")

	treatmentId, err := uuid.Parse(tId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	_, err = api.app.PetByUser(uId, petId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	t, err := api.app.TreatmentByPet(petId, treatmentId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	administer, err := api.app.User(t.AdministeredBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	verifier := user.Nil
	if t.VerifiedBy != uuid.Nil {
		verifier, err = api.app.User(t.VerifiedBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
	}

	c.JSON(http.StatusOK, treatmentConverter.TreatmentToResponse(t, administer, verifier))
}

func (api *API) updateTreatment(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	tId := c.Param("treatmentId")
	treatmentId, err := uuid.Parse(tId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	u, err := api.app.User(uId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")

	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	_, err = api.app.PetByUser(uId, petId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	t, err := api.app.TreatmentByPet(petId, treatmentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var requestBody typesTreatment.TreatmentUpdateRequest

	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	t, err = treatmentConverter.TreatmentUpdateRequestToTreatment(requestBody, t)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	var verifier = user.Nil
	if u.UserType == user.Vet {
		t.VerifiedBy = u.Id
		verifier = u
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	t, err = api.app.UpdateTreatment(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"treatment": treatmentConverter.TreatmentToResponse(t, u, verifier)})
}

func (api *API) deleteTreatment(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	tId := c.Param("treatmentId")
	treatmentId, err := uuid.Parse(tId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	_, err = api.app.User(uId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")
	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	_, err = api.app.PetByUser(uId, petId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	_, err = api.app.TreatmentByPet(petId, treatmentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	err = api.app.DeleteTreatment(treatmentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "treatment deleted"})
}
