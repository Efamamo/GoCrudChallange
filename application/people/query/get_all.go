package query

import (
	iquery "github.com/Efamamo/GoCrudChallange/application/common/cqrs/query"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
)

// Ensure GetPeopleHandler implements the IHandler interface for handling queries.
var _ iquery.IHandler[struct{}, []*model.Person] = &GetPeopleHandler{}

// GetPeopleHandler is a query handler for retrieving all people from the repository.
type GetPeopleHandler struct {
	repo irepo.IPerson // Repository interface for person operations.
}

// NewGetPeopleHandler creates a new instance of GetPeopleHandler with the provided repository.
func NewGetPeopleHandler(repo irepo.IPerson) *GetPeopleHandler {
	return &GetPeopleHandler{repo: repo}
}

// Handle processes the query to retrieve all people.
// It returns a slice of pointers to Person models and an error, if any.
func (h *GetPeopleHandler) Handle(_ struct{}) ([]*model.Person, error) {
	// Retrieve all people from the repository.
	people, err := h.repo.GetAll()
	if err != nil {
		return nil, err // Return nil and the error if the retrieval fails.
	}
	return people, nil // Return the slice of people and no error.
}
