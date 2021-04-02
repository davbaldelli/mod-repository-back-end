package main

import (
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/repositories"
	"github.com/davide/ModRepository/routes"
	"github.com/davide/ModRepository/routes/handlers"
)

func main() {

	web := routes.Web{CarHandler: handlers.CarsHandlerImpl{CarCtrl: controllers.CarControllerImpl{Repo: repositories.CarRepositoryImpl{}}}}
	web.Listen("8080")
}
