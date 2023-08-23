#!/usr/bin/env bash

# This script is used to build and start all services and API gateway

# Step 1. build all services and API gateway
echo "Building all services and API gateway ..."
for file in `ls cmd`
do
    if [ -d "cmd/$file" ]
    then
        echo "Building $file"
        cd cmd/$file
        ./build.sh
        cd ../..
    fi
done

# step 2. start docker compose
echo "Starting docker compose..."
docker compose up -d

# step 3. start all services and API gateway
echo "Starting all services and API gateway..."
for file in `ls cmd`
do
    if [ -d "cmd/$file" ]
    then
        sleep 5
        echo "Starting $file"
        screen -dmS $file -t $file cmd/$file/output/bootstrap.sh
    fi
done

# step 4. show all services and API gateway
screen -ls