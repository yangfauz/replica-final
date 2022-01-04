package main

import (
	"net/http"
	"replica-finalproject/api/controller"
	"replica-finalproject/api/repository"
	"replica-finalproject/api/responder"
	"replica-finalproject/api/service"
	"replica-finalproject/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//setup configuration
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	//migration
	//database.AutoMigrate(&migration.Role{}, &migration.Status{}, &migration.User{}, &migration.Payment{})

	//setup Repository
	userRepository := repository.NewUserRepository(database)
	roleRepository := repository.NewRoleRepository(database)
	paymentRepository := repository.NewPaymentRepository(database)
	statusRepository := repository.NewStatusRepository(database)

	//setup Service
	authService := service.NewAuthService(&userRepository)
	userService := service.NewUserService(&userRepository)
	roleService := service.NewRoleService(&roleRepository)
	paymentService := service.NewPaymentService(&paymentRepository)
	statusService := service.NewStatusService(&statusRepository)

	//setup Handler
	authController := controller.NewAuthController(&authService)
	userController := controller.NewUserController(&userService)
	roleController := controller.NewRoleController(&roleService)
	paymentController := controller.NewPaymentController(&paymentService)
	statusController := controller.NewStatusController(&statusService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	// app.Group("api/v1/")
	app.Use(cors.New())
	authController.Route(app)
	userController.Route(app)
	roleController.Route(app)
	paymentController.Route(app)
	statusController.Route(app)

	//Not Found in Last
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:  http.StatusNotFound,
			Error: &fiber.ErrNotFound.Message,
			Data:  nil,
		})
	})

	// Start App
	port := configuration.Get("APP_PORT")
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
