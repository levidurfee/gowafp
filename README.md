# gowafp

[![Build Status](https://travis-ci.org/levidurfee/gowafp.svg?branch=master)](https://travis-ci.org/levidurfee/gowafp)

A Go WAF (Web Application Firewall) that sits between your webserver (nginx)
and your FastCGI application.

nginx <- (tcp) -> gowafp <- (FastCGI) -> PHP-FPM

## usage

First, you need to have Go get the repo.

```Shell
go get github.com/levidurfee/gowafp
```

Below is simple `main.go` example.

```Go
package main

import (
	"github.com/levidurfee/gowafp"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", gowafp.AnalyzeRequest(gowafp.PhpHandler("/app/index.php", "tcp", "127.0.0.1:9000")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Then build and run it.

```Shell
go build main.go
./main
```
