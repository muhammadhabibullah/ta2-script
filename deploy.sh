#!/bin/sh
echo "Export env..."
source ./env.sh

echo "Building binary..."
GOOS=linux GOARCH=amd64 go build -o bin/ta2-script
echo "Sending binary..."
scp -P 8288 bin/ta2-script root@$SERVER_HOST:/usr/local/bin
