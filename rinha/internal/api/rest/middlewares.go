package rest

import "net/http"

type ResponseHeader struct {
	handler     http.Handler
	headerName  string
	headerValue string
}

func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(rh.headerName, rh.headerValue)
	rh.handler.ServeHTTP(w, r)
}

func NewResponseHeader(handlerToWrap http.Handler, headerName, headerValue string) *ResponseHeader {
	return &ResponseHeader{handlerToWrap, headerName, headerValue}
}
