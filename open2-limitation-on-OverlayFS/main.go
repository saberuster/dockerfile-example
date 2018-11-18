package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./file", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.OpenFile("./file", os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(f2, "hello world!")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("data: %s", string(data))
}
