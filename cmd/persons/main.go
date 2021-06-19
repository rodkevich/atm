package main

import (
	"fmt"
	"github.com/rodkevich/atm/persons"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &persons.GRPCServer{}
	persons.RegisterPersonsServer(s, srv)

	l, err := net.Listen("tcp", ":9090")
	fmt.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
