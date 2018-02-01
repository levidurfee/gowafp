#!/bin/bash
docker stack deploy -c docker-compose.yml --with-registry-auth gowafp
