package viewmodel

import "github.com/prakriti-yan/webshop/model"

type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

type ProductDetail struct {
	Title   string
	Active  string
	Product Product
}

func NewShopDetail(products []model.Product) ShopDetail {
	result := ShopDetail{
		Title:    "Lemonade Stand Supply - Juice Shop",
		Active:   "shop",
		Products: []Product{},
	}
	for _, p := range products {
		result.Products = append(result.Products, productToVM(p))
	}
	return result
}

func NewShopProduct(product model.Product) ProductDetail {
	tmp := productToVM(product)
	result := ProductDetail{
		Title:   "Lemonade Stand Supply - Juice Shop",
		Active:  "shop",
		Product: tmp,
	}
	return result
}
