package gowafp

import (
	"net/http"
	"testing"
)

func TestPossibleSqlInjection(t *testing.T) {
	result := possibleSqlInjection("1'or'1'='1")
	if result != true {
		t.Error("Possible SQL Injection failed")
	}
}

func TestPhpHandler(t *testing.T) {
	phpHandler := PhpHandler("/app/index.php", "tcp", "127.0.0.1:9000")
	_, ok := phpHandler.(http.Handler)
	if !ok {
		t.Error("Not a http handler")
	}
}

func TestAnalyzeRequest(t *testing.T) {
	handler := AnalyzeRequest(PhpHandler("/app/index.php", "tcp", "127.0.0.1:9000"))
	_, ok := handler.(http.Handler)
	if !ok {
		t.Error("Not a http handler")
	}
}
