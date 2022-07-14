package health

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/learnProto/configs"
)

// Server is the server object for this api service.
type Server struct {
	configs *configs.Configs
	logger  *logrus.Logger
}

// New creates a new server.
func New(configs *configs.Configs, logger *logrus.Logger) *Server {
	return &Server{
		configs: configs,
		logger:  logger,
	}
}
