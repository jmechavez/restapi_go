package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jmechavez/restapi_go/insurance/service"
)

type ClientHandlers struct {
	service service.ClientService
}

type FnameResponse struct {
	Fname string `json:"fname"`
}

func (ch *ClientHandlers) getAllClient(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	clients, err := ch.service.GetAllClient(status)

	// if xml format if not json format
	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(clients)
	// } else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(clients)
	// }
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, clients)
	}
}

func (ch *ClientHandlers) FindName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fname := vars["fname"]

	client, err := ch.service.FindName(fname)
	if err != nil {
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(err.Code)
		// json.NewEncoder(w).Encode(err.AsMessage())
		// fmt.Fprint(w, err.Message)
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, client)
	}
}

func (ch *ClientHandlers) JustFname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fname := vars["fname"]

	client, err := ch.service.FindName(fname)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		response := FnameResponse{
			Fname: client.Fname,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	// shorter format code
	// if err := json.NewEncoder(w).Encode(data); err != nil {
	// 	panic(err)
	// }

	// long format
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(fmt.Sprintf("failed to encode response: %v", err))
	}
}
