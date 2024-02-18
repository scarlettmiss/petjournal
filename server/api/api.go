package api

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	middlewares2 "github.com/scarlettmiss/petJournal/api/middlewares"
	"github.com/scarlettmiss/petJournal/application"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	"github.com/scarlettmiss/petJournal/application/domain/record"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	"github.com/scarlettmiss/petJournal/utils"
	"net/http"
)

type API struct {
	*gin.Engine
	app *application.Application
}

func New(application *application.Application, ui embed.FS) *API {
	api := &API{
		Engine: gin.Default(),
		app:    application,
	}
	api.NoRoute(middlewares2.NoRouteMiddleware("/", ui, "public"))

	api.POST("/api/auth/register", api.register)
	api.POST("/api/auth/login", api.login)
	api.GET("/api/vets", api.vets)

	userApi := api.Group("/").Use(middlewares2.Auth())
	userApi.GET("/api/users", api.users)
	userApi.GET("/api/user", api.user)
	userApi.GET("/api/user/:id", api.user)
	userApi.PATCH("/api/user", api.updateUser)
	userApi.DELETE("/api/user", api.deleteUser)

	petApi := api.Group("/").Use(middlewares2.Auth())
	petApi.POST("/api/pet", api.createPet)
	petApi.GET("/api/pets", api.pets)
	petApi.GET("/api/pet/:petId", api.pet)
	petApi.PATCH("/api/pet/:petId", api.updatePet)
	petApi.DELETE("/api/pet/:petId", api.deletePet)

	recordApi := api.Group("/").Use(middlewares2.Auth())
	recordApi.POST("/api/pet/:petId/record", api.createRecord)
	recordApi.POST("/api/pet/:petId/records", api.createRecords)
	recordApi.GET("/api/pet/:petId/records", api.recordsByPet)
	recordApi.GET("/api/records", api.records)
	recordApi.GET("/api/pet/:petId/record/:recordId", api.recordByPet)
	recordApi.PATCH("/api/pet/:petId/record/:recordId", api.updateRecord)
	recordApi.DELETE("/api/pet/:petId/record/:recordId", api.deleteRecord)

	return api
}

