package controllers

import (
	"2da_api_go/recibos/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateRecibosController struct {
	useCase *application.CreateRecibo
}

func NewCreateRecibosController(useCase *application.CreateRecibo) *CreateRecibosController {
	return &CreateRecibosController{useCase: useCase}
}

func (c *CreateRecibosController) Execute(ctx *gin.Context) {
	var input struct {
		pedido_id int `json:"pedido_id"`
		estatus string   `json:"estatus"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}
	if err := c.useCase.Execute(input.pedido_id, input.estatus); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el postre"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Postre creado creado :)"})
}
