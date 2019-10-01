package entity
import (
	"github.com/jinzhu/gorm"
)
type Status struct{
	gorm.Model
	Value string 
	Vehicles []Vehicle `gorm:"foreignkey:StatusRefer"`//liaison avec la struct Vehicle
	Historiques []Historique `gorm:"polymorphic:Owner;"`// liaison avec la struct Historique
}