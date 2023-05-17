#!/bin/bash

# Base URL of the repo's API
base_url="http://localhost:8000"

# Test GET request
echo "Testing GET request..."
response=$(curl -s -o /dev/null -w "%{http_code}" $base_url/movies)
if [ "$response" -ne 200 ]; then
	echo "GET request failed with status: $response"
else
	echo "GET request suceeded with status: $response"
fi

# Test POST request
echo "Testing POST request..."
response=$(curl -s -o /dev/null -w "%{http_code}" -X POST -d @data.json -H "Content-Type: application/json" $base_url/movies)
if [ "$response" -ne 201 ] && [ "$response" -ne 200 ]; then
    echo "POST request failed with status: $response"
else
    echo "POST request succeeded with status: $response"
fi
