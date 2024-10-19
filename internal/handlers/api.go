package handlers


import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)


func Handler(r *chi.Mux){
	r.use(chimiddle.StripSlashes);
	r.Route("/articles", func(router chil.Router){
		router.Use(middleware.Authorization);
		router.Get("/author", GetNumberOfArticles)
	})
}
