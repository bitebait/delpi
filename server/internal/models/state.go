package models

type Departamento struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre"`
	GPS    GPS    `json:"gps" gorm:"embedded"`
	Dato   string `json:"dato"`
}

func (Departamento) TableName() string {
	return "departamentos"
}
