package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
)

//DB ...
type DB struct {
	*sql.DB
}

const (
	//DbUser ...
	DbUser = "postgres"

	//DbPassword ...
	DbPassword = "example"

	//DbName ...
	DbName = "team_manage_app"
)

var db *gorp.DbMap

//Init ...
func Init() {
	dbinfo := os.Getenv("DATABASE_URL")
	if dbinfo == "" {
		dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			DbUser, DbPassword, DbName)
	}
	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
