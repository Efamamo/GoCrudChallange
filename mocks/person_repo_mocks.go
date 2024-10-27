package mocks

import (
	"sync"

	ierr "github.com/Efamamo/GoCrudChallange/domain/error"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

// MockPersonRepo is a mock implementation of the IPerson repository interface.
type MockPersonRepo struct {
	mutex      sync.RWMutex
	people     map[uuid.UUID]*model.Person
	SaveFunc   func(person *model.Person) ierr.IErr
	GetFunc    func(id uuid.UUID) (*model.Person, ierr.IErr)
	DeleteFunc func(id uuid.UUID) ierr.IErr
	GetAllFunc func() ([]*model.Person, ierr.IErr)
}

// NewMockPersonRepo creates a new instance of MockPersonRepo with default behavior.
func NewMockPersonRepo() *MockPersonRepo {
	return &MockPersonRepo{
		people: make(map[uuid.UUID]*model.Person),
	}
}

// Save mocks saving a person to the repository.
func (m *MockPersonRepo) Save(p *model.Person) ierr.IErr {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.SaveFunc != nil {
		return m.SaveFunc(p)
	}

	if p == nil {
		return ierr.NewValidation("person can't be empty")
	}

	m.people[p.Id()] = p
	return nil
}

// Get mocks retrieving a person by ID.
func (m *MockPersonRepo) Get(id uuid.UUID) (*model.Person, ierr.IErr) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if m.GetFunc != nil {
		return m.GetFunc(id)
	}

	if person, found := m.people[id]; found {
		return person, nil
	}
	return nil, ierr.NewNotFound("person not found")
}

// Delete mocks removing a person from the repository by ID.
func (m *MockPersonRepo) Delete(id uuid.UUID) ierr.IErr {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}

	if _, found := m.people[id]; found {
		delete(m.people, id)
		return nil
	}
	return ierr.NewNotFound("person not found")
}

// GetAll mocks retrieving all persons in the repository.
func (m *MockPersonRepo) GetAll() ([]*model.Person, ierr.IErr) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}

	var people []*model.Person
	for _, p := range m.people {
		people = append(people, p)
	}
	return people, nil
}
