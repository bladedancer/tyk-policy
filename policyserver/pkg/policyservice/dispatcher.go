package policyservice

import (
	"context"

	"github.com/bladedancer/tyk-policy/policyserver/pkg/coprocess"
)

func (ps *PolicyServer) Dispatch(ctx context.Context, req *coprocess.Object) (*coprocess.Object, error) {
	// See: https://tyk.io/docs/plugins/supported-languages/javascript-middleware/middleware-scripting-guide/
	log.Infof("Dispatch: %+v", req)

	// Change Method
	// req.Request.Method = "GET"

	// Add Headers
	if req.Response != nil {
		req.Response.Headers["custom"] = "response"
	} else {
		req.Request.SetHeaders = make(map[string]string)
		req.Request.SetHeaders["custom"] = "request"
	}

	return req, nil
}
func (ps *PolicyServer) DispatchEvent(ctx context.Context, req *coprocess.Event) (*coprocess.EventReply, error) {
	log.Infof("DispatchEvent: %+v", req)
	return &coprocess.EventReply{}, nil
}
