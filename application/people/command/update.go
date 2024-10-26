package command

import (
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	ierr "github.com/Efamamo/GoCrudChallange/domain/error"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

// UpdatePersonCommand represents the command to update a person's details.
type UpdatePersonCommand struct {
	ID      uuid.UUID
	Name    string
	Age     int16
	Hobbies []string
}

// UpdatePersonHandler is a command handler for updating a person's information.
type UpdatePersonHandler struct {
	repo irepo.IPerson // Repository interface for person operations.
}

// Ensure UpdatePersonHandler implements the IHandler interface for handling commands.
var _ icmd.IHandler[*UpdatePersonCommand, *model.Person] = &UpdatePersonHandler{}

// NewUpdatePersonHandler creates a new instance of UpdatePersonHandler with the provided repository.
func NewUpdatePersonHandler(repo irepo.IPerson) *UpdatePersonHandler {
	return &UpdatePersonHandler{repo: repo}
}

// Handle processes the command to update a person's information.
func (h *UpdatePersonHandler) Handle(command *UpdatePersonCommand) (*model.Person, ierr.IErr) {
	person, err := h.repo.Get(command.ID)
	if err != nil {
		return nil, err
	}

	if err := person.SetName(command.Name); err != nil {
		return nil, err
	}

	if err := person.SetAge(command.Age); err != nil {
		return nil, err
	}

	person.SetHobbies(command.Hobbies)

	if err := h.repo.Save(person); err != nil {
		return nil, err
	}

	return person, nil
}
