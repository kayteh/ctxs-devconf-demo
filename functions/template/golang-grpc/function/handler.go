package function

import (
	"context"

	proto "github.com/roleypoly/rpc/ctf"
	"google.golang.org/grpc"
)

type ExampleCTFService struct {
	proto.UnimplementedCTFServer
}

func Register(server *grpc.Server) {
	proto.RegisterCTFServer(server, &ExampleCTFService{})
}

var canaries = []*proto.Canary{
	{Name: "Test-1", Percent: -1},
	{Name: "Test0", Percent: 0},
	{Name: "Test30", Percent: 30},
	{Name: "Test50", Percent: 50},
	{Name: "Test100", Percent: 100},
}

func canaryMax(canaries []*proto.Canary, percentFloor float32) []*proto.Canary {
	matched := []*proto.Canary{}
	for _, canary := range canaries {
		if canary.Percent >= percentFloor {
			matched = append(matched, canary)
		}
	}

	return matched
}

func (e *ExampleCTFService) GetCanaries(ctx context.Context, req *proto.CanaryQuery) (*proto.Canaries, error) {
	return &proto.Canaries{
		Canaries: canaryMax(canaries, req.Threshold),
	}, nil
}
