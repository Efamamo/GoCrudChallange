package command

import (
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	ierr "github.com/Efamamo/GoCrudChallange/domain/error"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
)

// CreatePersonCommand holds the data required to create a new Person entity.
type CreatePersonCommand struct {
	Name    string
	Age     int16
	Hobbies []string
}

// CreatePersonHandler is responsible for handling the logic of creating a new Person entity.
type CreatePersonHandler struct {
	repo irepo.IPerson
}

// Compile-time check to ensure CreatePersonHandler implements IHandler for CreatePersonCommand.
var _ icmd.IHandler[*CreatePersonCommand, *model.Person] = &CreatePersonHandler{}

// NewCreatePersonHandler initializes a new CreatePersonHandler with a given IPerson repository.
func NewCreatePersonHandler(repo irepo.IPerson) *CreatePersonHandler {
	return &CreatePersonHandler{repo: repo}
}

// Handle processes the CreatePersonCommand to create a new Person entity.
func (h *CreatePersonHandler) Handle(command *CreatePersonCommand) (*model.Person, ierr.IErr) {
	person, err := model.CreatePerson(&model.PersonConfig{
		Name:    command.Name,
		Age:     command.Age,
		Hobbies: command.Hobbies,
	})
	if err != nil {
		return nil, err
	}

	if err := h.repo.Save(person); err != nil {
		return nil, err
	}

	return person, nil
}
