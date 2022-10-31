#!/bin/sh

# read the .env line by line
while read line; do
    # export the variable
    export $line
done < .env

# run the app
./clipboardsync start
