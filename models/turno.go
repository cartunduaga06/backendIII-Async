// app/models/turno.go
package models

type Turno struct {
	ID          string `json:"id"`
	FechaHora   string `json:"fechaHora"`
	Descripcion string `json:"descripcion"`
	Paciente    Paciente `json:"paciente"`
	Dentista    Dentista `json:"dentista"`
}

