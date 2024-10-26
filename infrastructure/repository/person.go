package repository

import (
	"sync"

	ierr "github.com/Efamamo/GoCrudChallange/domain/error"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

// PersonRepo is an in-memory repository for managing Person entities.
type PersonRepo struct {
	mutex  sync.RWMutex
	people []*model.Person
}

// NewPersonRepo creates and returns a new instance of PersonRepo.
func NewPersonRepo() *PersonRepo {
	return &PersonRepo{
		people: make([]*model.Person, 0),
	}
}

// Save saves a Person to the repository.
func (r *PersonRepo) Save(person *model.Person) ierr.IErr {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if person == nil {
		return ierr.NewValidation("person can't be empty")
	}

	var idx = -1

	// Check if the person already exists in the repository.
	for i, psn := range r.people {
		if psn.Id() == person.Id() {
			idx = i
		}
	}

	if idx != -1 {
		r.people[idx] = person
	} else {
		r.people = append(r.people, person)
	}

	return nil // Return nil indicating success.
}

// Get retrieves a Person by its ID from the repository.
func (r *PersonRepo) Get(id uuid.UUID) (*model.Person, ierr.IErr) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var person *model.Person

	// Search for the person in the repository.
	for _, p := range r.people {
		if p.Id() == id {
			person = p
		}
	}

	if person == nil {
		return nil, ierr.NewNotFound("person not found")
	}
	return person, nil
}

// Delete removes a Person from the repository by its ID.
func (r *PersonRepo) Delete(id uuid.UUID) ierr.IErr {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	idx := -1

	// Search for the person to delete.
	for i, p := range r.people {
		if p.Id() == id {
			idx = i
		}
	}

	if idx == -1 {
		return ierr.NewNotFound("person not found")
	}

	// Remove the person from the slice.
	r.people = append(r.people[0:idx], r.people[idx+1:]...)
	return nil
}

// GetAll retrieves all Person entities from the repository.
func (r *PersonRepo) GetAll() ([]*model.Person, ierr.IErr) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.people, nil
}
