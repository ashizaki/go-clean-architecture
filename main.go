package main

import (
	"github.com/ashizaki/go-clean-architecture/domain/repository"
	"github.com/ashizaki/go-clean-architecture/domain/service"
	"github.com/ashizaki/go-clean-architecture/infrastructure/orm"
	"github.com/ashizaki/go-clean-architecture/infrastructure/router"
	"github.com/ashizaki/go-clean-architecture/interface/controller"
	"github.com/ashizaki/go-clean-architecture/interface/database"
	"github.com/ashizaki/go-clean-architecture/usecase/interactor"
)

func main() {
	apiV1 := router.G.Group("/v1")

	handler, err := orm.OpenDatabase("test.db")
	if err != nil {
		panic(err)
	}

	userController := initializeUserAPI(handler)
	router.InitUserRouter(apiV1, userController)

	if err := router.G.Run(":8080"); err != nil {
		panic(err.Error())
	}
}

func initializeUserAPI(handler repository.DbHandler) controller.UserController {
	uRepo := database.NewUserRepository()
	uService := service.NewUserService(uRepo)
	uInteractor := interactor.NewUserInteractor(handler, uService, uRepo)
	return controller.NewUserController(uInteractor)
}
