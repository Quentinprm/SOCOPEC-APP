package entity
import (
	"github.com/jinzhu/gorm"
)
type Historique struct{
	gorm.Model
	VehicleRefer uint
	AgentRefer uint
	OwnerId uint
	OwnerType string 

}