package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"log"
	"database.entity"
)

var DB *gorm.DB
var err error

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
	if DB.HasTable(&Modele{}){
		log.Println("DropTable Modele")
		DB.DropTable(&Modele{})
	}
	if DB.HasTable(&UserGroup{}){
		log.Println("DropTable UserGroup")
		DB.DropTable(&UserGroup{})
	}
	if DB.HasTable(&Agence{}){
		log.Println("DropTable Agence")
		DB.DropTable(&Agence{})
	}
	if DB.HasTable(&Identity{}){
		log.Println("DropTable Identity")
		DB.DropTable(&Identity{})
	}
	if DB.HasTable(&Agent{}){
		log.Println("DropTable Agent")
		DB.DropTable(&Agent{})
	}
	if DB.HasTable(&Status{}){
		log.Println("DropTable Status")
		DB.DropTable(&Status{})
	}
	if DB.HasTable(&Vehicle{}){
		log.Println("DropTable Vehicle")
		DB.DropTable(&Vehicle{})
	}
	if DB.HasTable(&Photo{}){
		log.Println("DropTable Photo")
		DB.DropTable(&Photo{})
	}
	if DB.HasTable(&Historique{}){
		log.Println("DropTable Historique")
		DB.DropTable(&Historique{})
	}


	if !DB.HasTable(&Modele{}){
		log.Println("CreateTable Modele")
		DB.CreateTable(&Modele{})
	}
	if !DB.HasTable(&UserGroup{}){
		log.Println("CreateTable UserGroup")
		DB.CreateTable(&UserGroup{})
		GroupAdmin := UserGroup{Group: "Administrator"}
		GroupUser := UserGroup{Group: "User"}
		DB.Create(&GroupAdmin)
		DB.Create(&GroupUser)
	}

	if !DB.HasTable(&Agence{}){
		log.Println("CreateTable Agence")
		DB.CreateTable(&Agence{})
	}
	if(!DB.HasTable(&Identity{})){
		log.Println("CreateTable Identity")
		DB.CreateTable(&Identity{})
	}
	if !DB.HasTable(&Agent{}){
		log.Println("CreateTable Agent")
		DB.CreateTable(&Agent{})
	}
	if !DB.HasTable(&Status{}){
		log.Println("CreateTable Status")
		DB.CreateTable(&Status{})
		Pret := Status {Value: "prêter"}
		Location := Status{Value: "louer" }
		Demo :=  Status{Value: "démonstration"}
		DB.Create(&Pret)
		DB.Create(&Location)
		DB.Create(&Demo)
	}
	if !DB.HasTable(&Vehicle{}){
		log.Println("CreateTable Vehicle")
		DB.CreateTable(&Vehicle{})
	}
	if !DB.HasTable(&Photo{}){
		log.Println("CreateTable Photo")
		DB.CreateTable(&Photo{})
	}
	if !DB.HasTable(&Historique{}){
		log.Println("CreateTable Historique")
		DB.CreateTable(&Historique{})
	}
	return DB, err
}