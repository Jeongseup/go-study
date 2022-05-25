package main

import "github.com/aofei/air"

func main() {
	air.Default.GET("/", func(req *air.Request, res *air.Response) error {
		return res.WriteString("헬로우 월드2")
	})

	air.Default.Serve()
}
