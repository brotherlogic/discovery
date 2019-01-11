#!/bin/bash

if [ go test -v ./... | grep FAIL ]; then exit 1; else exit 0; fi
