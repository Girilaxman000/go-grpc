package main

import (
	"log"
	"net"

	pb "github.com/Girilaxman000/go-grpc/proto"
	"google.golang.org/grpc"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main() {
	list, err := net.Listen("tcp", ":8000")
	if err != nil{
		log.Fatalf("Failed to start the server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", list.Addr())
	if err := grpcServer.Serve(list); err != nil{
		log.Fatalf("Failed to start")
	}
}