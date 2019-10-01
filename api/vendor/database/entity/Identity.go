package database.entity
type Identity struct{
	gorm.Model
	User string `gorm:"type:varchar(50);unique_index"`
	Password string
	GroupRefer uint // liaison avec UserGoup
}
