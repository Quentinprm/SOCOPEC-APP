package database.entity
type Historique struct{
	gorm.Model
	VehicleRefer uint
	AgentRefer uint
	OwnerId uint
	OwnerType string 

}