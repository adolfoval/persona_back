package models

type CiudadResponse struct {
	CiuId     string `json:"ciuid" gorm:"column:ciu_id"`
	CiuNombre string `json:"ciunombre" grom:"column:ciu_nombre"`
}

func (CiudadResponse) TableName() string { return "backend.ciudad" }
