#!/bin/sh
echo "Export env..."
source ./env.sh

echo "Building raw script..."
GOOS=linux GOARCH=amd64 go build -o bin/ta2-script-raw
echo "Sending raw script..."
scp -P 8288 bin/ta2-script-raw root@$SERVER_HOST:/usr/local/bin

echo "Building finale script..."
GOOS=linux GOARCH=amd64 go build -o bin/ta2-script-fin
echo "Sending finale script..."
scp -P 8288 bin/ta2-script-fin root@$SERVER_HOST:/usr/local/bin