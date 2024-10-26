package main

import (
	"log"
	"net/http"
	"os"
)

func main(){

	log.Println("Starting shop microservices")

	ctx := cmd.Context()

	r := createShopMicroservices()

	server := &http.Server{Addr:os.Getenv("SHOP_PRODUCT_SERVICE_BIND_ADDR"),Handler:r}

	go func(){
		if err:=server.ListenAndServe(); err != http.ErrServerClosed{
			panic(err)
		}
	}()
	<= ctx.Done()
	log.Println("Closing shop microservices")
	if err := server.Close(); err != nil {
		panic(err)
	}
	
}

func createShopMicroservices()*chi.Mux{

	shopProductRepo := shop_infra_product.NewMemoryRepository()

	r := cmd.CreateRouter()

	shop_interfaces_public_http.AddRoutes(r,shopProductRepo)
	shop_interfaces_private_http.AddRoutes(r,shopProductRepo)

	return r,func ()  {
		
	}
}