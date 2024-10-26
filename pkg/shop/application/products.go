package application

type productReadModel interface {
	AllProducts() ([]products.Product, err)
}

type productsService struct {
	repo      products.respository
	readmodel productReadModel
}

func newProductsServices() productsService {

}

func (s productsService) AllProducts() {

}

func (s productsService) AddProduct() error {

}
