package infrastructure

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() Router {
	httpRouter := gin.Default()
	// msg := fmt.Sprintf("TeamPlace %s API Server is Running...", )
	return Router{httpRouter}
}

func (r *Router) RunServer() {
	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
