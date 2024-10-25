package cqrs

type Handler[Req any, Res any] interface {
	Handle(request Req) (Res, error)
}
