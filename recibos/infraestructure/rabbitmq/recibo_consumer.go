package rabbitmq

import (
	"encoding/json"
	"log"
	"2da_api_go/recibos/application"
	"github.com/streadway/amqp"
)

type PedidoMensaje struct {
	PedidoID int    `json:"pedido_id"`
	estatus  string `json:"estatus"`
}

func StartRabbitMQListener(useCase *application.CreateRecibo) {
	conn, err := amqp.Dial("amqp://anita:123456789@52.86.221.36:5672/")
	if err != nil {
		log.Fatal("Error al conectar a RabbitMQ: ", err)
	}
	defer conn.Close()

	log.Println("Conexión a RabbitMQ establecida")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error al abrir el canal: ", err)
	}
	defer ch.Close()

	log.Println("Canal abierto correctamente")
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
	} else {
		log.Println("Cola 'pedidos' declarada correctamente")
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
	} else {
		log.Println("Esperando mensajes en la cola 'pedidos'...")
	}
	go func() {
		for d := range msgs {
			log.Printf("Mensaje recibido: %s", d.Body)
			var pedido PedidoMensaje
			if err := json.Unmarshal(d.Body, &pedido); err != nil {
				log.Printf("Error decodificando JSON: %v", err)
				continue
			}
			if err := useCase.Execute(pedido.PedidoID, pedido.estatus); err != nil {
				log.Printf("Error al procesar el recibo: %v", err)
			} else {
				log.Println("Recibo procesado correctamente.")
				log.Println("Esperando mensajes en la cola 'pedidos'...")
			}
		}
	}()
	select {} 
}
