#!/bin/bash

# Check if a log file path is provided as an argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <path_to_log_file>" # Print usage message
    exit 1
fi

LOG_FILE=$1

# Using tail -F to follow the log file even if it is rotated or truncated
tail -F "$LOG_FILE" | while read -r line; do
    # Check if the line contains a HTTP 500 error
    echo "$line" | grep 'HTTP/1.1" 500' &>/dev/null
    if [ $? -eq 0 ]; then
        # Extract the path causing the HTTP 500 error
        HTTP_PATH=$(echo "$line" | grep -oP '(?<=GET |POST ).*?(?= HTTP/1.1")')
        
        # Send an email alert
        mail -s "Alert: HTTP 500 Error Detected" alert@project.com <<< "HTTP 500 on $HTTP_PATH"
    fi
done
