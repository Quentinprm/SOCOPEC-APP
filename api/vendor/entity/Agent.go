package entity
import (
	"github.com/jinzhu/gorm"
)
type Agent struct{
	gorm.Model
	Nom string 
	Prenom string
	Tel string 
	Fax string
	Mobile string
	Identity Identity `gorm:"foreignkey:IdentityRefer"`// liaison avec la struct Identity
	IdentityRefer uint 
	Historiques []Historique  `gorm:"foreignkey:AgentRefer"` // liaison avec la struct Historique
}
