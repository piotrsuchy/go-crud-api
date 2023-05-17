#!/bin/bash

# check for go compiler
if ! command -v go &> /dev/null
then
	echo "Go could not be found. Please install it first to continue"
	exit
fi

# navigate to main directory
cd ..
cd cmd/main

go build main.go

./main
