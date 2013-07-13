package main

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
)

type Config struct {
	Username string
	Password string
}

var c Config

func main() {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	if err := goyaml.Unmarshal([]byte(file), &c); err != nil {
		panic(err)
	}

	fmt.Printf("--- t:\n%v\n\n", c)

	d, err := goyaml.Marshal(&c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

}
