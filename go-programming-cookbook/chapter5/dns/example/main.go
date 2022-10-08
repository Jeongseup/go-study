package main

import (
	"dns"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <address>\n", os.Args[0])
		os.Exit(1)
	}

	address := os.Args[1]
	lookup, err := dns.LookupAddress(address)
	if err != nil {
		log.Panicf("failed to lookup: %s\n", err.Error())
	}
	fmt.Println(lookup)
	fmt.Println(lookup.String())
}
