package dtos

import "github.com/google/uuid"

type CreatePersonDTO struct {
	Name    string   `json:"name" binding:"required"`
	Age     int16    `json:"age" binding:"required"`
	Hobbies []string `json:"hobbies"`
}

type PersonResponseDTO struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Age     int16     `json:"age"`
	Hobbies []string  `json:"hobbies"`
}
