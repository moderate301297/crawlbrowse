package dbconnection

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

var (
	Connect *sql.DB
)

func init() {
	var err error
	Connect, err = sql.Open("mysql", "hihi:beeketing@tcp(localhost:3306)/data_walmart")
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func Close() {
	Connect.Close()
}
