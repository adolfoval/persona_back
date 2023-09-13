package models

import (
	"persona/db"

	"gorm.io/gorm"
)

type Ciudad struct {
	gorm.Model
	CiuId     string `json:"ciuid" gorm:"column:ciu_id;<-create;primaryKey;autoIncrement"`
	CiuNombre string `json:"ciunombre" gorm:"column:ciu_nombre"`
	DepId     string `json:"depid" gorm:"column:dep_id"`
}

type Persona struct {
	gorm.Model
	ID                int     `json:"id" gorm:"column:id"`
	PerId             int64   `json:"perid" gorm:"column:per_id;<-create;primaryKey;autoIncrement"`
	PerIdentificacion *string `json:"peridentificacion" gorm:"type:varchar(15);column:per_identificacion;index:per_idenx,unique"`
	PerNombre1        *string `json:"pernombre1" gorm:"type:varchar(20);column:per_nombre1;not null"`
	PerNombre2        string  `json:"pernombre2" gorm:"type:varchar(20);column:per_nombre2"`
	PerApellido1      *string `json:"perapellido1" gorm:"type:varchar(25);column:per_apellido1; not null"`
	PerApellido2      *string `json:"perapellido2" gorm:"type:varchar(25);column:per_apellido2; not null"`
	PerCorreo         string  `json:"percorreo" gorm:"type:varchar(60);column:per_correo"`
	PerTelefono       string  `json:"pertelefono" gorm:"type:varchar(20);column:per_telefono"`
	PerEstado         string  `json:"perestado" gorm:"type:varchar(1); not null; column:per_estado"`
	CiuIdRef          string  `json:"ciuid" gorm:"type:varchar(5);column:ciu_id;"`
	CiuIdFK           Ciudad  `json:"ciu" gorm:"ForeignKey:CiuIdRef;references:CiuId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type PersonaSelect struct {
	PerIdentificacion *string `json:"peridentificacion" gorm:"type:varchar(15);column:per_identificacion;index:per_idenx,unique"`
	PerNombre1        *string `json:"pernombre1" gorm:"type:varchar(20);column:per_nombre1;not null"`
	PerNombre2        string  `json:"pernombre2" gorm:"type:varchar(20);column:per_nombre2"`
	PerApellido1      *string `json:"perapellido1" gorm:"type:varchar(25);column:per_apellido1; not null"`
	PerApellido2      *string `json:"perapellido2" gorm:"type:varchar(25);column:per_apellido2; not null"`
	PerCorreo         string  `json:"percorreo" gorm:"type:varchar(60);column:per_correo"`
	PerTelefono       string  `json:"pertelefono" gorm:"type:varchar(20);column:per_telefono"`
	PerEstado         string  `json:"perestado" gorm:"type:varchar(1); not null; column:per_estado"`
	DepartamentoId    string  `json:"depid" gorm:"column:dep_id"`
	CiuId             string  `json:"ciuid" gorm:"column:ciu_id;<-create;primaryKey;autoIncrement"`
	CiuNombre         string  `json:"ciunombre" gorm:"column:ciu_nombre"`
}

func (PersonaSelect) TableName() string { return "backend.persona" }

func MigratePerson() {
	db.Database.AutoMigrate(Persona{})
}
