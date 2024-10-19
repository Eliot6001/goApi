package main 


import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.setReportCaller(true);
	var r *chi.Mux = chi.NewRouter();
	handlers.Handler(r);


	fmt.Println("Starting GO api");
	fmt.Println(`
	<----------------------- Welcome to this basic Go api ------------------------->
	`);
	err := http.ListenAndServe("localhost:8000", r);
	if err != nil {
		log.Error(err);
	}
}
