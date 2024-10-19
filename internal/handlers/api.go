package handlers


import (
	"github.com/Eliot6001/goApi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)


func Handler(r *chi.Mux){
	r.Use(chimiddle.StripSlashes)//removes the last slash; 
	r.Route("/articles", func(router chi.Router){
		router.Use(middleware.Authorization);
		router.Get("/author", GetNumberOfArticles)
	})
}
