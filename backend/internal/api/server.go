package api

import (
	"fmt"
	"log"
	"nxg/configs"
	"nxg/internal/api/rest"
	"nxg/internal/api/rest/handlers"
	"nxg/internal/domain"
	"nxg/internal/helper"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartServer(config configs.AppConfig) {
	fmt.Println("Server started at ", time.Now())
	app := fiber.New()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect to database")
	}

	log.Println("Database is connected")

	err = db.AutoMigrate(&domain.User{}, &domain.Designation{}, &domain.Product{}, &domain.ProductAttribute{}, &domain.Company{}, &domain.Employee{})
	if err != nil {
		fmt.Println(err)
		panic("failed to migrate to database")

	}

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: auth,
	}
	SetupRoutes(rh)
	app.Listen(config.ServerPort)
}

func SetupRoutes(rh *rest.RestHandler) {
	handlers.SetupUsersRoutes(rh)
	handlers.SetupDesignationRoutes(rh)
	handlers.SetupEmployeeRoutes(rh)
	handlers.SetupCompanyRoutes(rh)
	handlers.SetupProductRoutes(rh)
}
