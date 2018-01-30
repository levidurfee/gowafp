FROM golang:1.8
ADD . /app
WORKDIR /app
RUN \
    go get github.com/kellegous/fcgi
EXPOSE 8080
CMD ["go", "run", "main.go"]
