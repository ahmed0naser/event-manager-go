package main

import (
	"github.com/ahmed0naser/event-manager-go/db"
	"github.com/ahmed0naser/event-manager-go/routes"
	"github.com/gin-gonic/gin"
)
func main(){
db.InitDb()
server:=gin.Default()
routes.RegisterRoutes(server)
server.Run(":8080")
}
