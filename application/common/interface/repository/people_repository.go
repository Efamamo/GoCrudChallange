package irepo

import (
	ierr "github.com/Efamamo/GoCrudChallange/domain/common"
	model "github.com/Efamamo/GoCrudChallange/domain/model/person"
	"github.com/google/uuid"
)

type IPerson interface {
	Save(*model.Person) ierr.IErr

	Get(uuid.UUID) (*model.Person, ierr.IErr)

	Delete(uuid.UUID) ierr.IErr

	GetAll() ([]*model.Person, ierr.IErr)
}
