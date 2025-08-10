package main

import (
	"log"
)

func main() {
	config, err := ParseConfig()

	if err != nil {
		log.Fatal(err)
		return
	}

	ip, err := GetCurrentIp()

	if err != nil {
		log.Fatal(err)
		return
	}

	err = UpdateRecords(config, ip)

	if err != nil {
		log.Fatal(err)
	}
}
