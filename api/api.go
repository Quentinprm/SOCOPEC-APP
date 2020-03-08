package main

import (
	"database"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)
//Agency Handler

func indexAgencyHandler(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	setCors(w)
	var agencies[] database.Agency
	database.DB.Find(&agencies)
	res,err := json.Marshal(agencies)
	if err != nil{
		http.Error(w,err.Error(),500)
	}
	w.Write(res)
}
func showAgencyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var agency database.Agency
	database.DB.Where("ID = ?", ps.ByName("agencyId")).First(&agency)
	res, err := json.Marshal(agency)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(res)
}
func createAgencyHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	setCors(w)
	decoder := json.NewDecoder(r.Body)
	var newAgency database.Agency

	if err:= decoder.Decode(&newAgency);err != nil {
		http.Error(w,err.Error(),400)
		return
	}

	database.DB.Create(&newAgency)
	res,err := json.Marshal(newAgency)
	if err != nil {
		http.Error(w,err.Error(),500)
		return
	}
	w.Write(res)
}
func updateAgencyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	type body struct {
		Name 		string
		Adress 		string
		City		string
		PhoneNumber string
		Fax 		string
		Picture 	string
	}
	var updates body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updates); err != nil {
		http.Error(w, err.Error(), 400)
	}

	var updatedAgency database.Agency
	database.DB.Where("ID = ?", ps.ByName("agencyId")).First(&updatedAgency)

	updatedAgency.Name = updates.Name
	updatedAgency.Adress = updates.Adress
	updatedAgency.City= updates.City
	updatedAgency.PhoneNumber = updates.PhoneNumber
	updatedAgency.Fax = updates.Fax
	updatedAgency.Picture = updates.Picture

	database.DB.Save(&updatedAgency)
	res, err := json.Marshal(updatedAgency)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(res)
}
func deleteAgencyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var deletedAgency database.Agency
	database.DB.Where("ID = ?", ps.ByName("agencyId")).Delete(&deletedAgency)
	res, err := json.Marshal(deletedAgency)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(res)
}

//User Handler

func indexUserHandler(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	setCors(w)
	var users[] database.User
	database.DB.Find(&users)
	res,err := json.Marshal(users)
	if err != nil{
		http.Error(w,err.Error(),500)
	}
	w.Write(res)
}
func showUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var user database.User
	database.DB.Where("ID = ?", ps.ByName("userId")).First(&user)
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(res)
}
func createUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	setCors(w)
	decoder := json.NewDecoder(r.Body)
	var newUser database.User

	if err:= decoder.Decode(&newUser);err != nil {
		http.Error(w,err.Error(),400)
		return
	}

	database.DB.Create(&newUser)
	res,err := json.Marshal(newUser)
	if err != nil {
		http.Error(w,err.Error(),500)
		return
	}
	w.Write(res)
}
func updateUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	type body struct {
		Email			string
		FirstName		string
		LastName		string
		Password		string
		PhoneNumber		string
		Fax				string
		Mobile			string
		AgencyID		uint
		IsAdmin			bool
		AddRight		bool
	}
	var updates body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updates); err != nil {
		http.Error(w, err.Error(), 400)
	}

	var updatedUser database.User
	database.DB.Where("ID = ?", ps.ByName("userId")).First(&updatedUser)

	updatedUser.Email = updates.Email
	updatedUser.FirstName = updates.FirstName
	updatedUser.LastName = updates.LastName
	updatedUser.Password = updates.Password
	updatedUser.PhoneNumber = updates.PhoneNumber
	updatedUser.Fax = updates.Fax
	updatedUser.Mobile = updates.Mobile
	updatedUser.AgencyID = updates.AgencyID
	updatedUser.IsAdmin = updates.IsAdmin
	updatedUser.AddRight = updates.AddRight


	database.DB.Save(&updatedUser)
	res, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(res)
}
func deleteUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var deletedUser database.User
	database.DB.Where("ID = ?", ps.ByName("userId")).Delete(&deletedUser)
	res, err := json.Marshal(deletedUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(res)
}
//Car Handler

