package entity
import (
	"github.com/jinzhu/gorm"
)
type UserGroup struct{
	gorm.Model
	Group string
	Identitys []Identity `gorm:"foreignkey:GroupRefer"`// liaison avec la struct Identity
}

