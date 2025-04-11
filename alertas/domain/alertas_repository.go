package domain

import "2da_api_go/recibos/domain/entities"

type ReciboRepository interface {
	Save(recibo *entities.Recibo) error
}
