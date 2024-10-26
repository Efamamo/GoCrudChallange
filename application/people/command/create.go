package command

import (
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	ierr "github.com/Efamamo/GoCrudChallange/domain/common"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
)

// CreatePersonCommand holds the data required to create a new Person entity.
// It is used as an input to the CreatePersonHandler to execute the creation logic.
type CreatePersonCommand struct {
	Name    string   // Name of the person to create
	Age     int16    // Age of the person to create
	Hobbies []string // List of hobbies associated with the person
}

// CreatePersonHandler is responsible for handling the logic of creating a new Person entity.
// It utilizes an IPerson repository to persist the new Person in storage.
type CreatePersonHandler struct {
	repo irepo.IPerson // Repository interface for person-related database operations
}

// Compile-time check to ensure CreatePersonHandler implements IHandler for CreatePersonCommand.
var _ icmd.IHandler[*CreatePersonCommand, *model.Person] = &CreatePersonHandler{}

// NewCreatePersonHandler initializes a new CreatePersonHandler with a given IPerson repository.
func NewCreatePersonHandler(repo irepo.IPerson) *CreatePersonHandler {
	return &CreatePersonHandler{repo: repo}
}

// Handle processes the CreatePersonCommand to create a new Person entity.
// It first constructs a Person model and then saves it via the repository.
// Returns the created Person or an error if the process fails.
func (h *CreatePersonHandler) Handle(command *CreatePersonCommand) (*model.Person, ierr.IErr) {
	// Create a new Person model using the provided command data
	person, err := model.CreatePerson(&model.PersonConfig{
		Name:    command.Name,
		Age:     command.Age,
		Hobbies: command.Hobbies,
	})
	if err != nil {
		return nil, err // Return if model creation fails
	}

	// Save the Person model in the repository
	if err := h.repo.Save(person); err != nil {
		return nil, err // Return if saving fails
	}

	return person, nil // Return the newly created Person
}