func indexCarHandler(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	setCors(w)
	var cars[] database.Car
	database.DB.Find(&cars)
	res,err := json.Marshal(cars)
	if err != nil{
		http.Error(w,err.Error(),500)
	}
	w.Write(res)
}
func showCarHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var car database.Car
	database.DB.Where("ID = ?", ps.ByName("carId")).First(&car)
	res, err := json.Marshal(car)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(res)
}
func createCarHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	setCors(w)
	decoder := json.NewDecoder(r.Body)
	var newCar database.Car

	if err:= decoder.Decode(&newCar);err != nil {
		http.Error(w,err.Error(),400)
		return
	}

	database.DB.Create(&newCar)
	res,err := json.Marshal(newCar)
	if err != nil {
		http.Error(w,err.Error(),500)
		return
	}
	w.Write(res)
}
func updateCarHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	type body struct {
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
	var updates body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updates); err != nil {
		http.Error(w, err.Error(), 400)
	}

	var updatedCar database.Car
	database.DB.Where("ID = ?", ps.ByName("carId")).First(&updatedCar)

	updatedCar.Immatriculation = updates.Immatriculation
	updatedCar.ModelName = updates.ModelName
	updatedCar.Year = updates.Year
	updatedCar.Weight = updates.Weight
	updatedCar.Power = updates.Power
	updatedCar.Length = updates.Length
	updatedCar.Width = updates.Width
	updatedCar.Height = updates.Height
	updatedCar.Picture = updates.Picture
	updatedCar.AgencyID = updates.AgencyID

	database.DB.Save(&updatedCar)
	res, err := json.Marshal(updatedCar)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(res)
}
func deleteCarHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var deletedCar database.Car
	database.DB.Where("ID = ?", ps.ByName("carId")).Delete(&deletedCar)
	res, err := json.Marshal(deletedCar)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Write(res)
}
// Status Handler
func indexStatusHandler(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	setCors(w)
	var status[] database.State
	database.DB.Find(&status)
	res,err := json.Marshal(status)
	if err != nil{
		http.Error(w,err.Error(),500)
	}
	w.Write(res)
}
func showStatusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var status database.State
	database.DB.Where("ID = ?", ps.ByName("statusId")).First(&status)
	res, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(res)
}
func createStatusHandler(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	setCors(w)
	decoder := json.NewDecoder(r.Body)
	var newStatus database.State

	if err:= decoder.Decode(&newStatus);err != nil {
		http.Error(w,err.Error(),400)
		return
	}

	database.DB.Create(&newStatus)
	res,err := json.Marshal(newStatus)
	if err != nil {
		http.Error(w,err.Error(),500)
		return
	}
	w.Write(res)
}
func updateStatusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	type body struct {
		Libelle string
	}
	var updates body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updates); err != nil {
		http.Error(w, err.Error(), 400)
	}

	var updatedStatus database.State
	database.DB.Where("ID = ?", ps.ByName("statusId")).First(&updatedStatus)
	updatedStatus.Libelle = updates.Libelle
	database.DB.Save(&updatedStatus)
	res, err := json.Marshal(updatedStatus)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(res)
}
func deleteStatusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var deletedStatus database.State
	database.DB.Where("ID = ?", ps.ByName("statusId")).Delete(&deletedStatus)
	res, err := json.Marshal(deletedStatus)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(res)
}
//CarState Handler

