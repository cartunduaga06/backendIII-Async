// app/handlers/paciente_handler.go
package handlers

import (
	"net/http"
	"encoding/json"
	"path/to/your/models"
	"path/to/your/services"
)

type PacienteHandler struct {
	PacienteService services.PacienteService
}

// ... definir funciones para manejar las rutas relacionadas con Pacientes
