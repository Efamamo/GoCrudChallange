package icmd

type IHandler[Command any, Result any] interface {
	Handle(command Command) (Result, error)
}
