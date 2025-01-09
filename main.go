package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter() // Crea un nuevo router

	// Crea una ruta que responde a la raíz del servidor
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":3000", r) // Inicia el server en el puerto y utiliza r (el router creado)
}
