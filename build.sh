#!/bin/bash

GOOS=linux go build -o pls_linux_amd64 .
GOOS=darwin go build -o pls_darwin_amd64 .
GOOS=windows go build -o pls_windows_amd64.exe .

ls | grep pls
