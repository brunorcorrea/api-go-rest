package routes

import (
	"log"
	"net/http"

	"github.com/brunorcorrea/go-rest-api/controllers"
	"github.com/brunorcorrea/go-rest-api/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods(http.MethodPost)
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods(http.MethodDelete)
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
