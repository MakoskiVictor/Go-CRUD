package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Connect establece y devuelve una conexión a la base de datos PostgreSQL.
func Connect() (*sql.DB, error) {
	// Configuración de la cadena de conexión
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"))

	// Abre la conexión
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error al abrir la conexión: %v\n", err)
		return nil, err
	}

	// Prueba la conexión
	err = db.Ping()
	if err != nil {
		log.Printf("Error al probar la conexión: %v\n", err)
		return nil, err
	}

	log.Println("Conexión exitosa a la base de datos")
	return db, nil
}
