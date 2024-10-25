package query

import (
	iquery "github.com/Efamamo/GoCrudChallange/application/common/cqrs/query"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

var _ iquery.IHandler[uuid.UUID, *model.Person] = &GetPersonHandler{}

type GetPersonHandler struct {
	repo irepo.IPerson
}

func NewGetPersonHandler(repo irepo.IPerson) *GetPersonHandler {
	return &GetPersonHandler{repo: repo}
}

func (h *GetPersonHandler) Handle(id uuid.UUID) (*model.Person, error) {
	person, err := h.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return person, nil
}
