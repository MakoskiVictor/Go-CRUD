package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MakoskiVictor/Go-CRUD/db"
	"github.com/MakoskiVictor/Go-CRUD/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Conectar a la DB
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	defer database.Close() // Cierra la conexión a la base de datos al finalizar la función main

	fmt.Println("Conexión exitosa a la base de datos")

	r := mux.NewRouter() // Crea un nuevo router con gorilla/mux

	// Crea una ruta que responde a la raíz del servidor
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r) // Inicia el server en el puerto y utiliza r (el router creado)
}
