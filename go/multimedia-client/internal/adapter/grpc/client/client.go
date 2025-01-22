package client

import (
	"google.golang.org/grpc"
)

// GetConn function returns a new gRPC client connection
func GetConn(uri string, credentials grpc.DialOption) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(uri, credentials)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
