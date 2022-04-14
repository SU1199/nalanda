package parser

import (
	"log"
	"nalanda/models"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Parser(resp *http.Response, enum int) (models.Student, models.Address, models.Contact, models.Raw) {
	s := new(models.Student)
	a := new(models.Address)
	c := new(models.Contact)
	r := new(models.Raw)

	s.Enum = enum
	a.Enum = enum
	c.Enum = enum
	r.Enum = enum

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
	}

	//lname
	doc.Find("#borrower_surname").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			s.LName = val
		}
	})

	//fname
	doc.Find("#borrower_firstname").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			s.FName = val
		}
	})

	//dob
	doc.Find("#borrower_dateofbirth").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			s.Dob = val
		}
	})

	//street
	doc.Find("#borrower_streetnumber").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			a.Street = val
		}
	})

	//address
	doc.Find("#borrower_address").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			a.Address = val
		}
	})

	//address two
	doc.Find("#borrower_address2").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			a.AddressTwo = val
		}
	})

	//city
	doc.Find("#borrower_city").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			a.City = val
		}
	})

	//state
	doc.Find("#borrower_state").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			a.State = val
		}
	})

	//zip
	doc.Find("#borrower_zipcode").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			a.Zip = val
		}
	})

	//country
	doc.Find("#borrower_country").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			a.Country = val
		}
	})

	//phone
	doc.Find("#borrower_phone").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			c.PrimaryPhone = val
		}
	})

	//second_phone
	doc.Find("#borrower_phonepro").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			c.SecondaryPhone = val
		}
	})

	//other phone
	doc.Find("#borrower_mobile").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			c.OtherPhone = val
		}
	})

	//email
	doc.Find("#borrower_email").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			c.PrimaryEmail = val
		}
	})

	//second email
	doc.Find("#borrower_emailpro").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			c.SecondaryEmail = val
		}
	})

	//fax
	doc.Find("#borrower_fax").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			c.Fax = val
		}
	})

	//contact note
	doc.Find("#borrower_contactnote").Each(func(i int, sel *goquery.Selection) {
		val, err := sel.Attr("value")
		if err {
			s.ContactNote = val
		}
	})

	//exp date
	strippedRes := stripchars(doc.Find("#memberentry_library > ol > li:nth-child(2)").Text(), " \n")
	if len(strippedRes) > 15 {
		strippedRes = strippedRes[15:]
		s.ExpDate = strippedRes
	}

	//category
	strippedRes = stripchars(doc.Find("#memberentry_library > ol > li:nth-child(4)").Text(), " \n")
	if len(strippedRes) > 9 {
		strippedRes = strippedRes[9:]
		s.Category = strippedRes
	}

	//home lib
	doc.Find("#borrower_branchcode").Children().Each(func(i int, sel *goquery.Selection) {
		_, err := sel.Attr("selected")
		if err {
			s.HomeLib = sel.Text()
		}
	})

	//gender
	doc.Find("#memberentry_identity > ol > li.lradio").Children().Each(func(i int, sel *goquery.Selection) {
		_, err := sel.Attr("checked")
		if err {
			val, _ := sel.Attr("value")
			s.Gender = val
		}
	})

	//book category
	strippedRes = stripchars(doc.Find("#memberentry-form > div:nth-child(7) > div > fieldset > ol > li:nth-child(2)").Text(), " \n")[19:]
	if len(strippedRes) > 19 {
		strippedRes = strippedRes[19:]
		s.BookCategory = strippedRes
	}

	//roll number
	strippedRes = stripchars(doc.Find("#memberentry-form > div:nth-child(7) > div > fieldset > ol > li:nth-child(4)").Text(), " \n")
	if len(strippedRes) > 19 {
		strippedRes = strippedRes[19:]
		s.RollNo = strippedRes
	}
	return *s, *a, *c, *r
}

//courtesy -> http://rosettacode.org/wiki/Strip_a_set_of_characters_from_a_string#Go
func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
