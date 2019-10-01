package database.entity
// Identité 
type UserGroup struct{
	gorm.Model
	Group string
	Identitys []Identity `gorm:"foreignkey:GroupRefer"`// liaison avec la struct Identity
}

