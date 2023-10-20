#!/bin/bash
git branch -r | grep -Ev 'HEAD|main|develop' | awk -F/ '{print $2}' | xargs -r git push origin -d
