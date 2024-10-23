package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohit-051/hirego/internal/handlers"
	"github.com/mohit-051/hirego/internal/middleware"
)

func SetupV1Routes(router *gin.Engine) {

	v1 := router.Group("/auth")
	{
		v1.POST("/signup", handlers.UserHandler.Signup)
		v1.POST("/login", handlers.UserHandler.Login)
		v1.POST("/hrsignup", handlers.HRHandler.Signup)
		v1.POST("/hrlogin", handlers.HRHandler.Login)
	}
}

func SetupUserRoutes(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	user := router.Group("/user")
	{ // Add all the available middlewares which needs to be attached to the router.
		for _, middleware := range middlewares {
			user.Use(middleware)
		}
		user.POST("/userprofileinformation", handlers.UserHandler.SetUserProfileInformation)
		user.GET("/userprofileinformation", handlers.UserHandler.GetProfileInformation)
		user.POST("/userworkinformation", handlers.UserHandler.SetUserWrorkInformation)
		user.POST("/intenship", handlers.HRHandler.GetJobPosting)
		user.GET("/intenships", handlers.HRHandler.GetAllJobPosting)
		user.POST("/internship/apply", handlers.UserHandler.ApplyForJobPosting)
		user.GET("/internship/applied", handlers.UserHandler.GetAppliedJobPosting)
	}
}

func SetupHrRoutes(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	hr := router.Group("/hr")
	{
		for _, middleware := range middlewares {
			hr.Use(middleware)
		}
		hr.POST("/hrprofileinformation", handlers.HRHandler.SetHrProfileInformation)
		hr.GET("/hrprofileinformation", handlers.HRHandler.GetProfileInformation)
		hr.POST("/jobposting", handlers.HRHandler.JobPosting)          // Route for HR to post a job
		hr.GET("/jobposting", handlers.HRHandler.HrSpecificJobPosting) // Get all the job postings by the particular HR
		hr.POST("/userpublicinformation", handlers.UserHandler.GetUserWorkInformation)
		hr.POST("/applicants", handlers.HRHandler.GetAllApplicants)
		hr.POST("/userworkinformation", handlers.UserHandler.SetUserWrorkInformation)
	}
}

func SetupPublicRoutes(router *gin.Engine) {
	public := router.Group("/public")
	{
		public.GET("/internships", handlers.HRHandler.GetAllJobPosting)

		//Sample public route
	}
}

func InitializeRoutes(router *gin.Engine) {

	// Add the middleware to the parent route.
	SetupV1Routes(router)
	SetupUserRoutes(router, middleware.AuthMiddleware())
	SetupHrRoutes(router, middleware.AuthMiddleware())
	SetupPublicRoutes(router)
}
