package main

import (
	"fmt"
	"github.com/mohit-051/hirego/internal/core/services"
	"github.com/mohit-051/hirego/internal/handlers"
	"github.com/mohit-051/hirego/internal/repository"
	"github.com/mohit-051/hirego/internal/routes"
	// "github.com/mohit-051/hirego/pkg/cronJobs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	// Create a new cron service
	cronService := service.NewCronService()

	// Add a job to run every day at midnight
	_, err := cronService.AddJob("* * * * *", func() {
		fmt.Println("Running cron job every minutes")
		// Call the cron job function
		// cronjobs.Cronjobs()
	})

	if err != nil {
		panic(err)
	}

	// Start the cron scheduler
	cronService.Start()

	// Create a new gin router
	parent_route := gin.Default()

	// Allow all origins and proxy all requests.
	parent_route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize the routes.
	routes.InitializeRoutes(parent_route)

	log.Fatal(parent_route.Run())

}

func run() error {
	// Initialize the database.
	db, err := repository.InitializeDB()

	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	// Initialize the:
	// 1. User repository.
	// 2. User service.
	// 3. User handler.
	userRep := repository.UserRepo.Initialize(db)
	hrRep := repository.HRRepo.Initialize(db)

	users := service.InitializeUserService(userRep)
	hr_manager := service.InitializeHRService(hrRep)

	handlers.UserHandler.Initialize(users)
	handlers.HRHandler.Initialize(hr_manager)

	return nil
}
