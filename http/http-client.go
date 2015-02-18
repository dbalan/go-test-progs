// Override transport for weird but possible http client scenarios

package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(5 * time.Second)}
	resp, _ := client.Get("https://server.with.invalid.cert")
	fmt.Println(resp)
}
