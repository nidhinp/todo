package config

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//postgres database driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/nidhinp/todo/api/controllers"
	"github.com/nidhinp/todo/api/models"
)

// Server defines both DB and router
type Server struct {
	DB *gorm.DB
}

// Initialize the server with database and router
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Todo{})
}

// Run the server on the port provided
func (server *Server) Run(addr string) {
	router := gin.Default()

	router.GET("/", controllers.HomeController)
	router.POST("/login", controllers.Login)

	router.Run(addr)
}
