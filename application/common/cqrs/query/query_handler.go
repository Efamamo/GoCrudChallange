package iquery

// IHandler is a generic interface for handling queries in the CQRS architecture.
// It defines a method for processing queries and returning a result or an error.
//
// Query  - the type of query to handle (e.g., a data retrieval operation).
// Result - the type of result expected after handling the query.
// Handle method processes the query and returns a result of type Result or an error.
type IHandler[Query any, Result any] interface {
	Handle(query Query) (Result, error) // Executes the query and returns the result or an error
}
