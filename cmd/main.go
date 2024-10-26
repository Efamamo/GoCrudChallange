package main

import (
	"github.com/Efamamo/GoCrudChallange/api/controller"
	"github.com/Efamamo/GoCrudChallange/api/router"
	"github.com/Efamamo/GoCrudChallange/application/people/command"
	"github.com/Efamamo/GoCrudChallange/application/people/query"
	"github.com/Efamamo/GoCrudChallange/infrastructure/repository"
)

func main() {
	// Initialize the person repository.
	personRepo := repository.NewPersonRepo()

	// Create command handlers for various person-related operations.
	createPersonHandler := command.NewCreatePersonHandler(personRepo)
	updatePersonHandler := command.NewUpdatePersonHandler(personRepo)
	deletePersonHandler := command.NewDeletePersonHandler(personRepo)

	// Create query handlers for retrieving person data.
	getPersonHandler := query.NewGetPersonHandler(personRepo)
	getAllPersonsHandler := query.NewGetPeopleHandler(personRepo)

	// Create a PersonController with the initialized handlers.
	personController := controller.PersonController{
		CreateHandler: createPersonHandler,
		UpdateHandler: updatePersonHandler,
		DeleteHandler: deletePersonHandler,
		GetHandler:    getPersonHandler,
		GetAllHandler: getAllPersonsHandler,
	}

	// Start the API router with the person controller to handle requests.
	router.StartRouter(personController)
}
