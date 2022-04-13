package models

type Student struct {
	Enum         int
	ExpDate      string
	HomeLib      string
	Category     string
	FName        string
	LName        string
	Dob          string
	Gender       string
	ContactNote  string
	BookCategory string
	RollNo       string
}

type Address struct {
	Enum       int
	Street     string
	Address    string
	AddressTwo string
	City       string
	State      string
	Country    string
	Zip        string
}

type Contact struct {
	Enum           int
	PrimaryPhone   string
	SecondaryPhone string
	OtherPhone     string
	PrimaryEmail   string
	SecondaryEmail string
	Fax            string
}

type Raw struct {
	Enum     int
	Profile  string
	Checkout string
}
