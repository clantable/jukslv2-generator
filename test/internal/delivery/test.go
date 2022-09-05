package delivery

import (
	"context"

	"github.com/clantable/juksl-v2/services/lists//internal/delivery/converter"
	"github.com/clantable/juksl-v2/services/lists//internal/usecase"
	"github.com/clantable/juksl-v2/services/lists/pb"
)


func (s *ListsServer) StoreTest(ctx context.Context, in *pb.TestParams) (*pb.TestParams, error) {
	data := converter.TestConverterToDomain(in)
	if err := data.Validate(); err != nil {
		return nil, err
	}
	newData, err := cs.uc.Store(data)
	if err != nil {
		return nil, err
	}
	out := converter.TestConverterToRPC(newData)
	return out, nil
}

func (s *ListsServer) FetchTest(ctx context.Context, in *pb.FetchConfig) (*pb.FetchTestResponse, error) {
	query := converter.ConvertFetchQueryToDomain(in)
	data, _, err := s.svcCtx.Test.Fetch(query)
	if err != nil {
		return nil, err
	}
	res := converter.TestConverterOfFetchResultToRPC(data)
	return res, nil
}

func (s *ListsServer) GetByIDTest(ctx context.Context, in *pb.GetByIDParams) (*pb.TestParams, error) {
	test, err := s.svcCtx.Test.GetByID(uint(in.GetId()), in.GetPreload()...)
	if err != nil {
		return nil, err
	}
	out := converter.TestConverterToDomain(test)
	return out, nil
}

func (s *ListsServer) UpdateTest(ctx context.Context, in *pb.UpdateTestParams) (*pb.TestParams, error) {
	preTest := converter.TestConverterToDomain(in.Test)
	if err := preTest.Validate(); err != nil {
		return nil, err
	}
	updatedTest, err := s.svcCtx.Test.Update(uint(in.GetId()), preTest)
	if err != nil {
		return nil, err
	}
	out := converter.TestConverterToRPC(updatedTest)
	return out, nil
}

func (s *ListsServer) SoftDeleteTest(ctx context.Context, in *pb.ID) (*pb.TestParams, error) {
	, err := s.svcCtx.Test.SoftDelete(uint(in.GetId()))
	if err != nil {
		return nil, err
	}
	out := converter.TestConverterToDomain(test)
	return out, nil
}

func (s *ListsServer) HardDeleteTest(ctx context.Context, in *pb.ID) (*pb.TestParams, error) {
	test, err := s.svcCtx.Test.HardDelete(uint(in.GetId()))
	if err != nil {
		return nil, err
	}
	out := converter.TestConverterToDomain(test)
	return out, nil
}

func (s *ListsServer) RestoreTest(ctx context.Context, in *pb.ID) (*pb.TestParams, error) {
	test, err := s.svcCtx.Test.Restore(uint(in.GetId()))
	if err != nil {
		return nil, err
	}
	out := converter.TestConverterToDomain(test)
	return out, nil
}
