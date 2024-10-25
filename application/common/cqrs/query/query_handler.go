package iquery

type IHandler[Query any, Result any] interface {
	Handle(query Query) (Result, error)
}
