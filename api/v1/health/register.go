package health

import (
	"context"

	"gitlab.com/learnProto/pkg/v1/postgres"
	hlpb "gitlab.com/learnProto/proto/v1/health"
)

func (s *Server) Register(ctx context.Context, req *hlpb.Registrasi) (*hlpb.Data, error) {
	s.logger.Infof("[HEALTH][Registrasi] SUCCESS")

	rows, err := s.configs.Pg.CustomMainRegister(&postgres.Registrasi{
		UserId: req.GetUserId(),
		Pass:   req.GetPass(),
	})

	if err != nil {
		s.logger.Errorf("[HEALTH][REGISTRASI] ERROR %v", err)
		return nil, err
	}

	return &hlpb.Data{
		UserId: rows.Pass,
		Pass:   rows.UserId,
	}, nil
}

func (s *Server) DeleteRegister(ctx context.Context, req *hlpb.Id) (*hlpb.Id, error) {
	s.logger.Infof("[HEALTH][Delete Register] SUCCESS")

	rows, err := s.configs.Pg.CustomMainDeleteRegister(&postgres.Id{
		UserId: req.GetUserId(),
	})

	if err != nil {
		s.logger.Errorf("[HEALTH][DELETE REGISTRASI] ERROR %v", err)
		return nil, err
	}

	return &hlpb.Id{
		UserId: rows.UserId,
	}, nil
}
