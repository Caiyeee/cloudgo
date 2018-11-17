package main

import (
	"os"
	flag "github.com/spf13/pflag"
	"github.com/Caiyeee/cloudgo/service"
)

const(
	PORT string = "8080"
)
func main() {
	var port string
	port = os.Getenv("PORT")
	if len(port) <= 0 {
		port = PORT
	}

	flag.StringVarP(&port, "port", "p", "8080", "the port to listen")
	flag.Parse()

	server := service.NewServer()
	server.Run(":" + port)
}