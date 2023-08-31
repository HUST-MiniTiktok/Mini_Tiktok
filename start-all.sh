#!/usr/bin/env bash

# This script is used to build and start all services and API gateway

# Step 0. check and install screen & ffmpeg
echo "Checking and installing screen ..."
if [ ! -x "$(command -v screen)" ]; then
    echo "Installing screen ..."
    sudo apt-get install screen -y
fi

echo "Checking and installing ffmpeg ..."
if [ ! -x "$(command -v ffmpeg)" ]; then
    echo "Installing ffmpeg ..."
    sudo apt-get install ffmpeg -y
fi

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
echo "Starting docker compose ..."
docker compose up -d

# step 3. start all services and API gateway
echo "Starting all services and API gateway ..."
sleep 10
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
echo "Showing all services and API gateway ..."
screen -ls
echo ""
echo "Hint: You can use 'screen -r <service>' to attach to a screen and use Ctrl+A+D to detach from a screen"
echo "If you want to stop all services and API gateway, please run stop-all.sh"