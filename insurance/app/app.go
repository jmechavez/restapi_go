package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jmechavez/restapi_go/insurance/domain"
	"github.com/jmechavez/restapi_go/insurance/service"
)

func Start() {
	router := mux.NewRouter()

	// ch := ClientHandlers{service.NewClientService(domain.NewClientRepositoryStub())}
	ch := ClientHandlers{service.NewClientService(domain.NewClientRepositoryDb())}
	// define routes
	router.HandleFunc("/clients", ch.getAllClient).Methods(http.MethodGet)
	router.HandleFunc("/clients/{fname}", ch.FindName).Methods(http.MethodGet)
	// router.HandleFunc("/clients/{fname}", ch.JustFname).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
