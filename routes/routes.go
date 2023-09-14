package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/aksentijevicd1/reading-from-form-go/handlers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(Router *mux.Router) {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	oh := handlers.NewOpinions(l)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	Router.HandleFunc("/forma", oh.AddOpinion).Methods(http.MethodPost)
	//Router.HandleFunc("/getOpinion", oh.GetOpinion).Methods(http.MethodGet)
	//Router.HandleFunc("/getOpinion", oh.GetSpecificOpinion).Methods(http.MethodGet)

}
