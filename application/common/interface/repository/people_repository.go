package irepo

import (
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

type IPerson interface {
	Save(person *model.Person) error

	Get(id uuid.UUID) (*model.Person, error)

	Delete(id uuid.UUID) error

	GetAll() ([]*model.Person, error)
}
