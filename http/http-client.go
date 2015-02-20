// Override transport for weird but possible http client scenarios

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(5 * time.Second)}
	resp, _ := client.Get("http://path.to.wav")
	fmt.Println(resp.Header["Cache-Control"][0])

	for k, v := range strings.Split(resp.Header["Cache-Control"][0], ",") {
		fmt.Println(k, v)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	clength := fmt.Sprintf("%d", len(data))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "audio/wav")
		w.Header().Set("Content-Length", clength)
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("User-Agent", "PlivoGoTTS/0.1")
		w.Write(data)
	})
	http.ListenAndServe(":8000", nil)

}
