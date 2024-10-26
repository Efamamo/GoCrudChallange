package icmd

import ierr "github.com/Efamamo/GoCrudChallange/domain/common"

type IHandler[Command any, Result any] interface {
	Handle(command Command) (Result, ierr.IErr)
}
