#!/bin/bash

cd ..

echo "PWD: $(pwd)"

swag init -g cmd/main.go
cp docs/swagger.json docs/doc.json