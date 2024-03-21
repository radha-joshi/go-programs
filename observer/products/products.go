package products

type Product struct {
	name string
}

type ProductList struct {
	list          []Product
	observersList []ProductListObserver
}

func (p *ProductList) AddObserver(observer ProductListObserver) {
	p.observersList = append(p.observersList, observer)
}

func (p *ProductList) AddProduct(product Product) {
	p.list = append(p.list, product)
	for _, item := range p.observersList {
		item.ProductAdded(product)
	}
}

type ProductListObserver interface {
	ProductAdded(product Product)
}

func NewProductList() ProductList {
	return ProductList{
		list:          make([]Product, 0),
		observersList: make([]ProductListObserver, 0),
	}
}

func NewProduct(name string) Product {
	return Product{
		name: name,
	}
}
