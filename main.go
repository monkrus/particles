

package main

import (

  "encoding/json"

  "log"

  "net/http"



  "github.com/gorilla/mux"

  "github.com/jinzhu/gorm"

  "github.com/rs/cors"



  _ "github.com/jinzhu/gorm/dialects/postgres"

)

type Cell struct {

	gorm.Model

	OneRock string
	TwoRocks string
	TRock string
	EmptySpace string
    Rows []Row
  
  }

var db *gorm.DB

var err error

var (
rows = []Row {

	{OneRock : '.', EmptySpace: ' ', TRock : 'T' , TwoRocks :  ':' , TRock : 'T' },
	{OneRock : '.', EmptySpace: ' ', TRock : 'T' , TwoRocks :  ':' , TRock : 'T' },
	{EmptySpace: ' ', TRock : 'T' , TwoRocks :  ':' , TRock : 'T' , OneRock : '.'},
	{OneRock : '.', TRock : 'T' , TwoRocks :  ':' , EmptySpace: ' ', TRock : 'T' },
	{TwoRocks :  ':', OneRock : '.', EmptySpace: ' ', TRock : 'T' ,  TRock : 'T' },

}

func main() {

	router := mux.NewRouter()

	db, err = gorm.Open( "postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")

	if err != nil {

		panic("failed to connect database")
	
	  }


	  defer db.Close()
	  
	  db.AutoMigrate(&Row{})

	  for index := range rows {

		db.Create(&rows[index])
  
	}

	router.HandleFunc("/rows", GetRows).Methods("GET")

	handler := cors.Default().Handler(router)

    log.Fatal(http.ListenAndServe(":8080", handler))


	func GetCRows(w http.ResponseWriter, r *http.Request) {

		var rows []Row
	  
		db.Find(&rows)
	  
		json.NewEncoder(w).Encode(&rows)
	  
	  }