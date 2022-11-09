package main

import (
	"devfullcycle/gointensivo/internal/database_pg"
	"devfullcycle/gointensivo/internal/order/infra/database"
	"devfullcycle/gointensivo/internal/order/usecase"
	"devfullcycle/gointensivo/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// {"id":"abc","price":10.0,"tax":1.0}
func main() {
	database_pg.ConectaComBancoDeDados()

	//db, err := sql.Open("sqlite3", "./orders.db")
	db := database_pg.DB
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, out)

	for msg := range out {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println(outputDTO)
		//time.Sleep(500 * time.Millisecond)
	}

}

func main2() {

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, out)

	for msg := range out {
		msg.Ack(false)
		println(string(msg.Body))
	}

}
