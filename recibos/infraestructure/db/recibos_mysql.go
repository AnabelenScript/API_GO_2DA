package db


import (
	"database/sql"
	"2da_api_go/recibos/domain"
	"2da_api_go/recibos/domain/entities"
	"log"
)

type MySQLRecibosRepository struct {
	DB *sql.DB
}

func NewMySQLRecibosRepository(db *sql.DB) domain.ReciboRepository {
	return &MySQLRecibosRepository{DB: db}
}


func (r *MySQLRecibosRepository) Save(pedidos *entities.Recibo) error {
	query := "INSERT INTO recibos (pedido_id, estatus) VALUES (?, ?)"
	_, err := r.DB.Exec(query, pedidos.PedidoID, pedidos.Estatus)
	if err != nil {
		log.Printf("Error al agrgear el recibo: %v", err)
	}
	return err
}
