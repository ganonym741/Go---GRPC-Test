package router

import (
	"fmt"
	"log"
	"net"

	"github.com/sirupsen/logrus"

	"gitlab.com/go-grpc-learning/api/v1/health"
	"gitlab.com/go-grpc-learning/configs"
	hlpb "gitlab.com/go-grpc-learning/proto/v1/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewGRPCServer creates the grpc server serve mux.
func NewGRPCServer(configs *configs.Configs, logger *logrus.Logger) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configs.Config.Server.Grpc.Port))
	if err != nil {
		return err
	}

	// register grpc service server
	grpcServer := grpc.NewServer()

	hlpb.RegisterHealthServiceServer(grpcServer, health.New(configs, logger))

	// add reflection service
	reflection.Register(grpcServer)

	// running gRPC server
	log.Println("[SERVER] GRPC server is ready")

	grpcServer.Serve(lis)

	return nil
}
