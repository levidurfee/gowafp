# gowafp

A Go WAF (Web Application Firewall) that sits between your webserver (nginx)
and your FastCGI application.

nginx <- (tcp) -> gowafp <- (FastCGI) -> PHP-FPM
