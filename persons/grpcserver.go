package persons

import (
	"context"
)

// GRPCServer ...
type GRPCServer struct {
}

func (s *GRPCServer) mustEmbedUnimplementedPersonsServer() {
	panic("implement me")
}

func (s *GRPCServer) Person(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	return &TestResponse{
		Name:  req.Nickname,
		Id:    1010101010,
		Email: "World@Is.Big",
	}, nil
}
