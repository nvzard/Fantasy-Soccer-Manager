package server

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/nvzard/soccer-manager/controller"
	"github.com/nvzard/soccer-manager/middleware"
	"github.com/nvzard/soccer-manager/utils"
)

// var logger *zap.SugaredLogger

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// SetupApiServer attached routes and middleware and starts the server
func SetupApiServer() *gin.Engine {
	router := gin.New()

	// Middleware
	router.Use(gin.Recovery())
	router.Use(ginzap.Ginzap(utils.Logger(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(utils.Logger(), true))

	// Root Routes
	router.GET("/", root)
	router.GET("/health", healthcheck)

	// Public Routes: Create User and Login
	api := router.Group("/api")
	{
		api.POST("/user", controller.RegisterUser)
		api.POST("/auth", controller.GenerateToken)
	}

	// Secure Routes: Require JWT Token in Header
	secured := router.Group("/api").Use(middleware.Auth())
	{
		secured.GET("/ping", controller.Ping)
		secured.GET("/user/:email", controller.GetUser)

		secured.GET("/team/:id", controller.GetTeam)
		// secured.PATCH("/team/:id", controller.UpdateTeam)

		secured.GET("/player/:id", controller.GetPlayer)
		secured.PATCH("/player/:id", controller.UpdatePlayer)
	}

	return router
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from the Soccer Manager API!",
	})
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "OK",
	})
}
