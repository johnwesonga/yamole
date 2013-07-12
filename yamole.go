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

func main() {
	var c Config
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = goyaml.Unmarshal([]byte(data), &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- t:\n%v\n\n", c)

	d, err := goyaml.Marshal(&c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

}
