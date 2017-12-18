#!/bin/bash
# Makes and runs HealthChecker
go build
#if file exists, clear output and run
if [ -f ./healthChecker ]; then
	clear
	env GOOS=windows GOARCH=amd64 go build
fi
