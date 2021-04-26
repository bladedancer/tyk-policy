#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

export APIID=`curl -s -H "x-tyk-authorization: foo" http://localhost:8080/tyk/apis/ -d @$DIR/config/my-api -X POST | jq .key`
curl -s -H "x-tyk-authorization: foo" http://localhost:8080/tyk/reload
export APIKEY=`curl -s -H "x-tyk-authorization: foo" http://localhost:8080/tyk/keys/create -d @$DIR/config/key-of-my-api | jq .key`

echo Auth Key: $APIID
echo Auth Key: $APIKEY

curl  http://localhost:8080/my-api/ip -H "Authorization: $APIKEY"

curl  -H "x-tyk-authorization: foo" http://localhost:8080/tyk/apis/$APIID -X DELETE