package policyservice

import (
	"fmt"
	"net"

	"github.com/bladedancer/tyk-policy/policyserver/pkg/coprocess"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type PolicyServer struct {
	Config *PolicyServiceConfig
	done   chan struct{}
}

func (ps *PolicyServer) Run() error {
	ps.done = make(chan struct{})
	ps.runGRPC()
	select {
	case <-ps.done:
	}
	return nil
}

func (ps *PolicyServer) runGRPC() (*grpc.Server, error) {
	bind := fmt.Sprintf("%s:%d", ps.Config.Host, ps.Config.Port)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", bind)
	if err != nil {
		return nil, err
	}

	coprocess.RegisterDispatcherServer(grpcServer, ps)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
		ps.done <- struct{}{}
	}()

	log.Infof("GRPC Listening on %s", bind)

	return grpcServer, nil
}
