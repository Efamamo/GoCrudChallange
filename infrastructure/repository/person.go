package repository

import (
	"errors"

	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

type PersonRepo struct {
	people []*model.Person
}

func NewPersonRepo() *PersonRepo {
	return &PersonRepo{
		people: []*model.Person{},
	}
}

func (r *PersonRepo) Save(person *model.Person) error {
	if person == nil {
		return errors.New("person cant be empty")
	}

	var idx = -1

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

	return nil
}

func (r *PersonRepo) Get(id uuid.UUID) (*model.Person, error) {

	var person *model.Person

	for _, p := range r.people {
		if p.Id() == id {
			person = p
		}
	}

	if person == nil {
		return nil, errors.New("person not found")
	}
	return person, nil
}

func (r *PersonRepo) Delete(id uuid.UUID) error {

	idx := -1
	for i, p := range r.people {
		if p.Id() == id {
			idx = i
		}
	}

	if idx == -1 {
		return errors.New("person not found")
	}

	r.people = append(r.people[0:idx], r.people[idx+1:]...)
	return nil
}

func (r *PersonRepo) GetAll() ([]*model.Person, error) {
	return r.people, nil
}
