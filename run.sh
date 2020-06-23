#!/usr/bin/env bash
docker run --publish 8080:8080 -v "$(pwd)/logs:/app/logs" -d --name rete-challenge rete-challenge:latest
