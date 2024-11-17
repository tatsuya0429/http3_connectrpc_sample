package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/quic-go/quic-go/http3"
)

const (
	url     = "https://127.0.0.1:8080/greet.v1.GreetService/Greet"
	reqBody = `{"name": "world"}`
)

func main() {
	roundTripper := &http3.Transport{
		TLSClientConfig: &tls.Config{
			// we need this because our certificate is self signed
			InsecureSkipVerify: true,
		},
	}
	defer roundTripper.Close()
	client := &http.Client{
		Transport: roundTripper,
	}

	log.Println("connect: ", url)
	log.Println("send: ", reqBody)
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer resp.Body.Close()

	log.Println("recv: ", string(body))
}
