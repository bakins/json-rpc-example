# json-rpc-example
Example JSON-RPC over HTTP server in Go using Gorilla.

This is a trimmed down example showing how one may use Gorilla
JSON-RPC with the mux and middleware.

## Usage

### Install

```
go install github.com/bakins/json-rpc-example
```

Start the  server:

```
json-rpc-example
```

### Make RPC calls using curl

```
curl  -H "Content-Type: application/json"  -d '{"method":"Arith.Divide","params":[{"A": 10, "B":2}], "id": 0}' http://localhost:8080/rpc
```

```
{"id":0,"result":{"Quo":5,"Rem":0},"error":null}
```

Also supports compressed responses:

```
curl -sv --compressed -H "Content-Type: application/json" -d '{"method":"Arith.Divide","params":[{"A": 10, "B":2}], "id": 12}' http://localhost:8080/rpc
```

```
...
< HTTP/1.1 200 OK
< Content-Encoding: deflate
< Content-Type: application/json; charset=utf-8
< Vary: Accept-Encoding
< X-Content-Type-Options: nosniff
< Date: Wed, 21 Jan 2015 12:26:11 GMT
< Content-Length: 56
<
{"result":{"Quo":5,"Rem":0},"error":null,"id":12}
```

## Additional features

[expvar](http://golang.org/pkg/expvar) and
[pprof](http://golang.org/pkg/net/http/pprof/) are availible at there
normal endpoints (/debug/vars, etc).


## Acknowledgements

Inspired heavily by https://github.com/kelseyhightower/jsonrpc-server
