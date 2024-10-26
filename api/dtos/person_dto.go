package dtos

import "github.com/google/uuid"

// CreateDTO represents the data structure for creating or updating a Person.
// It is used to bind incoming JSON data and validate required fields.
type CreateDTO struct {
	Name    string   `json:"name" binding:"required"` // Name of the person; required for creating or updating
	Age     int16    `json:"age" binding:"required"`  // Age of the person; required for creating or updating
	Hobbies []string `json:"hobbies"`                 // List of hobbies for the person; optional
}

// ResponseDTO defines the data structure for returning Person data in responses.
// It is used to format the outgoing JSON response to the client.
type ResponseDTO struct {
	ID      uuid.UUID `json:"id"`      // Unique identifier of the person
	Name    string    `json:"name"`    // Name of the person
	Age     int16     `json:"age"`     // Age of the person
	Hobbies []string  `json:"hobbies"` // List of hobbies for the person
}
