package health

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"gitlab.com/learnProto/pkg/v1/utils/constants"
	hlpb "gitlab.com/learnProto/proto/v1/health"
)

var data = make(map[string]hlpb.ReqPatient)

// ProcessController acts as the main entry point for this get service
func (s *Server) Get(ctx context.Context, req *empty.Empty) (*hlpb.Response, error) {
	s.logger.Infof("[HEALTH][GET] SUCCESS")

	return &hlpb.Response{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
	}, nil
}

func (s *Server) GetPatient(ctx context.Context, req *empty.Empty) (*hlpb.ResPatient, error) {
	s.logger.Infof("[PATIENT][GET] SUCCESS")

	return &hlpb.ResPatient{
		Id:   "1",
		Name: "Hendro",
	}, nil
}

func (s *Server) CreatePatient(ctx context.Context, req *hlpb.ReqPatient) (*hlpb.ResPatient, error) {
	s.logger.Infof("[PATIENT][CREATE] SUCCESS")

	id := uuid.New().String()

	data[id] = hlpb.ReqPatient{
		Id:   req.Id,
		Name: req.Name,
	}

	return &hlpb.ResPatient{
		Id: id,
	}, nil
}

func (s *Server) GetAllPatients(ctx context.Context, req *empty.Empty) (*hlpb.ResAllPatients, error) {
	s.logger.Infof("[ALL PATIENT][GET] SUCCESS")

	resp := hlpb.ResAllPatients{}

	for key, v := range data {
		resp.AllPatient = append(resp.AllPatient, &hlpb.ResPatient{
			Id:   key,
			Name: v.Name,
		})
	}

	return &resp, nil

}
