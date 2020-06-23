#!/usr/bin/env bash
docker run --publish 80:8080 -v "$(pwd)/logs:/app/logs" -d --name rete-challenge rete-challenge:latest
