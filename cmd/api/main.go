package main 


import (
	"fmt"
	"net/http"
	"github.com/Eliot6001/goApi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true);
	var r *chi.Mux = chi.NewRouter();
	handlers.Handler(r) //passing request to handlers;


	fmt.Println("Starting GO api");
	fmt.Println(`
	<----------------------- Welcome to this basic Go api ------------------------->
	`);
	err := http.ListenAndServe("localhost:8000", r);
	if err != nil {
		log.Error(err);
	}
}
