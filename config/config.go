package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
)

func main() {

	type studentObj struct {
		Name   string
		Number int
	}
	type configObject struct {
		Student studentObj
	}

	blob, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Fatal("error reading file")
	}

	var con configObject
	data := string(blob)
	_, err = toml.Decode(data, &con)

	if err != nil {
		log.Fatal("decode failed")
	}
	fmt.Println(con.Student.Name)
}
