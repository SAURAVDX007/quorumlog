# quorumlog
A distributed, replicated commit log in Go — gRPC, Serf-based service discovery, and Raft consensus


Test current implementation with ->
Note: []byte are base64 encoded bcz of Go's encoding/json package

POST requests ->
curl -X POST localhost:8080 -d \
'{"record": {"value": "SGVsbG8gcXVvcnVtbG9n"}}'

curl -X POST localhost:8080 -d \
'{"record": {"value": "bXkgZmlyc3QgY29tbWl0IGxvZyBtZXNzYWdl"}}'

curl -X POST localhost:8080 -d \
'{"record": {"value": "YW5vdGhlciBtZXNzYWdl"}}'


GET requests ->
curl -X GET localhost:8080 -d '{"offset":0}'

curl -X GET localhost:8080 -d '{"offset":1}'

curl -X GET localhost:8080 -d '{"offset":2}'