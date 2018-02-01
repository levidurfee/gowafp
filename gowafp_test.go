package gowafp

import (
    "testing"
    "net/http"
)

func TestPhpHandler(t *testing.T) {
    phpHandler := PhpHandler()
    _, ok := phpHandler.(http.Handler)
    if !ok {
        t.Error("Not a http handler")
    }
}

func TestAnalyzeRequest(t *testing.T) {
    handler := AnalyzeRequest(PhpHandler())
    _, ok := handler.(http.Handler)
    if !ok {
        t.Error("Not a http handler")
    }
}
