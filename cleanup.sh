#!/bin/bash
set -e

echo "Stopping Faktory..."
docker stop faktory || echo "Faktory already stopped or not found."

echo "Removing Faktory container..."
docker rm faktory || echo "Faktory already removed or not found."
