package command

import (
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	ierr "github.com/Efamamo/GoCrudChallange/domain/common"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

// UpdatePersonCommand represents the command to update a person's details.
type UpdatePersonCommand struct {
	ID      uuid.UUID // Unique identifier of the person.
	Name    string    // New name for the person.
	Age     int16     // New age for the person.
	Hobbies []string  // Updated list of hobbies for the person.
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
// It returns the updated person and an error, if any.
func (h *UpdatePersonHandler) Handle(command *UpdatePersonCommand) (*model.Person, ierr.IErr) {
	// Retrieve the person from the repository using the provided ID.
	person, err := h.repo.Get(command.ID)
	if err != nil {
		return nil, err // Return nil and the error if the person is not found.
	}

	// Attempt to set the new name for the person.
	if err := person.SetName(command.Name); err != nil {
		return nil, err // Return nil and the error if setting the name fails.
	}

	// Attempt to set the new age for the person.
	if err := person.SetAge(command.Age); err != nil {
		return nil, err // Return nil and the error if setting the age fails.
	}

	// Set the hobbies for the person.
	person.SetHobbies(command.Hobbies)

	// Save the updated person back to the repository.
	if err := h.repo.Save(person); err != nil {
		return nil, err // Return nil and the error if saving fails.
	}

	return person, nil // Return the updated person and no error.
}
