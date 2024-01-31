// main.go
package main

import (
	"net/http"
	"path/to/your/app/database"
	"path/to/your/app/handlers"
	"path/to/your/app/services"

	"github.com/gorilla/mux"
)

func main() {
	// Conectar a la base de datos
	database.Connect()

	// Crear instancias de los servicios
	dentistaService := services.NewDentistaService()
	pacienteService := services.NewPacienteService()
	turnoService := services.NewTurnoService()

	// Crear instancias de los controladores
	dentistaHandler := handlers.NewDentistaHandler(dentistaService)
	pacienteHandler := handlers.NewPacienteHandler(pacienteService)
	turnoHandler := handlers.NewTurnoHandler(turnoService)

	// Configurar el enrutador
	r := mux.NewRouter()

	// Rutas para los controladores de Dentistas
	r.HandleFunc("/dentistas", dentistaHandler.GetAllDentistas).Methods("GET")
	r.HandleFunc("/dentistas/{id}", dentistaHandler.GetDentistaByID).Methods("GET")
	// ... otras rutas para Dentistas

	// Rutas para los controladores de Pacientes
	r.HandleFunc("/pacientes", pacienteHandler.GetAllPacientes).Methods("GET")
	r.HandleFunc("/pacientes/{id}", pacienteHandler.GetPacienteByID).Methods("GET")
	// ... otras rutas para Pacientes

	// Rutas para los controladores de Turnos
	r.HandleFunc("/turnos", turnoHandler.GetAllTurnos).Methods("GET")
	r.HandleFunc("/turnos/{id}", turnoHandler.GetTurnoByID).Methods("GET")
	// ... otras rutas para Turnos

	// Iniciar el servidor en el puerto 8080
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
