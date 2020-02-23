#!/bin/sh
echo "Export env..."
source ./env.sh

echo "Building first binary..."
GOOS=linux GOARCH=amd64 go build -o bin/ta2-script-raw
echo "Sending first binary..."
scp -P 8288 bin/ta2-script-raw root@$SERVER_HOST:/usr/local/bin

echo "Building second binary..."
GOOS=linux GOARCH=amd64 go build -o bin/ta2-script-fin
echo "Sending second binary..."
scp -P 8288 bin/ta2-script-fin root@$SERVER_HOST:/usr/local/bin