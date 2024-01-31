// app/services/turno_service.go
package services

import "path/to/your/models"

type TurnoService interface {
	GetAll() ([]models.Turno, error)
	GetByID(id string) (models.Turno, error)
	// ... otras operaciones específicas del servicio para Turnos
}
