package entity
import (
	"github.com/jinzhu/gorm"
)
type Modele struct{
	gorm.Model
	Nom	string
	Hauteur float64 
	Largeur float64
	Poid float64
	Puissance int
	Vehicles []Vehicle `gorm:"foreignkey:ModeleRefer"`//liaison avec la struct Vehicle
	Historiques []Historique `gorm:"polymorphic:Owner;"`//liaison vec la struct Historique
}
