package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/Rodrigo001-de/code-pix/application/grpc/pb"
	"github.com/Rodrigo001-de/code-pix/application/usecase"
	"github.com/Rodrigo001-de/code-pix/infrastructure/repositories"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	// essa linha vai debugar usando um client de grpc
	reflection.Register(grpcServer)

	pixRepository := repositories.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

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
