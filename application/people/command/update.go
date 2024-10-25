package command

import (
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

type UpdatePersonCommand struct {
	ID      uuid.UUID
	Name    string
	Age     int16
	Hobbies []string
}

type UpdatePersonHandler struct {
	repo irepo.IPerson
}

var _ icmd.IHandler[*UpdatePersonCommand, *model.Person] = &UpdatePersonHandler{}

func NewUpdatePersonHandler(repo irepo.IPerson) *UpdatePersonHandler {
	return &UpdatePersonHandler{repo: repo}
}

func (h *UpdatePersonHandler) Handle(command *UpdatePersonCommand) (*model.Person, error) {
	person, err := h.repo.Get(command.ID)
	if err != nil {
		return nil, err
	}

	if err := person.SetName(command.Name); err != nil {
		return nil, err
	}

	if err := person.SetAge(command.Age); err != nil {
		return nil, err
	}

	person.SetHobbies(command.Hobbies)

	if err := h.repo.Save(person); err != nil {
		return nil, err
	}

	return person, nil
}
