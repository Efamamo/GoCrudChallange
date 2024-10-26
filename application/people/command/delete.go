package command

import (
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	ierr "github.com/Efamamo/GoCrudChallange/domain/error"
	"github.com/google/uuid"
)

// DeletePersonHandler is a command handler for deleting a person by their ID.
type DeletePersonHandler struct {
	repo irepo.IPerson // Repository interface for person operations.
}

// Ensure DeletePersonHandler implements the IHandler interface for handling commands.
var _ icmd.IHandler[uuid.UUID, bool] = &DeletePersonHandler{}

// NewDeletePersonHandler creates a new instance of DeletePersonHandler with the provided repository.
func NewDeletePersonHandler(repo irepo.IPerson) *DeletePersonHandler {
	return &DeletePersonHandler{repo: repo}
}

// Handle processes the command to delete a person by their ID.
func (h *DeletePersonHandler) Handle(id uuid.UUID) (bool, ierr.IErr) {
	if err := h.repo.Delete(id); err != nil {
		return false, err
	}
	return true, nil
}
