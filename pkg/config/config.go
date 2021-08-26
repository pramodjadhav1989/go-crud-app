package config

import("log"
"database/sql"
_"github.com/go-sql-driver/mysql"
)
var DB *sql.DB
func Connect(){

db,err :=sql.Open("mysql","root:pass@123@tcp(127.0.0.1:3306)/testDb")
defer db.Close()
if err != nil {
	log.Fatal(err)
	
}
DB=db

}

func GetDB() *sql.DB{
	return DB;
}