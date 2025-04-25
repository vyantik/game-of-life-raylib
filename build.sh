#!/bin/bash

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o game_of_life.exe

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o game_of_life

echo "Done!"