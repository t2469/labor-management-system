package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/t2469/labor-management-system.git/config"
	"github.com/t2469/labor-management-system.git/controllers"
	"net/http"
	"time"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	setCors(router, cfg)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
	})

	employees := router.Group("/employees")
	{
		employees.GET("", controllers.GetEmployees)
		employees.GET("/:id", controllers.GetEmployee)
		employees.POST("", controllers.CreateEmployee)

		employees.POST("/:id/attendances", controllers.CreateAttendance)

		employees.GET("/:id/insurance", controllers.CalculateEmployeeInsurance)
		employees.GET("/:id/pension", controllers.CalculateEmployeePension)
		employees.GET("/:id/payroll", controllers.CalculateEmployeePayroll)
	}

	companies := router.Group("/companies")
	{
		companies.POST("", controllers.CreateCompany)
	}

	allowances := router.Group("/allowances")
	{
		allowances.POST("type", controllers.CrateAllowanceType)
		allowances.POST("", controllers.CreateEmployeeAllowance)
	}
	return router
}

func setCors(router *gin.Engine, cfg *config.Config) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.AllowedOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
