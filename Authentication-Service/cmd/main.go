package main

import (
	pb "Github.com/LocalEats/Authentication-Service/gen-proto/auth"
	"Github.com/LocalEats/Authentication-Service/internal/configs"
	"Github.com/LocalEats/Authentication-Service/internal/repositoty"
	"Github.com/LocalEats/Authentication-Service/internal/service"
	"Github.com/LocalEats/Authentication-Service/internal/storage"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config := configs.Load()

	db, err := storage.ConnectDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("Listening on port 8080")

	authStr := repositoty.NewAuthRepository(db)

	as := service.NewAuthService(*authStr)

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, as)

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
