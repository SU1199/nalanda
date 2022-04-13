package permutationGen

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "go"
	password = "aezakmi"
	dbname   = "nalanda"
)

func Generate(year int) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	makyr(db, year)

}

func makyr(db *sql.DB, year int) {
	for i := 1; i <= 99999; i++ {
		con := "10" + strconv.Itoa(year) + fmt.Sprintf("%05d", i)
		conInt, _ := strconv.Atoi(con)
		insertVal(db, conInt, year)
	}
}

// func createTable(db *sql.DB) {
// 	createTable := `CREATE TABLE permutations("id" integer not null primary key autoincrement, "enum" string not null, "year" string not null,  "ts" text, "hit" integer, "complete" integer);`

// 	statement, err := db.Prepare(createTable)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	statement.Exec()
// }

func insertVal(db *sql.DB, id int, year int) bool {
	insertEnum := `INSERT INTO permutations(enum,year) VALUES ($1, $2) ;`
	statement, err := db.Prepare(insertEnum)
	if err != nil {
		log.Println("A")
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, year)
	if err != nil {
		log.Println("B")
		log.Fatalln(err.Error())
		return false
	} else {
		return true
	}
}
