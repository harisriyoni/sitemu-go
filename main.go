package main

import (
	"log"
	"net/http"
	"os"

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
	// === Init DB & Validator ===
	db := app.NewDB()
	validate := validator.New()

	// === Init Google Drive Service ===
	err := helper.InitDriveService()
	if err != nil {
		log.Fatalf("Gagal inisialisasi Google Drive: %v", err)
	}

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

	router.GET("/api/berita/all", beritaController.GetAll)
	router.GET("/api/berita/detail/:id", beritaController.GetByID)

	router.GET("/api/organisasi/all", organisasiController.GetAll)

	router.GET("/api/type-galeri/all", typeGaleriController.GetAll)
	router.GET("/api/type-galeri/detail/:id", typeGaleriController.GetByID)

	router.GET("/api/prestasi/all", prestasiController.GetAll)
	router.GET("/api/prestasi/detail/:id", prestasiController.GetByID)

	router.GET("/api/galeri/all", galeriController.GetAll)
	router.GET("/api/galeri/detail/:id", galeriController.GetByID)

	// Serve file lokal (legacy fallback, bisa dihapus jika semua pindah ke Drive)
	router.ServeFiles("/public/berita/*filepath", http.Dir("public/berita"))
	router.ServeFiles("/public/organisasi/*filepath", http.Dir("public/organisasi"))
	router.ServeFiles("/public/galeri/*filepath", http.Dir("public/galeri"))

	// === PROTECTED ROUTES ===
	router.GET("/api/users/profile", middleware.AuthMiddleware(userController.Profile))
	router.PUT("/api/users/profile", middleware.AuthMiddleware(userController.Update))
	router.DELETE("/api/users/profile", middleware.AuthMiddleware(userController.Delete))

	router.POST("/api/organisasi", middleware.AuthMiddleware(organisasiController.Create))
	router.GET("/api/organisasi", middleware.AuthMiddleware(organisasiController.GetByUser))
	router.PUT("/api/organisasi/:id", middleware.AuthMiddleware(organisasiController.Update))
	router.DELETE("/api/organisasi/:id", middleware.AuthMiddleware(organisasiController.Delete))

	router.POST("/api/berita", middleware.AuthMiddleware(beritaController.Create))
	router.PUT("/api/berita/:id", middleware.AuthMiddleware(beritaController.Update))
	router.DELETE("/api/berita/:id", middleware.AuthMiddleware(beritaController.Delete))
	router.GET("/api/user/berita", middleware.AuthMiddleware(beritaController.GetByUser))

	router.POST("/api/type-galeri", middleware.AuthMiddleware(typeGaleriController.Create))
	router.PUT("/api/type-galeri/:id", middleware.AuthMiddleware(typeGaleriController.Update))
	router.DELETE("/api/type-galeri/:id", middleware.AuthMiddleware(typeGaleriController.Delete))
	router.GET("/api/user/type-galeri", middleware.AuthMiddleware(typeGaleriController.GetByUser))

	router.POST("/api/galeri", middleware.AuthMiddleware(galeriController.Create))
	router.PUT("/api/galeri/:id", middleware.AuthMiddleware(galeriController.Update))
	router.DELETE("/api/galeri/:id", middleware.AuthMiddleware(galeriController.Delete))

	router.POST("/api/prestasi", middleware.AuthMiddleware(prestasiController.Create))
	router.PUT("/api/prestasi/:id", middleware.AuthMiddleware(prestasiController.Update))
	router.DELETE("/api/prestasi/:id", middleware.AuthMiddleware(prestasiController.Delete))
	router.GET("/api/prestasi", middleware.AuthMiddleware(prestasiController.GetByUser))

	// === Error Handler ===
	router.PanicHandler = helper.ErrorHandler

	// === PORT ===
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// === Start Server with CORS Middleware ===
	server := http.Server{
		Addr:    ":" + port,
		Handler: middleware.CORSMiddleware(router),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
