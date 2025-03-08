package entities

type Recibo struct {
	ID        int    `json:"id"`
	PedidoID  int    `json:"pedido_id"`
	Estatus  string `json:"estatus"`
}