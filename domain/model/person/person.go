package model

import (
	"fmt"

	ierr "github.com/Efamamo/GoCrudChallange/domain/common"
	errdmn "github.com/Efamamo/GoCrudChallange/domain/error/common/error"
	"github.com/google/uuid"
)

// Person represents an individual with a unique ID, name, age, and hobbies.
type Person struct {
	id      uuid.UUID // Unique identifier for the person.
	name    string    // Name of the person.
	age     int16     // Age of the person.
	hobbies []string  // List of hobbies associated with the person.
}

// PersonConfig is a configuration struct used to create a new Person.
type PersonConfig struct {
	Name    string   // Name of the person.
	Age     int16    // Age of the person.
	Hobbies []string // List of hobbies for the person.
}

// CreatePerson initializes a new Person based on the provided configuration.
// Returns a pointer to the created Person and an error if validation fails.
func CreatePerson(pc *PersonConfig) (*Person, ierr.IErr) {
	newPerson := &Person{
		id: uuid.New(), // Generate a new unique ID for the person.
	}

	// Validate and set the name.
	nameErr := newPerson.SetName(pc.Name)
	if nameErr != nil {
		return nil, nameErr // Return nil and the error if name validation fails.
	}

	// Validate and set the age.
	ageErr := newPerson.SetAge(pc.Age)
	if ageErr != nil {
		return nil, ageErr // Return nil and the error if age validation fails.
	}

	// Set hobbies for the person.
	newPerson.SetHobbies(pc.Hobbies)

	return newPerson, nil // Return the newly created person and no error.
}

// SetName sets the name of the person after validating its length.
// Returns an error if the name does not meet validation criteria.
func (p *Person) SetName(name string) ierr.IErr {
	min := 5  // Minimum length for the name.
	max := 50 // Maximum length for the name.
	if len(name) < min || len(name) > max {
		return errdmn.NewValidation(fmt.Sprintf("name length should be between %d and %d", min, max))
	}

	p.name = name // Set the validated name.
	return nil    // Return nil if successful.
}

// SetAge sets the age of the person after validating it.
// Returns an error if the age is less than 0.
func (p *Person) SetAge(age int16) ierr.IErr {
	if age < 0 {
		return errdmn.NewValidation("age should be greater than or equal to 0")
	}
	p.age = age // Set the validated age.
	return nil  // Return nil if successful.
}

// SetHobbies sets the hobbies of the person.
func (p *Person) SetHobbies(hobbies []string) {
	p.hobbies = hobbies // Assign the hobbies to the person.
}

// Id returns the unique identifier of the person.
func (p *Person) Id() uuid.UUID {
	return p.id // Return the person's ID.
}

// Name returns the name of the person.
func (p *Person) Name() string {
	return p.name // Return the person's name.
}

// Age returns the age of the person.
func (p *Person) Age() int16 {
	return p.age // Return the person's age.
}

// Hobbies returns the hobbies of the person.
func (p *Person) Hobbies() []string {
	return p.hobbies // Return the person's hobbies.
}
