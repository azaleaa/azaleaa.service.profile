package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type KeyValPair struct {
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Profile struct {
	Email string `json:"email,omitempty"`
	Profilevals []KeyValPair `json:"profilevals,omitempty"`
}

type Configuration struct {
	AZALEAA_DB_DSN string
}


func CreateProfileEndpoint(w http.ResponseWriter, req *http.Request) {
//	params := mux.Vars(req)
	
//	fmt.Println(req)

	file, _ := os.Open("azaleaa.conf")
	decoder1 := json.NewDecoder(file)
	configuration := Configuration{}
	err1 := decoder1.Decode(&configuration)

	checkErr(err1)

	var profile Profile 
	decoder :=  json.NewDecoder(req.Body)

	err := decoder.Decode(&profile)
	
	fmt.Println(profile)

	db_dsn := configuration.AZALEAA_DB_DSN
	
	fmt.Printf("testvar value is %v", db_dsn)
	db, err := sql.Open("mysql", db_dsn)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO t_userprofile VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE value=?")
	checkErr(err)

	for _,v := range profile.Profilevals {
		//fmt.Println(i)
		//fmt.Println(v)
		res, err := stmt.Exec(profile.Email, v.Key, v.Value, v.Value)
		checkErr(err)
	
		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)
	}

	db.Close()
	return
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/profiles", CreateProfileEndpoint).Methods("POST")
	fmt.Printf("Starting server")
	log.Fatal(http.ListenAndServe(":80", router))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
