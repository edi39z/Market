package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	fmt.Println("Test server")

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}
func (server *Server) Run(addr string) {
	fmt.Printf("Serverto to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run(){
	var server = Server {}

	server.Initialize()
	server.Run(":9000")

}
