package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Person struct {
	id      uuid.UUID
	name    string
	age     int16
	hobbies []string
}

type PersonConfig struct {
	Name    string
	Age     int16
	Hobbies []string
}

func CreatePerson(pc *PersonConfig) (*Person, error) {
	newPerson := &Person{
		id:      uuid.New(),
		hobbies: pc.Hobbies,
	}

	nameErr := newPerson.SetName(pc.Name)
	if nameErr != nil {
		return nil, nameErr
	}

	ageErr := newPerson.SetAge(pc.Age)

	if ageErr != nil {
		return nil, ageErr
	}

	return newPerson, nil
}

func (p *Person) SetName(name string) error {
	min := 5
	max := 50
	if len(name) < min || len(name) > max {
		return fmt.Errorf("name length should be between %d and %d", min, max)
	}
	p.name = name
	return nil
}

func (p *Person) SetAge(age int16) error {
	if age < 0 {
		return fmt.Errorf("age should be greater than 0")
	}
	p.age = age
	return nil
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Age() int16 {
	return p.age
}

func (p *Person) Hobbies() []string {
	return p.hobbies
}
