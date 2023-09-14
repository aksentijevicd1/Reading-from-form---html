package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aksentijevicd1/reading-from-form-go/models"
)

type Opinions struct {
	l *log.Logger
}

func NewOpinions(l *log.Logger) *Opinions {
	return &Opinions{l}
}

func (o *Opinions) AddOpinion(w http.ResponseWriter, r *http.Request) {
	var newOpinion models.Opinion

	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		o.l.Printf("ParseForm() error:%s", err)
		return
	}

	newOpinion.FirstName = r.FormValue("firstName")
	newOpinion.LastName = r.FormValue("lastName")
	newOpinion.Address = r.FormValue("address")
	newOpinion.Opinion = r.FormValue("opinion")

	newOpinion.AddOpinion()
	fmt.Fprintf(w, "Name = %s\n", r.FormValue("firstName"))
	fmt.Fprintf(w, "Address = %s\n", r.FormValue("address"))

}

func (o *Opinions) GetOpinions(w http.ResponseWriter, r *http.Request) {

	allOpinions := models.GetOpinions()
	res, err := json.Marshal(allOpinions)

	if err != nil {
		http.Error(w, "error while marshaling", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
