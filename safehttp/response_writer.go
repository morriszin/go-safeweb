package safehttp

import (
	"net/http"
)

// ResponseWriter TODO
type ResponseWriter struct {
	d      Dispatcher
	rw     http.ResponseWriter
	header Header
}

func newResponseWriter(d Dispatcher, rw http.ResponseWriter) ResponseWriter {
	header := newHeader(rw.Header())
	return ResponseWriter{d: d, rw: rw, header: header}
}

// Result TODO
type Result struct{}

// Write TODO
func (w *ResponseWriter) Write(resp Response) Result {
	if err := w.d.Write(w.rw, resp); err != nil {
		panic("error")
	}
	return Result{}
}

// WriteTemplate TODO
func (w *ResponseWriter) WriteTemplate(t Template, data interface{}) Result {
	if err := w.d.ExecuteTemplate(w.rw, t, data); err != nil {
		panic("error")
	}
	return Result{}
}

// ServerError TODO
func (w *ResponseWriter) ServerError(code StatusCode, resp Response) Result {
	return Result{}
}

// Header returns the collection of headers that will be set
// on the response. Headers must be changed before Write, WriteTemplate
// or any other similar function is called for the first time.
func (w ResponseWriter) Header() Header {
	return w.header
}

// Dispatcher TODO
type Dispatcher interface {
	Write(rw http.ResponseWriter, resp Response) error
	ExecuteTemplate(rw http.ResponseWriter, t Template, data interface{}) error
}
