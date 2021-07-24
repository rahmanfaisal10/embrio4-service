#!/bin/bash

start:
	@go build -o bin/server .
	@bin/server -env=local

build:
	@go build -o bin/server .
