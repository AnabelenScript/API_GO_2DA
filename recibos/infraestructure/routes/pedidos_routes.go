package routes

import(
	"2da_api_go/recibos/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRecibosRoutes(
	r *gin.Engine,
	NewCreateRecibosController *controllers.CreateRecibosController,
	){
	r.POST("/pedidos", NewCreateRecibosController.Execute)
}
