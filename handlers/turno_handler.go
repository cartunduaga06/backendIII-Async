// app/handlers/turno_handler.go
package handlers

import (
	"net/http"
	"encoding/json"
	"path/to/your/models"
	"path/to/your/services"
)

type TurnoHandler struct {
	TurnoService services.TurnoService
}

// ... definir funciones para manejar las rutas relacionadas con Turnos
