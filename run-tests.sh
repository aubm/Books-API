#!/bin/bash

docker build -t aubm/books-api .
docker-compose up -d --force-recreate
docker-compose ps
echo "waiting for a few seconds for the app to be ready..."
sleep 40
go test ./...
newman -c postman/collection.json -e postman/environnement.json -d postman/data.json -x
