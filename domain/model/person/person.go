package model

import (
	"fmt"

	ierr "github.com/Efamamo/GoCrudChallange/domain/error"
	"github.com/google/uuid"
)

// Person represents an individual with a unique ID, name, age, and hobbies.
type Person struct {
	id      uuid.UUID
	name    string
	age     int16
	hobbies []string
}

// PersonConfig is a configuration struct used to create a new Person.
type PersonConfig struct {
	Name    string
	Age     int16
	Hobbies []string
}

// CreatePerson initializes a new Person based on the provided configuration.
func CreatePerson(pc *PersonConfig) (*Person, ierr.IErr) {
	newPerson := &Person{
		id: uuid.New(),
	}

	// Validate and set the name.
	nameErr := newPerson.SetName(pc.Name)
	if nameErr != nil {
		return nil, nameErr
	}

	// Validate and set the age.
	ageErr := newPerson.SetAge(pc.Age)
	if ageErr != nil {
		return nil, ageErr
	}

	// Set hobbies for the person.
	newPerson.SetHobbies(pc.Hobbies)

	return newPerson, nil
}

// SetName sets the name of the person after validating its length.
func (p *Person) SetName(name string) ierr.IErr {
	min := 5
	max := 50
	if len(name) < min || len(name) > max {
		return ierr.NewValidation(fmt.Sprintf("name length should be between %d and %d", min, max))
	}

	p.name = name
	return nil
}

// SetAge sets the age of the person after validating it.
func (p *Person) SetAge(age int16) ierr.IErr {
	if age < 0 {
		return ierr.NewValidation("age should be greater than or equal to 0")
	}
	p.age = age
	return nil
}

// SetHobbies sets the hobbies of the person.
func (p *Person) SetHobbies(hobbies []string) {
	p.hobbies = hobbies
}

// Id returns the unique identifier of the person.
func (p *Person) Id() uuid.UUID {
	return p.id
}

// Name returns the name of the person.
func (p *Person) Name() string {
	return p.name
}

// Age returns the age of the person.
func (p *Person) Age() int16 {
	return p.age
}

// Hobbies returns the hobbies of the person.
func (p *Person) Hobbies() []string {
	return p.hobbies
}
