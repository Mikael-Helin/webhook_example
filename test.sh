#!/bin/bash

# Generating timestamp
timestamp=$(date +%s)

# Your message
message="test_message"

# Your secret
secret="your_secret_here"

# Generating hash
hash=$(echo -n "${secret}${timestamp}${message}" | sha256sum | awk '{print $1}')

#curl -X POST "http://localhost:1113/hook" -d "{\"timestamp\": $timestamp, \"message\": \"$message\", \"hash\": \"$hash\"}"
curl -v -X POST "http://localhost:1113/hook" -d "{\"timestamp\": $timestamp, \"message\": \"$message\", \"hash\": \"$hash\"}"
