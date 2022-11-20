package server

import "google.golang.org/grpc"

type GrpcAPIServer struct {
	*grpc.Server
	address string
}
