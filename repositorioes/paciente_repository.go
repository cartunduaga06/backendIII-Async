// app/repositories/paciente_repository.go
package repositories

import "path/to/your/models"

type PacienteRepository interface {
	GetAll() ([]models.Paciente, error)
	GetByID(id string) (models.Paciente, error)
	// ... otras operaciones específicas del repositorio para Pacientes
}
