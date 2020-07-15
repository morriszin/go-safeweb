package safehttp

import (
	"net/http"
)

// Machinery TODO
type Machinery struct {
	h HandleFunc
	d Dispatcher
}

// NewMachinery TODO
func NewMachinery(h HandleFunc, d Dispatcher) *Machinery {
	return &Machinery{h: h, d: d}
}

// HandleRequest TODO
func (m *Machinery) HandleRequest(w http.ResponseWriter, req *http.Request) {
	rw := newResponseWriter(m.d, w)
	ir := newIncomingRequest(req)
	m.h(rw, &ir)
}
