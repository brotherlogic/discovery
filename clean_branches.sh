git branch -r | grep -Ev 'HEAD|main|develop' | awk -F/ '{print }' | xargs -r git push origin -d
