package rabbitmq

import (
	"2da_api_go/recibos/application"
	"2da_api_go/helpers"
	"encoding/json"
	"log"
)

type PedidoMensaje struct {
	PedidoID int    `json:"pedido_id"`
	Estatus  string `json:"estatus"`
}

func StartRabbitMQListener(useCase *application.CreateRecibo) {
	conn, ch, err := helpers.ConnectRabbitMQ()
	if err != nil {
		log.Fatal("No se pudo establecer conexión a RabbitMQ")
	}
	defer conn.Close()
	defer ch.Close()

	log.Println("Conexión y canal a RabbitMQ establecidos correctamente")

	_, err = ch.QueueDeclare(
		"pedidos",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error al declarar la cola 'pedidos': ", err)
	}
	msgs, err := ch.Consume(
		"pedidos",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error al consumir mensajes de la cola: ", err)
	}
	go func() {
		for d := range msgs {
			log.Printf("Mensaje recibido: %s", d.Body)

			var pedido PedidoMensaje
			if err := json.Unmarshal(d.Body, &pedido); err != nil {
				log.Printf("Error decodificando JSON: %v", err)
				continue
			}
			if err := useCase.Execute(pedido.PedidoID, pedido.Estatus); err != nil {
				log.Printf("Error al procesar el recibo: %v", err)
			} else {
				log.Println("Recibo procesado correctamente.")
			}
		}
	}()
	select {}
}
