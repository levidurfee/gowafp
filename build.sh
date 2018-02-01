#!/bin/bash

docker build -f Docker/nginx/Dockerfile -t levidurfee/gowafp_nginx:latest .
docker build -f Docker/Go/Dockerfile -t levidurfee/gowafp_fw:latest .
docker build -f Docker/php/Dockerfile -t levidurfee/gowafp_php:latest .

docker push levidurfee/gowafp_nginx:latest
docker push levidurfee/gowafp_fw:latest
docker push levidurfee/gowafp_php:latest
