// app/models/paciente.go
package models

type Paciente struct {
	ID         string `json:"id"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Domicilio  string `json:"domicilio"`
	DNI        string `json:"dni"`
	FechaAlta  string `json:"fechaAlta"`
}
