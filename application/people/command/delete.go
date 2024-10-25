package command

import (
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	"github.com/google/uuid"
)

type DeletePersonHandler struct {
	repo irepo.IPerson
}

var _ icmd.IHandler[uuid.UUID, bool] = &DeletePersonHandler{}

func NewDeletePersonHandler(repo irepo.IPerson) *DeletePersonHandler {
	return &DeletePersonHandler{repo: repo}
}

func (h *DeletePersonHandler) Handle(id uuid.UUID) (bool, error) {
	if err := h.repo.Delete(id); err != nil {
		return false, err
	}
	return true, nil
}
