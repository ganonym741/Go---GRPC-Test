package health

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.com/go-grpc-learning/pkg/v1/utils/constants"
	hlpb "gitlab.com/go-grpc-learning/proto/v1/health"
)

// ProcessController acts as the main entry point for this get service
func (s *Server) Get(ctx context.Context, req *empty.Empty) (*hlpb.Response, error) {
	s.logger.Infof("[HEALTH][GET] SUCCESS")

	return &hlpb.Response{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
	}, nil
}
