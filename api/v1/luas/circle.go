package luas

import (
	"context"
	"math"

	hlpb "gitlab.com/learnProto/proto/v1/luas"
)

func (s *Server) CalculateCircle(ctx context.Context, req *hlpb.Circle) (*hlpb.Response, error) {
	s.logger.Infof("[LUAS][POST] SUCCESS")

	tempResult := math.Pi * math.Pow(float64(req.Jari), 2)
	return &hlpb.Response{
		Result: int64(tempResult),
	}, nil
}
