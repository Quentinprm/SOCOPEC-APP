package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"log"
	"entity"
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
	if DB.HasTable(&entity.Modele{}){
		log.Println("DropTable Modele")
		DB.DropTable(&entity.Modele{})
	}
	if DB.HasTable(&entity.UserGroup{}){
		log.Println("DropTable UserGroup")
		DB.DropTable(&entity.UserGroup{})
	}
	if DB.HasTable(&entity.Agence{}){
		log.Println("DropTable Agence")
		DB.DropTable(&entity.Agence{})
	}
	if DB.HasTable(&entity.Identity{}){
		log.Println("DropTable Identity")
		DB.DropTable(&entity.Identity{})
	}
	if DB.HasTable(&entity.Agent{}){
		log.Println("DropTable Agent")
		DB.DropTable(&entity.Agent{})
	}
	if DB.HasTable(&entity.Status{}){
		log.Println("DropTable Status")
		DB.DropTable(&entity.Status{})
	}
	if DB.HasTable(&entity.Vehicle{}){
		log.Println("DropTable Vehicle")
		DB.DropTable(&entity.Vehicle{})
	}
	if DB.HasTable(&entity.Photo{}){
		log.Println("DropTable Photo")
		DB.DropTable(&entity.Photo{})
	}
	if DB.HasTable(&entity.Historique{}){
		log.Println("DropTable Historique")
		DB.DropTable(&entity.Historique{})
	}


	if !DB.HasTable(&entity.Modele{}){
		log.Println("CreateTable Modele")
		DB.CreateTable(&entity.Modele{})
	}
	if !DB.HasTable(&entity.UserGroup{}){
		log.Println("CreateTable UserGroup")
		DB.CreateTable(&entity.UserGroup{})
		GroupAdmin := entity.UserGroup{Group: "Administrator"}
		GroupUser := entity.UserGroup{Group: "User"}
		DB.Create(&GroupAdmin)
		DB.Create(&GroupUser)
	}

	if !DB.HasTable(&entity.Agence{}){
		log.Println("CreateTable Agence")
		DB.CreateTable(&entity.Agence{})
	}
	if(!DB.HasTable(&entity.Identity{})){
		log.Println("CreateTable Identity")
		DB.CreateTable(&entity.Identity{})
	}
	if !DB.HasTable(&entity.Agent{}){
		log.Println("CreateTable Agent")
		DB.CreateTable(&entity.Agent{})
	}
	if !DB.HasTable(&entity.Status{}){
		log.Println("CreateTable Status")
		DB.CreateTable(&entity.Status{})
		Pret := entity.Status{Value: "prêter"}
		Location := entity.Status{Value: "louer" }
		Demo :=  entity.Status{Value: "démonstration"}
		DB.Create(&Pret)
		DB.Create(&Location)
		DB.Create(&Demo)
	}
	if !DB.HasTable(&entity.Vehicle{}){
		log.Println("CreateTable Vehicle")
		DB.CreateTable(&entity.Vehicle{})
	}
	if !DB.HasTable(&entity.Photo{}){
		log.Println("CreateTable Photo")
		DB.CreateTable(&entity.Photo{})
	}
	if !DB.HasTable(&entity.Historique{}){
		log.Println("CreateTable Historique")
		DB.CreateTable(&entity.Historique{})
	}
	return DB, err
}