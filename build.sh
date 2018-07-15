#!/bin/bash

go get github.com/twitchtv/retool
retool sync
go generate
go build
