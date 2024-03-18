package main

import (
	"github.com/luisasacramento/web-app/routes"
	"net/http"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
