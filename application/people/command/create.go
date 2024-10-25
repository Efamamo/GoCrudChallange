package command

import (
	icmd "github.com/Efamamo/GoCrudChallange/application/common/cqrs/command"
	irepo "github.com/Efamamo/GoCrudChallange/application/common/interface/repository"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
)

type CreatePersonCommand struct {
	Name    string
	Age     int16
	Hobbies []string
}

type CreatePersonHandler struct {
	repo irepo.IPerson
}

var _ icmd.IHandler[*CreatePersonCommand, *model.Person] = &CreatePersonHandler{}

func NewCreatePersonHandler(repo irepo.IPerson) *CreatePersonHandler {
	return &CreatePersonHandler{repo: repo}
}

func (h *CreatePersonHandler) Handle(command *CreatePersonCommand) (*model.Person, error) {
	person, err := model.CreatePerson(&model.PersonConfig{
		Name:    command.Name,
		Age:     command.Age,
		Hobbies: command.Hobbies,
	})
	if err != nil {
		return nil, err
	}

	if err := h.repo.Save(person); err != nil {
		return nil, err
	}

	return person, nil
}
