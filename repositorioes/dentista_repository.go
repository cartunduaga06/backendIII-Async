// app/repositories/dentista_repository.go
package repositories

import "path/to/your/models"

type DentistaRepository interface {
	GetAll() ([]models.Dentista, error)
	GetByID(id string) (models.Dentista, error)
	// ... otras operaciones específicas del repositorio para Dentistas
}
