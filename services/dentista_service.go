// app/services/dentista_service.go
package services

import "path/to/your/models"

type DentistaService interface {
	GetAll() ([]models.Dentista, error)
	GetByID(id string) (models.Dentista, error)
	// ... otras operaciones específicas del servicio para Dentistas
}
