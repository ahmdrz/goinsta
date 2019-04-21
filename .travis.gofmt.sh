#!/bin/bash
if [ -n "$(gofmt -l .)" ]; then
    echo "GoInsta is not formatted !"
    gofmt -d .
    exit 1
else
    echo "GoInsta is well formatted ;)"
fi
