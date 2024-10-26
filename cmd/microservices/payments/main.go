package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Starting payment microservices")

	defer log.Println("Closing payments microservices")

	ctx := cmd.Context()

	paymentInterface := createPaymentMicroservice()

	if err := paymentInterface.Run(ctx); err != nil {
		panic(err)
	}
}

func createPaymentMicroservice() amqp.paymentInterface{
	cmd.WaitForSe+rvice(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsServices := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)

	paymentInterface , err := amqp.NewPaymentsInterface(
		fmt.Sprintf("amqp://%s/",os.Getenv("SHOP_RABBITMQ_ADDR"))
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		paymentsServices,
	)
}
