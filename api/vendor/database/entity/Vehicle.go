package database.entity
type Vehicle struct{
	gorm.Model
	Date *time.Time
	ModeleRefer uint //liaison avec la struct Modele
	AgenceRefer uint	//liaison avec la struct Agence
	StatusRefer uint	// liaison avec la struct Status
	Photos []Photo `gorm:"polymorphic:Owner;"`// liaison avec la struct Photo
	Historiques []Historique  `gorm:"foreignkey:VehicleRefer"` // liaison avec la struct Historique
}
