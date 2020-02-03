package main

import (
	"log"
	"strconv"
	"time"

	handler "handler/function"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/roleypoly/gripkit"
)

func parseIntOrDurationValue(val string, fallback time.Duration) time.Duration {
	if len(val) > 0 {
		parsedVal, parseErr := strconv.Atoi(val)
		if parseErr == nil && parsedVal >= 0 {
			return time.Duration(parsedVal) * time.Second
		}
	}

	duration, durationErr := time.ParseDuration(val)
	if durationErr != nil {
		return fallback
	}
	return duration
}

func main() {
	gk := gripkit.Create(
		gripkit.WithGrpcWeb(
			grpcweb.WithOriginFunc(func(o string) bool { return true }),
		),
		gripkit.WithHTTPOptions(gripkit.HTTPOptions{
			Addr: ":8082",
		}),
	)

	handler.Register(gk.Server)

	err := gk.Serve()
	if err != nil {
		log.Fatalln("gRPC server ended fatally", err)
	}
}
