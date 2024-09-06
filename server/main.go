package main

import (
	"log"
	"os"
	"server/app/config"
	"server/app/firebase"
	"server/app/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	db := config.InitDB()
	r := gin.Default()
	secret := os.Getenv("SECRET_KEY")
	expiration := 24 * time.Hour
	cs, err := firebase.DefaultCloudStorage()
	if err != nil {
		panic(err)
	}
	r.Use(cors.Default())
	r.Use(gin.Logger())

	routes.SetUpUserRoutes(r, db, []byte(secret), expiration)
	routes.SetupCourseRoutes(r, db)
	routes.SetupGradeRoutes(r, db)
	routes.SetupAssignmentRoutes(r, db)
	routes.SetupEnrollmentRoutes(r, db)
	routes.SetupSubmissionRoutes(r, db, cs)
	routes.SetupMaterialRoutes(r, db, cs)
	routes.SetupQuizRoutes(r, db, secret)
	_ = r.Run()
}
