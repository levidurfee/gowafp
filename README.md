# gowafp

[![Build Status](https://travis-ci.org/levidurfee/gowafp.svg?branch=master)](https://travis-ci.org/levidurfee/gowafp)

A Go WAF (Web Application Firewall) that sits between your webserver (nginx)
and your FastCGI application.

nginx <- (tcp) -> gowafp <- (FastCGI) -> PHP-FPM

The goal of this package is to prevent any attacks from reaching the FastCGI
application. It should block all
[SQL injection](https://www.owasp.org/index.php/SQL_Injection)
attempts and filter
[XSS](https://www.owasp.org/index.php/Cross-site_Scripting_(XSS))
attempts.
Maybe down the road it could also handle
[CSRF](https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)).

The Application and the Web Application Firewall should not be on different
servers. While the services out there do a good job, they're expensive and
slower. Of course, you could recompile your web server with some additional
features, but that is harder to deploy while scaling.

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
