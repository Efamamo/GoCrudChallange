package repository

import (
	ierr "github.com/Efamamo/GoCrudChallange/domain/common"
	errdmn "github.com/Efamamo/GoCrudChallange/domain/error/common/error"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

type PersonRepo struct {
	people []*model.Person
}

func NewPersonRepo() *PersonRepo {
	return &PersonRepo{
		people: make([]*model.Person, 0),
	}
}

func (r *PersonRepo) Save(person *model.Person) ierr.IErr {
	if person == nil {
		return errdmn.NewValidation("person cant be empty")
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

func (r *PersonRepo) Get(id uuid.UUID) (*model.Person, ierr.IErr) {

	var person *model.Person

	for _, p := range r.people {
		if p.Id() == id {
			person = p
		}
	}

	if person == nil {
		return nil, errdmn.NewNotFound("person not found")
	}
	return person, nil
}

func (r *PersonRepo) Delete(id uuid.UUID) ierr.IErr {

	idx := -1
	for i, p := range r.people {
		if p.Id() == id {
			idx = i
		}
	}

	if idx == -1 {
		return errdmn.NewNotFound("person not found")
	}

	r.people = append(r.people[0:idx], r.people[idx+1:]...)
	return nil
}

func (r *PersonRepo) GetAll() ([]*model.Person, ierr.IErr) {
	return r.people, nil
}
