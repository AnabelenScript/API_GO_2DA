package main

import (
	"2da_api_go/helpers"
	"2da_api_go/recibos/application"
	"2da_api_go/recibos/infraestructure/controllers"
	"2da_api_go/recibos/infraestructure/db"
	"2da_api_go/recibos/infraestructure/rabbitmq"
	"2da_api_go/recibos/infraestructure/routes"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	dbConn := helpers.ConnectToMySQL()
	defer dbConn.Close()
	reciboRepo := db.NewMySQLRecibosRepository(dbConn)
	createRecibo := application.NewCreateRecibo(reciboRepo)
	createReciboController :=  controllers.NewCreateRecibosController(createRecibo)
	r := gin.Default()
	r.Use(cors.Default())
	routes.SetupRecibosRoutes(r, createReciboController)
	go rabbitmq.StartRabbitMQListener(createRecibo)

	log.Println("Server running on :8081")
	r.Run(":8081")
}
