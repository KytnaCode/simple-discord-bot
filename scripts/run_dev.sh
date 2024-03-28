#!/usr/bin/sh

# All paths are relative to project's root.

docker compose -f ./build/docker-compose.yaml down --volumes
docker compose -f ./build/docker-compose.yaml up --build
