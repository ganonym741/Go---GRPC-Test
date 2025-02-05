package health

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/go-grpc-learning/configs"
)

// Server is the server object for this api service.
type Server struct {
	config *configs.Configs
	logger *logrus.Logger
}

// New creates a new server.
func New(config *configs.Configs, logger *logrus.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}
