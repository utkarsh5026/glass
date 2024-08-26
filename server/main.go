package main

import (
	"server/app/config"
	"server/app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gin.Logger())

	routes.SetupCourseRoutes(r, db)
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
