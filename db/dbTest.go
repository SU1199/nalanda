package db

import (
	models "nalanda/models"
)

func test() {
	ConnectDB()
	s := models.Student{
		Enum:         102199003,
		ExpDate:      "ED",
		HomeLib:      "HL",
		Category:     "C",
		FName:        "FN",
		LName:        "LN",
		Dob:          "D",
		Gender:       "G",
		ContactNote:  "CN",
		BookCategory: "BC",
		RollNo:       "RN",
	}
	a := models.Address{
		Enum:       102199003,
		Street:     "S",
		Address:    "A",
		AddressTwo: "AT",
		City:       "C",
		State:      "S",
		Country:    "C",
		Zip:        "Z",
	}
	c := models.Contact{
		Enum:           102199003,
		PrimaryPhone:   "PP",
		SecondaryPhone: "SP",
		OtherPhone:     "OP",
		PrimaryEmail:   "PE",
		SecondaryEmail: "SE",
		Fax:            "F",
	}
	r := models.Raw{
		Enum:     102199003,
		Profile:  "P",
		Checkout: "C",
	}

	CreateRecord(s, a, c, r)
}
