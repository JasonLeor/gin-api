package main

import (
	"encoding/json"
	"log"
	"os"
)

func init() {
	var data interface{}
	fr, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	decoder := json.NewDecoder(fr)
	err = decoder.Decode(&data)
}
