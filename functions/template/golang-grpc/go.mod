module handler

go 1.13

replace handler/function => ./function

require (
	github.com/improbable-eng/grpc-web v0.12.0
	github.com/roleypoly/gripkit v0.0.0-20200110031819-2c4872b1be7c
	google.golang.org/grpc v1.26.0
	handler/function v0.0.0-00010101000000-000000000000
)
