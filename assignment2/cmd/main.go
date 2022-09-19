package main

import (
	"gin-clean-architecture/httpserver"
)

func main() {
	app := httpserver.CreateRouter()
	app.Run(":1234")
}
