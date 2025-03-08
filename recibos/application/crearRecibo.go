package application

import (
	"2da_api_go/recibos/domain"
	"2da_api_go/recibos/domain/entities"
)

type CreateRecibo struct {
	Repo domain.ReciboRepository
}

func NewCreateRecibo(repo domain.ReciboRepository) *CreateRecibo {
	return &CreateRecibo{Repo: repo}
}

func (uc *CreateRecibo) Execute(pedidoID int, estatus string) error {
	recibo := &entities.Recibo{PedidoID: pedidoID, Estatus: estatus}
	return uc.Repo.Save(recibo)
}
