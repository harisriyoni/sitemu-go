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

	beritaRepository := repository.NewBeritaRepository(db)
	beritaService := service.NewBeritaService(beritaRepository)
	beritaController := controller.NewBeritaController(beritaService)

	typeGaleriRepo := repository.NewTypeGaleriRepository(db)
	typeGaleriService := service.NewTypeGaleriService(typeGaleriRepo)
	typeGaleriController := controller.NewTypeGaleriController(typeGaleriService)

	galeriRepo := repository.NewGaleriRepository(db)
	galeriService := service.NewGaleriService(galeriRepo)
	galeriController := controller.NewGaleriController(galeriService)

	prestasiRepository := repository.NewPrestasiRepository(db)
	prestasiService := service.NewPrestasiService(prestasiRepository)
	prestasiController := controller.NewPrestasiController(prestasiService)

	// === Router ===
	router := httprouter.New()

	// === PUBLIC ROUTES ===
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)

	router.GET("/api/berita/all", beritaController.GetAll)         // Public: semua berita
	router.GET("/api/berita/detail/:id", beritaController.GetByID) // Public: detail berita

	router.GET("/api/organisasi/all", organisasiController.GetAll) // Public: semua organisasi

	router.GET("/api/type-galeri/all", typeGaleriController.GetAll)         // Publik: semua tipe galeri
	router.GET("/api/type-galeri/detail/:id", typeGaleriController.GetByID) // Publik: detail tipe galeri

	router.GET("/api/prestasi/all", prestasiController.GetAll)         // Public
	router.GET("/api/prestasi/detail/:id", prestasiController.GetByID) // Public

	router.GET("/api/galeri/all", galeriController.GetAll)
	router.GET("/api/galeri/detail/:id", galeriController.GetByID)

	router.ServeFiles("/public/berita/*filepath", http.Dir("public/berita"))
	router.ServeFiles("/public/organisasi/*filepath", http.Dir("public/organisasi"))
	router.ServeFiles("/public/galeri/*filepath", http.Dir("public/galeri"))

	// === PROTECTED ROUTES ===
	// User
	router.GET("/api/users/profile", middleware.AuthMiddleware(userController.Profile))
	router.PUT("/api/users/profile", middleware.AuthMiddleware(userController.Update))
	router.DELETE("/api/users/profile", middleware.AuthMiddleware(userController.Delete))

	// Organisasi
	router.POST("/api/organisasi", middleware.AuthMiddleware(organisasiController.Create))
	router.GET("/api/organisasi", middleware.AuthMiddleware(organisasiController.GetByUser))
	router.PUT("/api/organisasi/:id", middleware.AuthMiddleware(organisasiController.Update))
	router.DELETE("/api/organisasi/:id", middleware.AuthMiddleware(organisasiController.Delete))

	// Berita
	router.POST("/api/berita", middleware.AuthMiddleware(beritaController.Create))
	router.PUT("/api/berita/:id", middleware.AuthMiddleware(beritaController.Update))
	router.DELETE("/api/berita/:id", middleware.AuthMiddleware(beritaController.Delete))
	router.GET("/api/user/berita", middleware.AuthMiddleware(beritaController.GetByUser))

	// Type Galeri
	router.POST("/api/type-galeri", middleware.AuthMiddleware(typeGaleriController.Create))
	router.PUT("/api/type-galeri/:id", middleware.AuthMiddleware(typeGaleriController.Update))
	router.DELETE("/api/type-galeri/:id", middleware.AuthMiddleware(typeGaleriController.Delete))
	router.GET("/api/user/type-galeri", middleware.AuthMiddleware(typeGaleriController.GetByUser))

	// Galeri
	router.POST("/api/galeri", middleware.AuthMiddleware(galeriController.Create))
	router.PUT("/api/galeri/:id", middleware.AuthMiddleware(galeriController.Update))
	router.DELETE("/api/galeri/:id", middleware.AuthMiddleware(galeriController.Delete))

	// Prestasi
	router.POST("/api/prestasi", middleware.AuthMiddleware(prestasiController.Create))
	router.PUT("/api/prestasi/:id", middleware.AuthMiddleware(prestasiController.Update))
	router.DELETE("/api/prestasi/:id", middleware.AuthMiddleware(prestasiController.Delete))
	router.GET("/api/prestasi", middleware.AuthMiddleware(prestasiController.GetByUser))

	// === Error Handler dan Server ===
	router.PanicHandler = helper.ErrorHandler

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
