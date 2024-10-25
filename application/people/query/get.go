package query

import (
	iquery "github.com/Efamamo/GoCrudChallange/application/common/cqrs/query"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
)

var _ iquery.IHandler[struct{}, []*model.Person] = &GetPeopleHandler{}

type GetPeopleHandler struct {
	repo irepo.IPerson
}

func NewGetAllPersonHandler(repo irepo.IPerson) *GetPeopleHandler {
	return &GetPeopleHandler{repo: repo}
}

func (h *GetPeopleHandler) Handle(_ struct{}) ([]*model.Person, error) {
	persons, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return persons, nil
}
