#!/bin/bash
set -e


if docker ps -a --format '{{.Names}}' | grep -q '^faktory$'; then
  echo "Faktory already started"
else
  echo "Starting Faktory..."
  docker run -d \
    --name faktory \
    -v "$(pwd)/data:/var/lib/faktory/db" \
    -e "FAKTORY_PASSWORD=some_password" \
    -p 127.0.0.1:7419:7419 \
    -p 127.0.0.1:7420:7420 \
    contribsys/faktory:latest \
    /faktory -b :7419 -w :7420 -e production

  echo "Waiting for Faktory to be ready..."
  sleep 1
fi

echo "Running Go app..."
go run main.go
