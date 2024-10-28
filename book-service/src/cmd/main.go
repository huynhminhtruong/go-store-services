package main

import (
	"log"

	"github.com/huynhminhtruong/go-store-services/book-service/src/src/config"
	"github.com/huynhminhtruong/go-store-services/book-service/src/src/internal/adapters/db"
	"github.com/huynhminhtruong/go-store-services/book-service/src/src/internal/adapters/grpc"
	"github.com/huynhminhtruong/go-store-services/book-service/src/src/internal/application/core/api"
)

/*
	Hexagonal architecture(Ports and Adapters):

		1. primary adapter(chủ động điều khiển các luồng xử lý của hệ thống từ bên ngoài vào):
			storing application

		2. second adapter(chủ động gửi request ra ngoài để gọi các dịch vụ bên ngoài):
			database
			shipping service
*/

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	// init and connect to inside business logic
	application := api.NewApplication(dbAdapter)
	// init storing adapter which will receive requests from outside
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
