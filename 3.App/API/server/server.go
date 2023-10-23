package server

import (
	"api/model"
	"api/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router  *mux.Router
	service *service.Service
}

func New(s *service.Service) Server {
	server := &Server{service: s}

	router := mux.NewRouter()
	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")

			h.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/", server.getMunicipios).Methods(http.MethodGet)

	//municipio
	router.HandleFunc("/municipios", server.getMunicipios).Methods(http.MethodGet) //.Headers("Content-Type", "application/json")

	//parqueadero
	router.HandleFunc("/parqueaderos", server.GetParqueaderos).Methods(http.MethodGet) //.Headers("Content-Type", "application/json")
	router.HandleFunc("/parqueaderos", server.CreateParqueaderos).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	//buses
	router.HandleFunc("/buses", server.GetBuses).Methods(http.MethodGet) //.Headers("Content-Type", "application/json")
	router.HandleFunc("/buses", server.CreateBuses).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	//programas
	router.HandleFunc("/programas", server.GetProgramaciones).Methods(http.MethodGet) //.Headers("Content-Type", "application/json")
	router.HandleFunc("/programas", server.CreateProgramacion).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	server.Router = router
	return *server
}

func (s *Server) helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello wolrd!"))
}

func (s *Server) getMunicipios(w http.ResponseWriter, r *http.Request) {
	municipios, err := s.service.GetMunicipios()
	if err != nil {
		w.WriteHeader(401)
		return
	}

	//convertir municipios a json
	responseJson(w, municipios)
}

func (s *Server) GetParqueaderos(w http.ResponseWriter, r *http.Request) {
	parqueaderos, err := s.service.GetParqueaderos()
	if err != nil {
		w.WriteHeader(401)
		return
	}

	//convertir municipios a json
	responseJson(w, parqueaderos)
}

func (s *Server) CreateParqueaderos(w http.ResponseWriter, r *http.Request) {
	var parqueadero model.Parqueadero
	err := json.NewDecoder(r.Body).Decode(&parqueadero)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	parqueaderos, err := s.service.CreateParqueadero(parqueadero)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	//convertir municipios a json
	responseJson(w, parqueaderos)
}

func (s *Server) GetBuses(w http.ResponseWriter, r *http.Request) {
	buses, err := s.service.GetBuses()
	if err != nil {
		w.WriteHeader(401)
		return
	}

	//convertir municipios a json
	responseJson(w, buses)
}

func (s *Server) CreateBuses(w http.ResponseWriter, r *http.Request) {
	var bus model.Bus
	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	buses, err := s.service.CreateBus(bus)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	//convertir municipios a json
	responseJson(w, buses)
}

func (s *Server) GetProgramaciones(w http.ResponseWriter, r *http.Request) {
	programas, err := s.service.GetProgramaciones()

	if err != nil {
		w.WriteHeader(401)
		return
	}

	//convertir municipios a json
	responseJson(w, programas)
}

func (s *Server) CreateProgramacion(w http.ResponseWriter, r *http.Request) {
	var programa model.Programacion
	err := json.NewDecoder(r.Body).Decode(&programa)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	buses, err := s.service.CreateProgramacion(programa)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	//convertir municipios a json
	responseJson(w, buses)
}

func responseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
