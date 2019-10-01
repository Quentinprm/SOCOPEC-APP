package entity
import (
	"github.com/jinzhu/gorm"
)
type Photo struct{
	gorm.Model
	Data []byte
	OwnerID uint //liaison 
	OwnerType string //liaison
	Historiques []Historique `gorm:"polymorphic:Owner;"`//liaison vec la struct Historique
}
