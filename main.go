package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Will override these

	userPtr := flag.String("username", "", "username as string")
	passPtr := flag.String("password", "", "password as string")
	methodPtr := flag.String("method", "GET", "method as string")
	urlPtr := flag.String("url", "http://google.com", "url as string")
	headerKeyPtr := flag.String("headerKey", "", "headerKey as string")
	headerValPtr := flag.String("headerVal", "", "headerVal as string")
	payloadPtr := flag.String("p", "", "payload p as string")

	payload := []byte(*payloadPtr)

	req := setupRequest(payload, *methodPtr, *urlPtr)

	if *userPtr != "" && *passPtr != "" {
		req.SetBasicAuth(*userPtr, *passPtr)
	}

	if *headerKeyPtr != "" && *headerValPtr != "" {
		req.Header.Set(*headerKeyPtr, *headerValPtr)
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal()
	}
	s := string(bodyText)

	if len(s) > 0 {
		fmt.Println("Worked")
	} else {
		fmt.Println("NOOP")
	}

	//This will need to take a semi curl input and execute the
	//user/pass and method and header
	//Build scratch container then pass in dockercompose as command
	//Check equi dockerfiles and compose

}

func setupRequest(payload []byte, method, url string) *http.Request {
	if len(payload) > 0 {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
		if err != nil {
			log.Fatal()
		}
		return req
	} else {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			log.Fatal()
		}
		return req
	}
}
