#!/usr/bin/env bash
go build -ldflags "-w -s" main.go
mv -f main agent-proxy
upx --brute agent-proxy