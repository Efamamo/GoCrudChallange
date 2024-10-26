package irepo

import (
	ierr "github.com/Efamamo/GoCrudChallange/domain/common"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

// IPerson defines the interface for the repository layer responsible for CRUD operations on Person entities.
// Each method returns an ierr.IErr to handle custom errors as defined in the domain layer.
type IPerson interface {
	// Save adds a new Person to the repository or updates an existing one.
	// Returns a custom error type if the operation fails.
	Save(*model.Person) ierr.IErr

	// Get retrieves a Person by their unique UUID.
	// Returns the Person if found, or a custom error if not.
	Get(uuid.UUID) (*model.Person, ierr.IErr)

	// Delete removes a Person from the repository by their UUID.
	// Returns a custom error if the operation fails.
	Delete(uuid.UUID) ierr.IErr

	// GetAll retrieves all Person entities in the repository.
	// Returns a slice of pointers to Person and a custom error if any issues occur.
	GetAll() ([]*model.Person, ierr.IErr)
}
