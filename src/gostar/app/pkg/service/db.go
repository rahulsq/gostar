package service

import (
    "fmt"
    // "log"
    // "net/http"
    // "os"

    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/joho/godotenv"
)

var db *gorm.DB

const (
    dbhost = "DBHOST"
    dbport = "DBPORT"
    dbuser = "DBUSER"
    dbpass = "DBPASS"
    dbname = "DBNAME"
)

func init() {
    config := dbConfig()
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        config[dbhost], config[dbport],
        config[dbuser], config[dbpass], config[dbname])

    db, err = gorm.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	host := "localhost"
	port := "5432"
	user := "odoo"
	name := "star_odoo"
	password := "test123"
    conf := make(map[string]string)
    // host, ok := os.LookupEnv(dbhost)
    // if !ok {
    //     panic("DBHOST environment variable required but not set")
    // }
    // port, ok := os.LookupEnv(dbport)
    // if !ok {
    //     panic("DBPORT environment variable required but not set")
    // }
    // user, ok := os.LookupEnv(dbuser)
    // if !ok {
    //     panic("DBUSER environment variable required but not set")
    // }
    // password, ok := os.LookupEnv(dbpass)
    // if !ok {
    //     panic("DBPASS environment variable required but not set")
    // }
    // name, ok := os.LookupEnv(dbname)
    // if !ok {
    //     panic("DBNAME environment variable required but not set")
    // }
    conf[dbhost] = host
    conf[dbport] = port
    conf[dbuser] = user
    conf[dbpass] = password
    conf[dbname] = name
    return conf
}

func GetDBConn() *gorm.DB {
    return db
}

func Close() {
    db.Close()
}