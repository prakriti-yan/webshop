package controller

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/prakriti-yan/webshop/model"
	"github.com/prakriti-yan/webshop/viewmodel"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
	productTemplate  *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	http.HandleFunc("/shop/", s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	productPattern, _ := regexp.Compile(`/shop/(\d+)/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	matchesP := productPattern.FindStringSubmatch(r.URL.Path)
	if len(matchesP) > 0 {
		categoryID, _ := strconv.Atoi(matchesP[1])
		ID, _ := strconv.Atoi(matchesP[2])
		s.handleProduct(w, r, categoryID, ID)
	} else if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		s.handleCategory(w, r, categoryID)
	} else {
		categories := model.GetCategories()
		vw := viewmodel.NewShop(categories)
		s.shopTemplate.Execute(w, vw)
	}
}

func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	products := model.GetProductsForCategory(categoryID)
	vw := viewmodel.NewShopDetail(products)
	s.categoryTemplate.Execute(w, vw)
}

func (s shop) handleProduct(w http.ResponseWriter, r *http.Request, categoryID int, ID int) {
	product := model.GetProduct(categoryID, ID)
	vw := viewmodel.NewShopProduct(product)
	s.productTemplate.Execute(w, vw)
}