func indexCarStateHandler(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	setCors(w)
	var carStates[] database.CarState
	database.DB.Find(&carStates)
	res,err := json.Marshal(carStates)
	if err != nil{
		http.Error(w,err.Error(),500)
	}
	w.Write(res)
}
func showCarStateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var carState database.CarState
	database.DB.Where("ID = ?", ps.ByName("carStateId")).First(&carState)
	res, err := json.Marshal(carState)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(res)
}
func createCarStateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	setCors(w)
	decoder := json.NewDecoder(r.Body)
	var newCarState database.CarState

	if err:= decoder.Decode(&newCarState);err != nil {
		http.Error(w,err.Error(),400)
		return
	}

	database.DB.Create(&newCarState)
	res,err := json.Marshal(newCarState)
	if err != nil {
		http.Error(w,err.Error(),500)
		return
	}
	w.Write(res)
}
func updateCarStateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	type body struct {
		CarID 				uint
		StateID				uint
		UserID				uint
		DateReservation 	time.Time
		BeginDate			time.Time
		EndDate				time.Time
		}
	var updates body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updates); err != nil {
		http.Error(w, err.Error(), 400)
	}

	var updatedCarState database.CarState
	database.DB.Where("ID = ?", ps.ByName("carStateId")).First(&updatedCarState)

	updatedCarState.CarID = updates.CarID
	updatedCarState.StateID = updates.StateID
	updatedCarState.UserID = updates.UserID
	updatedCarState.DateReservation = updates.DateReservation
	updatedCarState.BeginDate = updates.BeginDate
	updatedCarState.EndDate = updates.EndDate

	database.DB.Save(&updatedCarState)
	res, err := json.Marshal(updatedCarState)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(res)
}
func deleteCarStateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setCors(w)
	var deletedCarState database.CarState
	database.DB.Where("ID = ?", ps.ByName("carStateId")).Delete(&deletedCarState)
	res, err := json.Marshal(deletedCarState)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Write(res)
}
//Home
func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
	fmt.Fprintf(w, "This is the RESTful api")
}
// used for COR preflight checks
func corsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
}

// util
func getFrontendUrl() string {
	if os.Getenv("APP_ENV") == "production" {
		return "http://localhost:3000" // change this to production domain
	} else {
		return "http://localhost:3000"
	}
}

func setCors(w http.ResponseWriter) {
	frontendUrl := getFrontendUrl()
	w.Header().Set("Access-Control-Allow-Origin", frontendUrl)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Temporary Canary test to make sure Travis-CI is working
func Canary(word string) string {
	return word
}

func main() {
	defer database.DB.Close()



	// add router and routes
	router := httprouter.New()
	//Home
	router.GET("/", indexHandler)
	//Agency routes
	router.GET("/agency",indexAgencyHandler)
	router.GET("/agency/:agencyId",showAgencyHandler)
	router.POST("/agency",createAgencyHandler)
	router.PUT("/agency/:agencyId",updateAgencyHandler)
	router.DELETE("/agency/:agencyId",deleteAgencyHandler)
	//User routes
	router.GET("/user",indexUserHandler)
	router.GET("/user/:userId",showUserHandler)
	router.POST("/user",createUserHandler)
	router.PUT("/user/:userId",updateUserHandler)
	router.DELETE("/user/:userId",deleteUserHandler)
	//Car routes
	router.GET("/car",indexCarHandler)
	router.GET("/car/:carId",showCarHandler)
	router.POST("/car",createCarHandler)
	router.PUT("/car/:carId",updateCarHandler)
	router.DELETE("/car/:carId",deleteCarHandler)
	//Status routes
	router.GET("/status",indexStatusHandler)
	router.GET("/status/:statusId",showStatusHandler)
	router.POST("/status",createStatusHandler)
	router.PUT("/status/:statusId",updateStatusHandler)
	router.DELETE("/status/:statusId",deleteStatusHandler)
	//CarState routes
	router.GET("/carstate",indexCarStateHandler)
	router.GET("/carstate/:carStateId",showCarStateHandler)
	router.POST("/carstate",createCarStateHandler)
	router.PUT("/carstate/:carStateId",updateCarStateHandler)
	router.DELETE("/carstate/:carStateId",deleteCarStateHandler)
	//OPTIONS
	router.OPTIONS("/*any", corsHandler)


	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	log.Println("connected to DB")

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}
