package requests

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Body       string
	StatusCode int
	Headers    map[string]string
}

func Get(url string, data []byte, headers map[string]string) (body string, statusCode int, response_headers http.Header, err error) {
	return processRequest("GET", url, data, headers)
}

func Post(url string, data []byte, headers map[string]string) (body string, statusCode int, response_headers http.Header, err error) {
	return processRequest("POST", url, data, headers)
}

func Put(url string, data []byte, headers map[string]string) (body string, statusCode int, response_headers http.Header, err error) {
	return processRequest("PUT", url, data, headers)
}

func Delete(url string, data []byte, headers map[string]string) (body string, statusCode int, response_headers http.Header, err error) {
	return processRequest("DELETE", url, data, headers)
}

func processRequest(method string, url string, data []byte, headers map[string]string) (body string, statusCode int, response_headers http.Header, err error) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return body, statusCode, response_headers, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	r, err := http.DefaultClient.Do(req)
	defer r.Body.Close()

	if err != nil {
		log.Printf("Response Error: %v | Response Object:  %+v", err, r)
		return body, statusCode, response_headers, err
	}

	response_body, _ := ioutil.ReadAll(r.Body)
	return string(response_body), r.StatusCode, r.Header, err
}
