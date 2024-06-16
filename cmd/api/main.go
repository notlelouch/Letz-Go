package main

import (
	"fmt"
	"net/http"

	"github.com/Letz-Go/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handlers(r)

	fmt.Println("Strating Go api service ...")
	fmt.Println("Alhamdulillah")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}

}
