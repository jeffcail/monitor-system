package request

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	def := http.DefaultTransport
	defPot, ok := def.(*http.Transport)
	if !ok {
		panic("init transport出错")
	}
	defPot.MaxIdleConns = 100
	defPot.MaxIdleConnsPerHost = 100
	defPot.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client = &http.Client{
		Timeout:   time.Second * time.Duration(20),
		Transport: defPot,
	}
}

// Get
func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return nil, err
	}
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return bb, err
}
