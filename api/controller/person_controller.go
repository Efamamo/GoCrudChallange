package controller

import (
	errapi "github.com/Efamamo/GoCrudChallange/api/error"
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	iquery "github.com/Efamamo/GoCrudChallange/application/common/cqrs/query"
	"github.com/Efamamo/GoCrudChallange/application/people/command"
	ierr "github.com/Efamamo/GoCrudChallange/domain/error"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PersonController defines handlers for managing CRUD operations on Person entities.
type PersonController struct {
	CreateHandler icmd.IHandler[*command.CreatePersonCommand, *model.Person]
	UpdateHandler icmd.IHandler[*command.UpdatePersonCommand, *model.Person]
	DeleteHandler icmd.IHandler[uuid.UUID, bool]
	GetHandler    iquery.IHandler[uuid.UUID, *model.Person]
	GetAllHandler iquery.IHandler[struct{}, []*model.Person]
}

// Create handles the creation of a new Person.
// It parses JSON data from the request body, validates it, and calls the CreateHandler to add a new Person.
// Responds with a 201 status code if successful, or 400 if the input is invalid.
func (pc *PersonController) Create(c *gin.Context) {
	var dto CreateDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		e := errapi.NewBadRequest("Invalid input data format")
		c.IndentedJSON(e.StatusCode(), gin.H{"error": e.Error()})
		return
	}

	command := &command.CreatePersonCommand{
		Name:    dto.Name,
		Age:     dto.Age,
		Hobbies: dto.Hobbies,
	}

	// Calls CreateHandler to process the creation command
	p, err := pc.CreateHandler.Handle(command)

	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	// Prepares the response DTO
	response := ResponseDTO{
		ID:      p.Id(),
		Name:    p.Name(),
		Age:     p.Age(),
		Hobbies: p.Hobbies(),
	}

	c.IndentedJSON(201, response)
}

// Update handles updating an existing Person's details.
// It retrieves the ID from the URL, binds JSON data from the request, and calls UpdateHandler to apply changes.
// Returns a 201 status code if successful or relevant errors for invalid input or update issues.
func (pc *PersonController) Update(c *gin.Context) {
	var dto CreateDTO

	// Parse and validate UUID from the URL
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		e := errapi.NewBadRequest("Invalid id format")
		c.IndentedJSON(e.StatusCode(), gin.H{"error": e.Error()})
		return
	}

	// Bind JSON data to the DTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		e := errapi.NewBadRequest(err.Error())
		c.IndentedJSON(e.StatusCode(), gin.H{"error": e.Error()})
		return
	}

	command := &command.UpdatePersonCommand{
		ID:      id,
		Name:    dto.Name,
		Age:     dto.Age,
		Hobbies: dto.Hobbies,
	}

	// Call UpdateHandler to process the update command
	person, err := pc.UpdateHandler.Handle(command)
	if err != nil {
		// Handle custom errors defined by ierr.IErr interface
		if customErr, ok := err.(ierr.IErr); ok {
			e := errapi.Map(customErr)
			c.IndentedJSON(e.StatusCode(), gin.H{"error": e.Error()})
			return
		} else {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	// Prepare response DTO
	response := ResponseDTO{
		ID:      person.Id(),
		Name:    person.Name(),
		Age:     person.Age(),
		Hobbies: person.Hobbies(),
	}

	c.IndentedJSON(201, response)
}

// Delete handles the deletion of a Person by ID.
// It validates the ID, then calls DeleteHandler to remove the specified Person.
// Responds with a 204 status code if successful, or 404 if the Person was not found.
func (pc *PersonController) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		e := errapi.NewBadRequest("Invalid id format")
		c.IndentedJSON(e.StatusCode(), gin.H{"error": e.Error()})
		return
	}

	_, err = pc.DeleteHandler.Handle(id)
	if err != nil {
		// Handle custom error types using ierr.IErr interface mapping
		if customErr, ok := err.(ierr.IErr); ok {
			e := errapi.Map(customErr)
			c.IndentedJSON(e.StatusCode(), gin.H{"error": e.Error()})
			return
		} else {
			c.IndentedJSON(404, gin.H{"error": err.Error()})
			return
		}
	}

	c.IndentedJSON(204, nil)
}

// Get retrieves a Person by their ID.
// It parses the ID from the URL, calls GetHandler to fetch the Person, and returns a 200 status code with Person data if found.
func (pc *PersonController) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		e := errapi.NewBadRequest("Invalid id format")
		c.IndentedJSON(e.StatusCode(), gin.H{"error": e.Error()})
		return
	}

	person, err := pc.GetHandler.Handle(id)
	if err != nil {
		// Map custom errors to HTTP response codes using ierr.IErr interface
		if customErr, ok := err.(ierr.IErr); ok {
			e := errapi.Map(customErr)
			c.IndentedJSON(e.StatusCode(), gin.H{"error": e.Error()})
			return
		} else {
			c.IndentedJSON(404, gin.H{"error": err.Error()})
			return
		}
	}

	// Prepare response DTO
	response := ResponseDTO{
		ID:      person.Id(),
		Name:    person.Name(),
		Age:     person.Age(),
		Hobbies: person.Hobbies(),
	}

	c.IndentedJSON(200, response)
}

// GetAll retrieves all Person entities.
// It calls GetAllHandler to fetch all Persons and returns them as a JSON array with a 200 status code.
func (pc *PersonController) GetAll(c *gin.Context) {
	persons, _ := pc.GetAllHandler.Handle(struct{}{})

	// Prepare response list by mapping each Person to ResponseDTO
	var responses = make([]ResponseDTO, 0)
	for _, person := range persons {
		responses = append(responses, ResponseDTO{
			ID:      person.Id(),
			Name:    person.Name(),
			Age:     person.Age(),
			Hobbies: person.Hobbies(),
		})
	}

	

	c.IndentedJSON(200, responses)
}
