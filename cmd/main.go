package main

import (
	"github.com/Efamamo/GoCrudChallange/api/controller"
	"github.com/Efamamo/GoCrudChallange/api/router"
	"github.com/Efamamo/GoCrudChallange/application/people/command"
	"github.com/Efamamo/GoCrudChallange/application/people/query"
	"github.com/Efamamo/GoCrudChallange/infrastructure/repository"
)

func main() {
	personRepo := repository.NewPersonRepo()
	createPersonHandler := command.NewCreatePersonHandler(personRepo)
	updatePersonHandler := command.NewUpdatePersonHandler(personRepo)
	deletePersonHandler := command.NewDeletePersonHandler(personRepo)
	getPersonHandler := query.NewGetPersonHandler(personRepo)
	getAllPersonsHandler := query.NewGetAllPersonHandler(personRepo)

	pc := controller.PersonController{
		CreateHandler: createPersonHandler,
		UpdateHandler: updatePersonHandler,
		DeleteHandler: deletePersonHandler,
		GetHandler:    getPersonHandler,
		GetAllHandler: getAllPersonsHandler,
	}

	router.StartRouter(pc)
}
