package service

import (
	"context"
	"github.com/bufbuild/protovalidate-go"

	pb "demo/api/acme/demo/v1"
	"demo/internal/biz"
)

type DemoService struct {
	pb.UnimplementedDemoServiceServer

	uc *biz.DemoUsecase
}

func NewDemoService(uc *biz.DemoUsecase) *DemoService {
	return &DemoService{uc: uc}
}

func (s *DemoService) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	result := &pb.CreateResponse{
		Code:    200,
		Message: "create demo success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	return result, nil
}

func (s *DemoService) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	result := &pb.GetResponse{
		Code:    200,
		Message: "get demo success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	return result, nil
}

func (s *DemoService) GetList(ctx context.Context, request *pb.GetListRequest) (*pb.GetListResponse, error) {
	result := &pb.GetListResponse{
		Code:    200,
		Message: "get demo list success",
	}

	return result, nil
}

func (s *DemoService) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	result := &pb.UpdateResponse{}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	return result, nil
}

func (s *DemoService) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	result := &pb.DeleteResponse{}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	return result, nil
}
