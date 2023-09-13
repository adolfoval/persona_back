package models

type Departamento struct {
	DepartamentoId     string `json:"depid" gorm:"column:dep_id"`
	DepartamentoNombre string `json:"depnombre" gorm:"column:dep_nombre"`
}

func (Departamento) TableName() string { return "backend.departamento" }
