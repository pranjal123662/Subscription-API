package router

import (
	"subscription/helper"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/subscribeapi", helper.Subscribeapi)
	router.HandleFunc("/sendnotification", helper.Sendnotification)
	return router
}
