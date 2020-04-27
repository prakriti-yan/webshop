package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/prakriti-yan/webshop/viewmodel"
)

type locator struct {
	standLocatorTemplate *template.Template
}

func (l locator) registerRoutes() {
	http.HandleFunc("/stand_locator", l.handleLocator)
	http.HandleFunc("/api/stands", l.handleApiStands)
}

func (l locator) handleLocator(w http.ResponseWriter, r *http.Request) {
	vw := viewmodel.NewStandLocator()
	w.Header().Add("Content-Type", "text/html")
	l.standLocatorTemplate.Execute(w, vw)
}

func (l locator) handleApiStands(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var loc struct {
		ZipCode string `json: "zipCode"`
	}
	err := dec.Decode((&loc))
	if err != nil {
		log.Println(fmt.Errorf("Error retrieving location: %v", err))
		enc := json.NewEncoder(w)
		enc.Encode([]viewmodel.StandCoordinate{})
		return
	}
	log.Println("Location: ", loc)
	vm := coords
	enc := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json")
	enc.Encode(vm)
	// fmt.Println(coords)
}

var coords []viewmodel.StandCoordinate = []viewmodel.StandCoordinate{
	viewmodel.StandCoordinate{
		Latitude:  60.205,
		Longitude: 25.109,
		Title:     "Rebecca's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  60.220,
		Longitude: 25.109,
		Title:     "Chris's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  60.208,
		Longitude: 25.108,
		Title:     "Carson's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  60.210,
		Longitude: 25.107,
		Title:     "Lorelei's stand",
	},
}
