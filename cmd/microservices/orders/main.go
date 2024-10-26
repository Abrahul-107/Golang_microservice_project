package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting order microservices")

	// Get the context for handling shutdown signals
	ctx := cmd.Context()

	// Initialize the router and closure function for the microservice
	r, closefn := createOrderMicroservices()
	defer closefn() // Ensure resources are properly released on exit

	// Set up the server with address and handler from environment variables
	server := &http.Server{Addr: os.Getenv("Shop_order_service_bind_addr"), Handler: r}

	go func() {

		// Listen and serve, and panic if the server closes unexpectedly
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}

		// Block until context signals shutdown
		<-ctx.Done()

		log.Println("Closing order microservices")
		// Close the server on shutdown and handle any errors
		if err := server.Close(); err != nil {
			panic(err)
		}

	}
}

/*createOrderMicroservices sets up the necessary dependencies and routing for the order service.
Returns the router and a closure function to release resources when the service stops.*/

func createOrderMicroservices() (router *chi.Mux, closeFn func()) {

	// Wait until RabbitMQ service is available at the given address
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	// Initialize an HTTP client for communication with other services
	shopHttpClient := order_infra_product.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))

	// Create a new router instance
	r := cmd.CreateRouter()

	// Initialize order service and repository dependencies
	order_public_http.AddRoutes(r, orderService, ordersRepo)
	order_private_http.AddRoutes(r, orderService, ordersRepo)

	return r, func() {

	}

}
