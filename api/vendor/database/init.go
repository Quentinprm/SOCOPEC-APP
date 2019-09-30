package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var DB *gorm.DB
var err error

type Post struct {
	gorm.Model
	Author  string
	Message string
}

type Modele struct{
	gorm.Model
	Hauteur float64 
	Largeur float64
	Poid float64
	Puissance int
}

type UserGroup struct{
	gorm.Model
	Group string
}

type Agence struct{
	gorm.Model
	Nom string
	Addresse string
	Tel string
	Fax string
	Photo []byte
}
type Agent struct{
	gorm.Model
	Nom string 
	Prenom string
	Tel string 
	Fax string
	Mobile string
	User string
	Password string
	Group UserGroup
}
type Status struct{
	gorm.Model
	Value string 
}
type Vehicle struct{
	gorm.Model
	Modele Modele
	Date time.Time
	Agence Agence
	Status Status
}

type Photo struct{
	gorm.Model
	Data []byte
	Vehicle Vehicle
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

	// create table if it does not exist
	if DB.HasTable(&Post{}) {
		DB.Delete(&Post{})
	}

	if !DB.HasTable(&Modele{}){
		DB.CreateTable(&Modele{})
	}

	if !DB.HasTable(&UserGroup{}){
		DB.CreateTable(&UserGroup{})
		GroupAdmin := UserGroup{Group: "Administrator"}
		GroupUser := UserGroup{Group: "User"}
		DB.Create(&GroupAdmin)
		DB.Create(&GroupUser)
	}

	if !DB.HasTable(&Agence{}){
		DB.CreateTable(&Agence{})
	}
	if !DB.HasTable(&Agent{}){
		DB.CreateTable(&Agent{})
	}
	if !DB.HasTable(&Status{}){
		DB.CreateTable(&Status{})
		Pret := Status {Value: "prêter"}
		Location := Status{Value: "loue" }
		Demo :=  Status{Value: "démonstration"}
		DB.Create(&Pret)
		DB.Create(&Location)
		DB.Create(&Demo)
	}
	if !DB.HasTable(&Vehicle{}){
		DB.CreateTable(&Vehicle{})
	}
	if !DB.HasTable(&Photo{}){
		DB.CreateTable(&Photo{})
	}
	return DB, err
}