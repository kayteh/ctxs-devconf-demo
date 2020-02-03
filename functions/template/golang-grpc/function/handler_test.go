package function

import (
	"context"
	proto "github.com/roleypoly/rpc/ctf"
	"testing"
)

func TestGetCanaries(t *testing.T) {
	e := ExampleCTFService{}

	testCases := []struct {
		desc      string
		threshold float32
	}{
		{
			desc:      "threshold = 0%",
			threshold: 0,
		},
		{
			desc:      "threshold = 50%",
			threshold: 50,
		},
		{
			desc:      "threshold = 100%",
			threshold: 100,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			canaries, err := e.GetCanaries(context.TODO(), &proto.CanaryQuery{
				Threshold: tC.threshold,
			})
			if err != nil {
				t.Error(err)
			}

			for _, canary := range canaries.Canaries {
				if canary.Percent < tC.threshold {
					t.Error("Returned out-of-threshold canary")
				}
			}
		})
	}
}
