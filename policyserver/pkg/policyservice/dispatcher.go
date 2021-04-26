package policyservice

import (
	"context"

	"github.com/bladedancer/tyk-policy/policyserver/pkg/coprocess"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ps *PolicyServer) Dispatch(ctx context.Context, req *coprocess.Object) (*coprocess.Object, error) {
	log.Infof("%v", req)
	return nil, status.Errorf(codes.Unimplemented, "mine: method Dispatch not implemented")
}
func (ps *PolicyServer) DispatchEvent(ctx context.Context, req *coprocess.Event) (*coprocess.EventReply, error) {
	log.Infof("%v", req)
	return nil, status.Errorf(codes.Unimplemented, "mine:  method DispatchEvent not implemented")
}
