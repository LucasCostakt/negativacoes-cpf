package server

import (
	"io"
	"log"
	"net/http"
	"os"
	"service-import-data/repository"
	"service-import-data/service"
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
		port = "5050"
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
	
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Println("FormFile error in csvOperator(), ", err)
		}
		defer file.Close()
		
		f, err := os.OpenFile("archives/"+handler.Filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Println("open file error in  csvOperator(), ", err)
		}
		defer f.Close()
		
		_, err = io.Copy(f, file)
		if err != nil {
			log.Println("open file error in csvOperator(), ", err)
		}

		data := &service.Negativacoes{}

		ng := service.NewNegativacoes(data.Data)

		json, err := ng.ConvertFileToStruct("archives/" + handler.Filename)
		if err != nil {
			log.Println("open file error in csvOperator(), ", err)
		}

		db, err := repository.OpenConnection()
		if err != nil {
			log.Println("open file error in  csvOperator(), ", err)
		}

		repo := repository.NewRepository(db)


		err = repo.SaveJson(json)
		if err != nil {
			log.Println("open file error in  csvOperator(), ", err)
		}

		_, err = io.WriteString(w, "File Uploaded successfully\n")
		if err != nil {
			log.Println(" write http.ResponseWriter error in csvOperator(), ", err)
		}
		db.Close()
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
