package main

import (
	"my-favorite-pokemon-rest-api/controller"
	"my-favorite-pokemon-rest-api/db"
	"my-favorite-pokemon-rest-api/repository"
	"my-favorite-pokemon-rest-api/router"
	"my-favorite-pokemon-rest-api/usecase"
	"my-favorite-pokemon-rest-api/validator"
)

func main() {
	connectDB := db.NewDB()
	userValidation := validator.NewUserValidation()
	starValidation := validator.NewStarValidation()
	userRepository := repository.NewUserRepository(connectDB)
	starRepository := repository.NewStarRepository(connectDB)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidation)
	starUsecase := usecase.NewStarUsecase(starRepository, starValidation)
	userController := controller.NewUserController(userUsecase)
	starController := controller.NewStarController(starUsecase)
	e := router.NewRouter(userController, starController)
	e.Logger.Fatal(e.Start(":8080"))
}
