package router

import (
	"io"
	"os"

	"github.com/abdolrhman/simple-go-lang-rest-api/controller"
	"github.com/abdolrhman/simple-go-lang-rest-api/pkg/logger"
	"github.com/abdolrhman/simple-go-lang-rest-api/pkg/middleware"
	"github.com/gin-contrib/gzip"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	// Write gin access log to file
	f, err := os.OpenFile("ugin.access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Errorf("Failed to create access log file: %v", err)
	} else {
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// Set default middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Set custom middlewares
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middleware.Security())
	r.Use(middleware.MyLimit())

	api := controller.Controller{DB: db}

	// Non-protected routes
	posts := r.Group("/customers")
	{
		posts.GET("/", api.GetCustomers)
	}

	return r
}