func (api *API) register(c *gin.Context) {
	var requestBody UserCreateRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	uOpts := UserCreateRequestToUserCreateOptions(requestBody)
	u, err := api.app.CreateUser(uOpts)
	if err != nil {
		switch err {
		case user.ErrNoValidType,
			user.ErrNoValidMail,
			user.ErrMailExists,
			user.ErrNoValidName,
			user.ErrNoValidSurname,
			utils.ErrPasswordLength,
			utils.ErrPasswordLowerCase,
			utils.ErrPasswordUpperCase,
			utils.ErrPasswordDigit,
			utils.ErrPasswordSpecialChar:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	token, err := api.app.UserToken(u)
	if err != nil {
		switch err {
		case user.ErrUserDeleted:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": UserToResponse(u), "token": token})
}

func (api *API) users(c *gin.Context) {
	users, err := api.app.Users(true)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	usersResp := make([]*UserResponse, 0, len(users))
	for _, u := range users {
		userResponse := UserToResponse(u)
		usersResp = append(usersResp, userResponse)
	}

	c.JSON(http.StatusOK, usersResp)
}

func (api *API) vets(c *gin.Context) {
	users, err := api.app.UsersByType(user.Vet, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	usersResp := make([]*UserResponse, 0, len(users))
	for _, u := range users {
		userResponse := UserToResponse(u)
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
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	u, err := api.app.User(uId)
	if err != nil {
		switch err {
		case user.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, UserToResponse(u))
}

func (api *API) updateUser(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var requestBody UserUpdateRequest
	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	uOpts := UserUpdateRequestToUserOptions(requestBody, uId)

	u, err := api.app.UpdateUser(uOpts, false)
	if err != nil {
		switch err {
		case user.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		case user.ErrNoValidType,
			user.ErrNoValidMail,
			user.ErrMailExists,
			user.ErrNoValidName,
			user.ErrNoValidSurname:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	c.JSON(http.StatusOK, UserToResponse(u))
}

func (api *API) deleteUser(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	err = api.app.DeleteUser(uId)
	if err != nil {
		switch err {
		case user.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (api *API) login(c *gin.Context) {
	var requestBody LoginRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	loginOpts := application.LoginOptions{Email: requestBody.Email, Password: requestBody.Password}

	u, err := api.app.Authenticate(loginOpts)
	if err != nil {
		if err == user.ErrUserDeleted {
			c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token, err := api.app.UserToken(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": UserToResponse(u), "token": token})

}

func (api *API) createPet(c *gin.Context) {
	var requestBody PetCreateRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	opts, err := PetCreateRequestToPetCreateOpts(requestBody, uId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	p, err := api.app.CreatePet(opts)
	if err != nil {
		switch err {
		case pet.ErrNoValidName, pet.ErrNoValidBreedname, pet.ErrNoValidBirthDate:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	owner, vet, err := api.ownerVetResponse(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})

		return
	}

	c.JSON(http.StatusCreated, PetToResponse(p, owner, vet))
}

func (api *API) pets(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	pets, err := api.app.PetsByUser(uId, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	petsResp := make([]PetResponse, 0, len(pets))
	var hasError bool
	for _, p := range pets {
		owner, vet, err := api.ownerVetResponse(p)
		if err != nil {
			hasError = true
		}

		petResponse := PetToResponse(p, owner, vet)
		petsResp = append(petsResp, petResponse)
	}

	if hasError {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, petsResp)
}

func (api *API) pet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
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
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	p, err := api.app.PetByUser(uId, pId, false)
	if err != nil {
		switch err {
		case pet.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	owner, vet, err := api.ownerVetResponse(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, PetToResponse(p, owner, vet))
}

func (api *API) updatePet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	id := c.Param("petId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	pId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var requestBody PetUpdateRequest
	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	opts, err := PetUpdateRequestToPetUpdateOpts(requestBody, pId, uId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	p, err := api.app.UpdatePet(opts)
	if err != nil {
		switch err {
		case pet.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		case pet.ErrNoValidName,
			pet.ErrNoValidBreedname,
			pet.ErrNoValidBirthDate:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	owner, vet, err := api.ownerVetResponse(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, PetToResponse(p, owner, vet))
}

func (api *API) ownerVetResponse(p pet.Pet) (user.User, user.User, error) {
	owner, err := api.app.User(p.OwnerId)
	if err != nil {
		return user.Nil, user.Nil, err
	}

	vet, err := api.app.UserByType(p.VetId, user.Vet, false)
	if err != nil && err != user.ErrNotFound {
		return user.Nil, user.Nil, err
	}

	return owner, vet, nil
}

func (api *API) deletePet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	id := c.Param("petId")
	pId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	err = api.app.DeletePet(uId, pId)
	if err != nil {
		switch err {
		case user.ErrNotFound, pet.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pet deleted"})
}

func (api *API) createRecord(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")
	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	u, err := api.app.User(uId)
	if err != nil {
		switch err {
		case user.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	var requestBody RecordCreateRequest
	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	opts := RecordCreateRequestToRecord(requestBody, petId, u)
	r, err := api.app.CreateRecord(opts)
	if err != nil {
		switch err {
		case record.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		case pet.ErrNotFound, record.ErrNotValidType, record.ErrNotValidResult,
			record.ErrNotValidName, record.ErrNotValidDate, record.ErrNotValidVerifier:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
	}

	p, err := api.app.Pet(r.PetId)
	if err != nil {
		switch err {
		case pet.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	administer, err := api.app.User(r.AdministeredBy)
	if err != nil && err != user.ErrNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	verifier, err := api.app.User(r.VerifiedBy)
	if err != nil && err != user.ErrNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, RecordToResponse(r, p, administer, verifier))
}

func (api *API) recordsToRecordsResponse(records map[uuid.UUID]record.Record) ([]RecordResponse, bool) {
	hasError := false
	recordsResp := make([]RecordResponse, 0, len(records))
	for _, r := range records {
		administer, err := api.app.User(r.AdministeredBy)
		if err != nil && err != user.ErrNotFound {
			hasError = true
		}

		verifier, err := api.app.User(r.VerifiedBy)
		if err != nil && err != user.ErrNotFound {
			hasError = true
		}

		p, err := api.app.Pet(r.PetId)
		if err != nil {
			hasError = true
		}

		recordResponse := RecordToResponse(r, p, administer, verifier)
		recordsResp = append(recordsResp, recordResponse)
	}

	return recordsResp, !hasError
}

func (api *API) createRecords(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")
	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	u, err := api.app.User(uId)
	if err != nil {
		switch err {
		case user.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	var requestBody RecordCreateRequest
	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	opts := RecordsCreateRequestToRecord(requestBody, petId, u)
	records, err := api.app.CreateRecords(opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	recordsResp, ok := api.recordsToRecordsResponse(records)
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, recordsResp)
}

func (api *API) recordsByPet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")
	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	records, err := api.app.RecordsByUserPet(uId, petId, false)
	if err != nil {
		switch err {
		case pet.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	recordsResp, ok := api.recordsToRecordsResponse(records)
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, recordsResp)
}

func (api *API) records(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	records, err := api.app.RecordsByUser(uId, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	recordsResp, ok := api.recordsToRecordsResponse(records)
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, recordsResp)
}

func (api *API) recordByPet(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")
	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	rId := c.Param("recordId")
	recordId, err := uuid.Parse(rId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	r, err := api.app.RecordByUserPet(uId, petId, recordId, false)
	if err != nil {
		switch err {
		case pet.ErrNotFound, record.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	p, err := api.app.Pet(r.PetId)
	if err != nil {
		switch err {
		case pet.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	administer, err := api.app.User(r.AdministeredBy)
	if err != nil && err != user.ErrNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	verifier, err := api.app.User(r.VerifiedBy)
	if err != nil && err != user.ErrNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, RecordToResponse(r, p, administer, verifier))
}

func (api *API) updateRecord(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	rId := c.Param("recordId")
	recordId, err := uuid.Parse(rId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var requestBody RecordUpdateRequest
	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	u, err := api.app.User(uId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	opts := RecordUpdateRequestToRecord(requestBody, recordId, u)
	r, err := api.app.UpdateRecord(opts)
	if err != nil {
		switch err {
		case record.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		case record.ErrNotValidType, record.ErrNotValidResult,
			record.ErrNotValidName, record.ErrNotValidDate, record.ErrNotValidVerifier:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	p, err := api.app.Pet(r.PetId)
	if err != nil {
		switch err {
		case pet.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	administer, err := api.app.User(r.AdministeredBy)
	if err != nil && err != user.ErrNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	verifier, err := api.app.User(r.VerifiedBy)
	if err != nil && err != user.ErrNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, RecordToResponse(r, p, administer, verifier))
}

func (api *API) deleteRecord(c *gin.Context) {
	uId, err := uuid.Parse(c.GetString("UserId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	rId := c.Param("recordId")
	recordId, err := uuid.Parse(rId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	pId := c.Param("petId")
	petId, err := uuid.Parse(pId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	err = api.app.DeleteRecordUserPet(uId, petId, recordId)
	if err != nil {
		switch err {
		case pet.ErrNotFound, user.ErrNotFound, record.ErrNotFound:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}
