package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/huynhminhtruong/go-store-services/app-services/storing/config"
	"github.com/huynhminhtruong/go-store-services/app-services/storing/internal/ports"
	"github.com/huynhminhtruong/go-store-services/app-services/storing/services/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	book.UnimplementedBookServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	// Đăng ký service BookServer vào grpcServer
	// a (biến Adapter) là struct mà bạn đã implement để xử lý các phương thức của BookServer
	// Điều này nói với gRPC server rằng grpcServer
	// sẽ thực thi các service gRPC của BookServer bằng cách
	// sử dụng các hàm đã được triển khai trong Adapter
	book.RegisterBookServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}
