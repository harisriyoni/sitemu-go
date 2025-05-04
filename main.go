package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/harisriyoni/sitemu-go/app"
	"github.com/harisriyoni/sitemu-go/controller"
	"github.com/harisriyoni/sitemu-go/helper"
	"github.com/harisriyoni/sitemu-go/middleware"
	"github.com/harisriyoni/sitemu-go/repository"
	"github.com/harisriyoni/sitemu-go/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	// === Dependency Injection ===
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	organisasiRepository := repository.NewOrganisasiRepository(db)
	organisasiService := service.NewOrganisasiService(organisasiRepository, db, validate)
	organisasiController := controller.NewOrganisasiController(organisasiService)

	// === Router Setup ===
	router := httprouter.New()

	// -- User Routes --
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.GET("/api/users/profile", userController.Profile)
	router.PUT("/api/users/profile", userController.Update)
	router.DELETE("/api/users/profile", userController.Delete)

	// -- Organisasi Routes --
	router.POST("/api/organisasi", organisasiController.Create)
	router.GET("/api/organisasi", organisasiController.GetByUser)
	router.PUT("/api/organisasi/:id", organisasiController.Update)
	router.DELETE("/api/organisasi/:id", organisasiController.Delete)
	router.GET("/api/organisasi/all", organisasiController.GetAll)

	// -- Serve Image Publicly --
	router.ServeFiles("/public/organisasi/*filepath", http.Dir("public/organisasi"))

	// -- Middleware & Server --
	router.PanicHandler = helper.ErrorHandler
	server := http.Server{
		Addr: ":8080",
		Handler: middleware.NewAuthMiddlewareWithExclusion(router, []string{
			"/api/users/login",
			"/api/users/register",
			"/api/organisasi/:id",
			"/api/organisasi/all",
			"/public/",
		}),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
