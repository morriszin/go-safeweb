package safehttp

import "net/http"

// IncomingRequest TODO
type IncomingRequest struct {
	req    *http.Request
	Header Header
}

func newIncomingRequest(req *http.Request) IncomingRequest {
	return IncomingRequest{req: req, Header: newHeader(req.Header)}
}
