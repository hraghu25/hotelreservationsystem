package main

import (
	"flag"

	"github.com/hraghu25/hotelreservationsystem/server"
)

func main() {
	ListenAddress := flag.String("listenAddre", ":8080", "The listen address of API server")

	server.Execute(ListenAddress)
}
