// app/handlers/dentista_handler.go
package handlers

import (
	"net/http"
	"encoding/json"
	"path/to/your/models"
	"path/to/your/services"
)

type DentistaHandler struct {
	DentistaService services.DentistaService
}

// ... definir funciones para manejar las rutas relacionadas con Dentistas
