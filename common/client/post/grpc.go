package post

import (
	"common/proto/post"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewProfileClient()(client post.PostServiceClient,close func()error,err error){
	grpcAddr := os.Getenv("TRAINER_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env TRAINER_GRPC_ADDR")
	}
	opts, err := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}
	return post.NewPostServiceClient(conn),conn.Close,nil
}

func grpcDialOpts(grpcAddr string) ([]grpc.DialOption, error) {
	if noTLS, _ := strconv.ParseBool(os.Getenv("GRPC_NO_TLS")); noTLS {
		return []grpc.DialOption{grpc.WithInsecure()}, nil
	}

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.New("cannot load root CA cert")
	}
	creds := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	return []grpc.DialOption{
		grpc.WithTransportCredentials(creds),		
	}, nil
}
