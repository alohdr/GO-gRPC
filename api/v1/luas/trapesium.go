package luas

import (
	"context"

	hlpb "gitlab.com/learnProto/proto/v1/luas"
)

func (s *Server) CalculateTrapesium(ctx context.Context, req *hlpb.Trapesium) (*hlpb.Response, error) {
	s.logger.Infof("[LUAS][POST] SUCCESS")

	tempResult := 0.5 * float64(req.Height) * float64(req.Base1+req.Base2)
	return &hlpb.Response{
		Result: int64(tempResult),
	}, nil
}
