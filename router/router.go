package router

import (
	"net/http"

	"github.com/ckaminer/obfl-api/stats"
	"github.com/go-chi/chi"
)

func InitializeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", HelloHandler)
	r.Get("/owners", stats.GetAllOwnersHandler)
	return r
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from OBFL"))
}
