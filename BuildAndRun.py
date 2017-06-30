import os
import subprocess

name = "discovery"

current_hash = ""
if os.path.isfile('hash'):
    current_hash = open('hash').readlines()[0]
new_hash = os.popen('git rev-parse HEAD').readlines()[0]
open('hash','w').write(new_hash)
    
# Move the old version over
for line in os.popen('cp ' + name + ' old' + name).readlines():
    print line.strip()

# Rebuild
for line in os.popen('go build ./...').readlines():
    print line.strip()

# Rebuild
for line in os.popen('go build').readlines():
    print line.strip()

size_1 = os.path.getsize('./old' + name)
size_2 = os.path.getsize('./' + name)

running = len(os.popen('ps -ef | grep ' + name).readlines()) > 3

              
if size_1 != size_2 or new_hash != current_hash or not running:
    if not running:
        for line in os.popen('cat out.txt | mail -s "Crash Report ' + name + '" brotherlogic@gmail.com').readlines():
            pass
    for line in os.popen('echo "" > out.txt').readlines():
        pass
    for line in os.popen('killall ' + name).readlines():
        pass
    subprocess.Popen(['./' + name])
