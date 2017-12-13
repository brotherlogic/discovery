#!/bin/bash
grep log.Print * -Rl | grep .go$ | grep -v _test.go
RESULT=$?
if [ $RESULT != 1 ]; then
    exit 1
fi
