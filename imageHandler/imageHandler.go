package imagehandler

import (
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

//fetches image converts it into base64 and returns the data url
func ConvImage(cookie string, enum int) string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("GET", "https://library.thapar.edu/cgi-bin/koha/opac-patron-image.pl", nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Cookie", "CGISESSID="+cookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	b64 := "data:image/png;base64,"
	b64 += toBase64(bytes)

	return b64
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
