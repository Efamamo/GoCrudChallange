package query

import (
	iquery "github.com/Efamamo/GoCrudChallange/application/common/cqrs/query"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

// Ensure GetPersonHandler implements the IHandler interface for handling queries.
var _ iquery.IHandler[uuid.UUID, *model.Person] = &GetPersonHandler{}

// GetPersonHandler is a query handler for retrieving a specific person by their ID.
type GetPersonHandler struct {
	repo irepo.IPerson // Repository interface for person operations.
}

// NewGetPersonHandler creates a new instance of GetPersonHandler with the provided repository.
func NewGetPersonHandler(repo irepo.IPerson) *GetPersonHandler {
	return &GetPersonHandler{repo: repo}
}

// Handle processes the query to retrieve a person by their ID.
// It returns a pointer to the Person model and an error, if any.
func (h *GetPersonHandler) Handle(id uuid.UUID) (*model.Person, error) {
	// Retrieve the person from the repository using the provided ID.
	person, err := h.repo.Get(id)
	if err != nil {
		return nil, err // Return nil and the error if the retrieval fails.
	}

	return person, nil // Return the person and no error if successful.
}
