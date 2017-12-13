package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	//"github.com/weirdsnap/agendaweb/service"
	"github.com/weirdsnap/upgraded-agenda/service/service"
	"os"
)

const (
	PORT string = "8080"
)

func main() {
	//use environment variables if exist
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
		fmt.Println("prot change to ", port)
	}
	//use flap "p" to set port if p's value exist
	pPort := flag.StringP("port", "p", port, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}
