package main

import (
	"os"

	"github.com/Rodrigo001-de/code-pix/application/grpc"
	"github.com/Rodrigo001-de/code-pix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

// a função main é o ponto de entrada de executar o meu programa em go
func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
