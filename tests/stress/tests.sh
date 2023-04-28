#!/bin/sh
# Stress Test 

# Reading from Environment Variables
HOST=http://$KV_SERVER_HOST:$KV_SERVER_PORT

# Run Tests
echo "Running Tests"

ddosify -t $HOST/api/kv/{{_randomColor}} -n 10000 -d 200 -m POST -b {{_randomInt}} 
ddosify -t $HOST/api/kv/{{_randomColor}} -n 1000 -d 20 -m GET 
ddosify -t $HOST/api/kv/{{_randomColor}} -n 1000 -d 20 -m DELETE

