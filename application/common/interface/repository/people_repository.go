package irepo

import (
	ierr "github.com/Efamamo/GoCrudChallange/domain/error"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

// IPerson defines the interface for the repository layer responsible for CRUD operations on Person entities.
type IPerson interface {
	// Save adds a new Person to the repository or updates an existing one.
	Save(*model.Person) ierr.IErr

	// Get retrieves a Person by their unique UUID.
	Get(uuid.UUID) (*model.Person, ierr.IErr)

	// Delete removes a Person from the repository by their UUID.
	Delete(uuid.UUID) ierr.IErr

	// GetAll retrieves all Person entities in the repository.
	GetAll() ([]*model.Person, ierr.IErr)
}
