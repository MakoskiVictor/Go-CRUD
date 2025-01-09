package main

import (
	"net/http"

	"github.com/MakoskiVictor/Go-CRUD/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter() // Crea un nuevo router con gorilla/mux

	// Crea una ruta que responde a la raíz del servidor
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r) // Inicia el server en el puerto y utiliza r (el router creado)
}
