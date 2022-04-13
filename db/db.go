package db

import (
	"database/sql"
	"fmt"
	"log"
	models "nalanda/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "go"
	password = "aezakmi"
	dbname   = "nalanda"
)

//global db pointer
var DBCon *sql.DB

func ConnectDB() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DBCon = db

	log.Println("database connected...")
}

func CreateRecord(s models.Student, a models.Address, c models.Contact, r models.Raw) {
	counter := 0 //for checking if all db operations were completed

	err := DBCon.Ping()
	if err != nil {
		panic((err))
	}
	insertStudent := `INSERT INTO student(enum,expDate,homeLib,category,fName,lName,dob,gender,contactNote,bookCategory,rollNo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ;`
	insertStudentStmt, err := DBCon.Prepare(insertStudent)
	if err != nil {
		log.Println(err)
	}

	insertAddress := `INSERT INTO address(enum,street,address,addressTwo,city,state,country,zip) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ;`
	insertAddressStmt, err := DBCon.Prepare(insertAddress)
	if err != nil {
		log.Println(err)
	}

	insertContact := `INSERT INTO contact(enum,primaryPhone,secondaryPhone,otherPhone,primaryEmail,secondaryEmail,fax) VALUES ($1, $2, $3, $4, $5, $6, $7) ;`
	insertContactStmt, err := DBCon.Prepare(insertContact)
	if err != nil {
		log.Println(err)
	}

	insertRaw := `INSERT INTO raw(enum,profile,checkout) VALUES ($1, $2, $3) ;`
	insertRawStmt, err := DBCon.Prepare(insertRaw)
	if err != nil {
		log.Println(err)
	}

	_, err = insertStudentStmt.Exec(s.Enum, s.ExpDate, s.HomeLib, s.Category, s.FName, s.LName, s.Dob, s.Gender, s.ContactNote, s.BookCategory, s.RollNo)
	if err != nil {
		log.Println(err)
	} else {
		counter++
	}

	_, err = insertAddressStmt.Exec(a.Enum, a.Street, a.Address, a.AddressTwo, a.City, a.State, a.Country, a.Zip)
	if err != nil {
		log.Println(err)
	} else {
		counter++
	}

	_, err = insertContactStmt.Exec(c.Enum, c.PrimaryPhone, c.SecondaryPhone, c.OtherPhone, c.PrimaryEmail, c.SecondaryEmail, c.Fax)
	if err != nil {
		log.Println(err)
	} else {
		counter++
	}

	_, err = insertRawStmt.Exec(r.Enum, r.Profile, r.Checkout)
	if err != nil {
		log.Println(err)
	} else {
		counter++
	}

	if counter == 4 {
		UpdatePermutation(2, s.Enum) //mark this enum as completed
		log.Println(s.Enum, " complete")
	}

}

func UpdatePermutation(status int, enum int) { // status 1 for hit 2 for complete
	update := `UPDATE permutations SET hit=0 WHERE enum=$1`
	if status == 1 {
		update = `UPDATE permutations SET hit=1 WHERE enum=$1`
	}
	if status == 2 {
		update = `UPDATE permutations SET complete=1 WHERE enum=$1`
	}
	updateStmt, err := DBCon.Prepare(update)
	if err != nil {
		log.Println(err)
	}
	_, err = updateStmt.Exec(enum)
	if err != nil {
		log.Println(err)
	}
}

//generate a slice of all the enums yet to be touched
func JobsGen() []int {
	enumFetch := "SELECT enum FROM permutations WHERE hit IS null"
	rows, err := DBCon.Query(enumFetch)
	var res []int
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var e int
		err := rows.Scan(&e)
		if err != nil {
			log.Println(err)
		}
		res = append(res, e)
	}
	return res
}
