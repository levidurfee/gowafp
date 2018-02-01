FROM golang:1.8
ADD . /app
WORKDIR /app
RUN \
    go get github.com/tomasen/fcgi_client && \
    go get github.com/microcosm-cc/bluemonday
EXPOSE 8080
CMD ["go", "run", "main.go"]
