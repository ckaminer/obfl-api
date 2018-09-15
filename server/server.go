package server

import (
	"net/http"
	"os"

	"github.com/ckaminer/obfl/router"
)

func StartServer() {
	r := router.InitializeRouter()
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
