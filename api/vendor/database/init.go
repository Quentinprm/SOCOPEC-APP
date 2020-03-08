package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var DB *gorm.DB
var err error


type Agency struct {
	ID              	uint `gorm:"primary_key"`
	Name 				string
	Adress 				string
	City				string
	PhoneNumber 		string
	Fax 				string
	Picture 			string
}

type User struct{
	ID              	uint `gorm:"primary_key"`
	Email				string
	FirstName			string
	LastName			string
	Password			string
	PhoneNumber			string
	Fax					string
	Mobile				string
	AgencyID			uint
	IsAdmin				bool
	AddRight			bool
}

type State struct {
	ID              	uint `gorm:"primary_key"`
	Libelle 			string
}

type Car struct{
	ID              	uint `gorm:"primary_key"`
	Immatriculation 	string
	ModelName			string
	Year				int
	Weight				int
	Power 				float32
	Length				int
	Width				int
	Height				int
	Picture				string
	AgencyID 			uint
}

type CarState struct{
	ID              	uint `gorm:"primary_key"`
	CarID 				uint
	StateID				uint
	UserID				uint
	DateReservation 	time.Time
	BeginDate			time.Time
	EndDate				time.Time
}

func addDatabase(dbname string) error {
	// create database with dbname, won't do anything if db already exists
	DB.Exec("CREATE DATABASE " + dbname)

	// connect to newly created DB (now has dbname param)
	connectionParams := "dbname=" + dbname + " user=docker password=docker sslmode=disable host=db"
	DB, err = gorm.Open("postgres", connectionParams)
	if err != nil {
		return err
	}

	return nil
}

func Init() (*gorm.DB, error) {
	// set up DB connection and then attempt to connect 5 times over 25 seconds
	connectionParams := "user=docker password=docker sslmode=disable host=db"
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open("postgres", connectionParams) // gorm checks Ping on Open
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return DB, err
	}


	// Angency Base init 
	if !DB.HasTable(&Agency{}){
		DB.CreateTable(&Agency{})
		BaseAgency := Agency{Name: "Nancy",Adress:"4 terrasse des vosges", City:"Nancy",PhoneNumber:"0695348705",Fax:"0695348706",Picture:""}
		DB.Create(&BaseAgency)
	}

	
	if !DB.HasTable(&Car{}){
		DB.CreateTable(&Car{})
		DB.Model(&Car{}).AddForeignKey("agency_id","agencies(id)","CASCADE","CASCADE")
		BaseCar1 := Car{Immatriculation:"HGJVBFGH1",ModelName:"toyota yaris",Year:2007,Weight:9556,Power:70,Length:265,Width:124,Height:215,Picture:"",AgencyID:1}
		BaseCar2 := Car{Immatriculation:"HGJVBFGH2",ModelName:"scoda fabia",Year:2007,Weight:9556,Power:70,Length:265,Width:124,Height:215,Picture:"",AgencyID:1}
		DB.Create(&BaseCar1)
		DB.Create(&BaseCar2)
	}

	// Status Base init 
	if !DB.HasTable(&State{}){
		DB.CreateTable(&State{})
		BaseStatusPreter:= State{Libelle:"Prêter"}
		BaseStatusLouer := State{Libelle:"Louer"}
		BaseStatusDemo	:= State{Libelle:"Démonstration"}
		DB.Create(&BaseStatusPreter)
		DB.Create(&BaseStatusLouer)
		DB.Create(&BaseStatusDemo)
	}
	

	if !DB.HasTable(&User{}){
		DB.CreateTable(&User{})
		DB.Model(&User{}).AddForeignKey("agency_id","agencies(id)","CASCADE","CASCADE")
		AdminUser := User{	Email:"admin@admin.com",FirstName:"admin",LastName:"admin",Password:"admin",PhoneNumber:"0654875214",Fax:"0356892510",Mobile:"0654875214",IsAdmin:true,AddRight:true,AgencyID:1}
		DB.Create(&AdminUser)
	}

	if !DB.HasTable(&CarState{}){
		DB.CreateTable(&CarState{})
		DB.Model(&CarState{}).AddForeignKey("car_id","cars(id)","CASCADE","CASCADE")
		DB.Model(&CarState{}).AddForeignKey("user_id","users(id)","RESTRICT","CASCADE")
		DB.Model(&CarState{}).AddForeignKey("state_id","states(id)","RESTRICT","CASCADE")
	}
	
	return DB, err
}
