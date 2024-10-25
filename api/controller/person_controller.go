package controller

import (
	"github.com/Efamamo/GoCrudChallange/api/dtos"
	"github.com/Efamamo/GoCrudChallange/application/people/command"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/Efamamo/GoCrudChallange/internal/cqrs"
	"github.com/gin-gonic/gin"
)

type PersonController struct {
	CreateHandler cqrs.Handler[*command.CreatePersonCommand, *model.Person]
}

func (pc *PersonController) Create(c *gin.Context) {
	var dto dtos.CreatePersonDTO
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

	response := dtos.PersonResponseDTO{
		ID:      p.Id(),
		Name:    p.Name(),
		Age:     p.Age(),
		Hobbies: p.Hobbies(),
	}

	c.IndentedJSON(201, response)

}
