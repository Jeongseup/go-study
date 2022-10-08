package main

import "global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
