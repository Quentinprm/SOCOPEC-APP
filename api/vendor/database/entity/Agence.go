package database.entity
type Agence struct{
	gorm.Model
	Nom string
	Addresse string
	Tel string
	Fax string
	Photo Photo `gorm:"polymorphic:Owner;"` // liaison avec la struct Photo
	Vehicles []Vehicle `gorm:"foreignkey:AgenceRefer"`//liaison avec la struct Vehicle
	Historiques []Historique `gorm:"polymorphic:Owner;"`//liaison vec la struct Historique
}
