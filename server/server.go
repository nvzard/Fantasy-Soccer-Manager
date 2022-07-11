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
	r := gin.New()

	// Middleware
	r.Use(gin.Recovery())
	r.Use(ginzap.Ginzap(utils.Logger(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(utils.Logger(), true))

	api := r.Group("/api")
	{
		api.POST("/auth", controller.GenerateToken)
		api.POST("/user", controller.RegisterUser)
	}

	secured := r.Group("/secured").Use(middleware.Auth())
	{
		secured.GET("/ping", controller.Ping)
	}

	// Root Routes
	r.GET("/", root)
	r.GET("/health", healthcheck)
	// r.POST("/token", controller.GenerateToken)
	// r.POST("/user/register", controller.RegisterUser)

	return r
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
