package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc/tweak"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("error dialing", err)
	}

	args := tweak.Args{
		String:  "this string should be blah blah",
		ToUpper: true,
		Reverse: true,
	}

	var result string
	err = client.Call("StringTweaker.Tweak", args, &result)
	if err != nil {
		log.Fatal("client call with err: ", err)
	}
	fmt.Printf("the result is: %s", result)
}
