package models

type VehicleType string

const (
	Motor VehicleType = "motor"
	Car               = "car"
)

type Vehicle struct {
	Model

	RegNumber    string      `gorm:"not null;unique" json:"regNumber"`
	VehicleModel string      `json:"vehicleModel"`
	Type         VehicleType `json:"vehicleType"`
	ElectronID   uint        `json:"electronId"`

	State  State `json:"state"`
	Active bool  `json:"active" gorm:"default=false"`
}
