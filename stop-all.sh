#!/usr/bin/env bash

# This script is used to stop all services and API gateway

# Step 1. stop all services and API gateway
echo "Stopping all services and API gateway..."
for file in `ls cmd`
do
    if [ -d "cmd/$file" ]
    then
        echo "Stoping $file"
        screen -S $file -X quit
    fi
done

# Step 2. stop docker compose
echo "Stopping docker compose..."
docker compose down