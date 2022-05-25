package http

import "net/http"

type ResponseWriter interface {
    WriteHttpResponse(writer http.ResponseWriter) error
}
