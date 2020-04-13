package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   string
	port     int
	user     string
	password string
	database string

	db *sql.DB
)

func main() {

	flag.StringVar(&server, "server", "localhost", "the server address")
	flag.IntVar(&port, "port", 1433, "the port for connetihg db")
	flag.StringVar(&user, "user", "sa", "DB user name")
	flag.StringVar(&password, "password", "enoviaV6", "DB Password")
	flag.StringVar(&database, "database", "spacedb", "Database Name")

	flag.Parse()

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Printf("Error while connecting db : %v", err)
	}

	defer db.Close()

	SelectVersion()

	SelectData()
}

// SelectData Mtheod to return the lxTable of DB
func SelectData() {
	ctx := context.Background()

	err := db.PingContext(ctx)

	if err != nil {
		fmt.Printf("Unable to Ping db %v", err)
	}

	var fileName = "20079659.pdf"
	lxRows, queryErr := db.QueryContext(ctx, "SELECT lxPath FROM spacedb.lxFile_eb44cf05 where lxPath='?'", fileName)
	if queryErr != nil {
		fmt.Printf("Error while executing query %v", queryErr)
		return
	}
	defer lxRows.Close()
	names := make([]string, 0)

	for lxRows.Next() {
		var name string
		if err != lxRows.Scan(&name) {
			log.Fatal("Error while reading row:", err)
			return
		}
		names = append(names, name)
	}
	fmt.Println(names)

}

// SelectVersion Mtheod to return the version of DB
func SelectVersion() {

	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		fmt.Printf("Unable to Ping db %v", err)
		return
	}
	query := "SELECT @@VERSION"
	var result string
	var queryErr = db.QueryRowContext(ctx, query).Scan(&result)
	if queryErr != nil {
		fmt.Printf("Error while executing query %v", err)
		return
	}

	fmt.Printf("Output:| %s\n", result)
}
