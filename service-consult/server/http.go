package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"service-consult/repository"
)

type httpServer struct {
	http.Handler
}

func NewServer() Storage {
	return new(httpServer)
}

//Init new routes
func (h *httpServer) NewRoutes() {
	log.Println("Init Routes")
	router := http.NewServeMux()

	//Create the endpoins
	router.Handle("/", http.HandlerFunc(jsonOperator))

	h.Handler = router
}

//Run server
func (h *httpServer) StartAPI() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Println("Start API")
	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, h); err != nil {
		log.Fatal("init server error in StartApi(), ", err)
	}
}

func jsonOperator(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		cpf := struct {
			CPF string `json:"cpf"`
		}{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		json.Unmarshal(body, &cpf)

		db, err := repository.OpenConnection()
		if err != nil {
			log.Println("Error OpenConnection(), ", err)
		}

		repo := repository.NewRepository(db)

		cval, err := repo.ConsultNegativacoes(cpf.CPF)
		if err != nil {
			log.Println("Error ConsultNegativacoes(), ", err)
		}

		json.NewEncoder(w).Encode(cval)
		db.Close()
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
