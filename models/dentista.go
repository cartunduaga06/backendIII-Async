// app/models/dentista.go
package models

type Dentista struct {
	ID        string `json:"id"`
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
	Matricula string `json:"matricula"`
}
