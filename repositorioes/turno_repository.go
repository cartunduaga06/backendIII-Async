// app/repositories/turno_repository.go
package repositories

import "path/to/your/models"

type TurnoRepository interface {
	GetAll() ([]models.Turno, error)
	GetByID(id string) (models.Turno, error)
	// ... otras operaciones específicas del repositorio para Turnos
}
