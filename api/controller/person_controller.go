package controller

import (
	"github.com/Efamamo/GoCrudChallange/api/dtos"
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	iquery "github.com/Efamamo/GoCrudChallange/application/common/cqrs/query"
	"github.com/Efamamo/GoCrudChallange/application/people/command"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PersonController struct {
	CreateHandler icmd.IHandler[*command.CreatePersonCommand, *model.Person]
	UpdateHandler icmd.IHandler[*command.UpdatePersonCommand, *model.Person]
	DeleteHandler icmd.IHandler[uuid.UUID, bool]
	GetHandler    iquery.IHandler[uuid.UUID, *model.Person]
	GetAllHandler iquery.IHandler[struct{}, []*model.Person]
}

func (pc *PersonController) Create(c *gin.Context) {
	var dto dtos.CreateDTO
	err := c.ShouldBindJSON(&dto)

	if err != nil {
		c.IndentedJSON(400, err.Error())
		return
	}

	command := &command.CreatePersonCommand{
		Name:    dto.Name,
		Age:     dto.Age,
		Hobbies: dto.Hobbies,
	}

	p, err := pc.CreateHandler.Handle(command)

	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dtos.ResponseDTO{
		ID:      p.Id(),
		Name:    p.Name(),
		Age:     p.Age(),
		Hobbies: p.Hobbies(),
	}

	c.IndentedJSON(201, response)

}

func (pc *PersonController) Update(c *gin.Context) {
	var dto dtos.CreateDTO
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid Id Format"})
		return
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.IndentedJSON(400, err.Error())
		return
	}

	command := &command.UpdatePersonCommand{
		ID:      id,
		Name:    dto.Name,
		Age:     dto.Age,
		Hobbies: dto.Hobbies,
	}

	person, err := pc.UpdateHandler.Handle(command)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dtos.ResponseDTO{
		ID:      person.Id(),
		Name:    person.Name(),
		Age:     person.Age(),
		Hobbies: person.Hobbies(),
	}

	c.IndentedJSON(201, response)
}

func (pc *PersonController) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid Id Format"})
		return
	}

	_, err = pc.DeleteHandler.Handle(id)

	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(204, nil)
}

func (pc *PersonController) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid Id Format"})
		return
	}

	person, err := pc.GetHandler.Handle(id)
	if err != nil {
		c.IndentedJSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dtos.ResponseDTO{
		ID:      person.Id(),
		Name:    person.Name(),
		Age:     person.Age(),
		Hobbies: person.Hobbies(),
	}

	c.IndentedJSON(200, response)
}

func (pc *PersonController) GetAll(c *gin.Context) {
	persons, err := pc.GetAllHandler.Handle(struct{}{})
	if err != nil {
		c.IndentedJSON(500, err)
		return
	}

	var responses = make([]dtos.ResponseDTO, 0)
	for _, person := range persons {
		responses = append(responses, dtos.ResponseDTO{
			ID:      person.Id(),
			Name:    person.Name(),
			Age:     person.Age(),
			Hobbies: person.Hobbies(),
		})
	}

	c.IndentedJSON(200, responses)
}
