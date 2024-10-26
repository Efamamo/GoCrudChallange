package icmd

import ierr "github.com/Efamamo/GoCrudChallange/domain/common"

// IHandler is a generic interface for handling commands in the CQRS architecture.
// It defines a method for processing commands and returning a result or error.
//
// Command - the type of command to handle (e.g., a specific operation or action).
// Result  - the type of result expected after handling the command.
// Handle method processes the command and returns a result of type Result and an error of type ierr.IErr.
type IHandler[Command any, Result any] interface {
	Handle(command Command) (Result, ierr.IErr) // Executes the command and returns the result or an error
}
