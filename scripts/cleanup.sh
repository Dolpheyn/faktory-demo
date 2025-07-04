#!/bin/bash
set -e

# Check if container exists
if docker ps -a --format '{{.Names}}' | grep -q '^faktory$'; then
  echo "Stopping Faktory..."
  docker stop faktory || echo "Faktory may already be stopped."
  echo "Removing Faktory..."
  docker rm faktory || echo "Faktory may already be removed."
else
  echo "Faktory container not found. Skipping cleanup."
fi
