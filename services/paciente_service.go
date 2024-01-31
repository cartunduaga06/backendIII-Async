// app/services/paciente_service.go
package services

import "path/to/your/models"

type PacienteService interface {
	GetAll() ([]models.Paciente, error)
	GetByID(id string) (models.Paciente, error)
	// ... otras operaciones específicas del servicio para Pacientes
}
