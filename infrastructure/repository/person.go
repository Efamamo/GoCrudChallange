package repository

import (
	ierr "github.com/Efamamo/GoCrudChallange/domain/common"
	errdmn "github.com/Efamamo/GoCrudChallange/domain/error/common/error"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

// PersonRepo is an in-memory repository for managing Person entities.
type PersonRepo struct {
	people []*model.Person // Slice to store Person entities.
}

// NewPersonRepo creates and returns a new instance of PersonRepo.
func NewPersonRepo() *PersonRepo {
	return &PersonRepo{
		people: make([]*model.Person, 0), // Initialize the slice for storing persons.
	}
}

// Save saves a Person to the repository.
// If the person already exists, it updates the existing entry.
func (r *PersonRepo) Save(person *model.Person) ierr.IErr {
	if person == nil {
		return errdmn.NewValidation("person can't be empty") // Validate that the person is not nil.
	}

	var idx = -1 // Initialize index to -1 to check if person exists.

	// Check if the person already exists in the repository.
	for i, psn := range r.people {
		if psn.Id() == person.Id() {
			idx = i // Update index if the person is found.
		}
	}

	if idx != -1 {
		r.people[idx] = person // Update the existing person.
	} else {
		r.people = append(r.people, person) // Add a new person to the repository.
	}

	return nil // Return nil indicating success.
}

// Get retrieves a Person by its ID from the repository.
// Returns the Person and an error if not found.
func (r *PersonRepo) Get(id uuid.UUID) (*model.Person, ierr.IErr) {
	var person *model.Person // Initialize variable to hold the found person.

	// Search for the person in the repository.
	for _, p := range r.people {
		if p.Id() == id {
			person = p // Set the found person.
		}
	}

	if person == nil {
		return nil, errdmn.NewNotFound("person not found") // Return error if not found.
	}
	return person, nil // Return the found person.
}

// Delete removes a Person from the repository by its ID.
// Returns an error if the person is not found.
func (r *PersonRepo) Delete(id uuid.UUID) ierr.IErr {
	idx := -1 // Initialize index to -1 to check if person exists.

	// Search for the person to delete.
	for i, p := range r.people {
		if p.Id() == id {
			idx = i // Update index if the person is found.
		}
	}

	if idx == -1 {
		return errdmn.NewNotFound("person not found") // Return error if not found.
	}

	// Remove the person from the slice.
	r.people = append(r.people[0:idx], r.people[idx+1:]...)
	return nil // Return nil indicating success.
}

// GetAll retrieves all Person entities from the repository.
func (r *PersonRepo) GetAll() ([]*model.Person, ierr.IErr) {
	return r.people, nil // Return the slice of all persons.
}
