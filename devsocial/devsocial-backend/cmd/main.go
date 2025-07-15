package cmd

import (
	"devsocial-backend/database"
	"devsocial-backend/controllers"
	"github.com/gin-gonic/gin"
)

func main () {
	// grab and database function from our devsocial-backend/database; to implement the connection
	database.connect()

	// initialize the routes
	routes := gin.Default()

}