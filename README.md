# gowafp

A Go WAF (Web Application Firewall) that sits between your webserver (nginx)
and your FastCGI application.

nginx <- (tcp) -> gowafp <- (FastCGI) -> PHP-FPM

## notes

### dev

```
docker-compose up --build
```

### build

```
./build.sh
```

### create a machine

```
docker-machine create --driver digitalocean \
--digitalocean-image  ubuntu-16-04-x64 \
--digitalocean-private-networking \
--digitalocean-access-token $DOTOKEN node-1 &
```

### deploy

```
./deploy.sh
```

### checking services

```
docker service ls
docker service ps CONTAINER_ID
```
