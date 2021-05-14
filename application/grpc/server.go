package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	address := fmt.Sprintf("0.0.0.0:%d", port)
	// o listener vai ficar escutando a conexão
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("Grpc server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
