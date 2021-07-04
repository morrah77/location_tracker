#Simple in-memory location history server.

#Build

Running `go build` in the root directory of this module produces a working binary.

The server serves HTTPS at the address specified in the environment variable HISTORY_SERVER_LISTEN_ADDR. If this environment variable is not set, the listen address defaults to :8080.

Clients are able to speak JSON over HTTP to the server. The three endpoints it supports are:

PUT /location/{order_id}
GET /location/{order_id}?max=<N>
DELETE /location/{order_id}

Details about the endpoints:

PUT /location/{order_id} - append a location to the history for the specified order.

Example interaction:

```
PUT /location/def456
{
  "lat": 12.34,
  "lng": 56.78
}

200 OK
```

GET /location/{order_id}?max=<N> - Retrieve at most N items of history for the specified order. The most recent locations (in chronological order of insertion) should be returned first, if history is truncated by the max parameter.

Example interaction:

```
GET /location/abc123?max=2

200 OK
{
  "order_id": "abc123",
  "history": [
    {"lat": 12.34, "lng": 56.78},
    {"lat": 12.34, "lng": 56.79}
  ]
}
```

The max query parameter may or may not be present. If it is not present, the entire history should be returned.


DELETE /location/{order_id} - delete history for the specified order.

Example interaction:

```
DELETE /location/xyz987

200 OK
```

#Run

```
./location_tracker
```

Or

```
HISTORY_SERVER_LISTEN_ADDR=:8081 ./location_tracker
```

#Test

##Unit tests (TODO: add mmore tests)

```
go test ./...
```

##Manually

```
HISTORY_SERVER_LISTEN_ADDR=:8081 ./location_tracker


curl --insecure -iv -X GET https://127.0.0.1:8081/location/123abc

404 {"error":"No entities found!"}


curl --insecure -iv -X GET https://127.0.0.1:8081/location/124abc

404 {"error":"No entities found!"}


curl --insecure -iv -X DELETE https://127.0.0.1:8081/location/123abc

404 {"error":"No entities found!"}


curl --insecure -iv -X PUT https://127.0.0.1:8081/location/123abc -d '{"lat":11.23,"lng":41.22}'

200


curl --insecure -iv -X GET https://127.0.0.1:8081/location/123abc

200 {"order_id":"123abc","history":[{"lat":11.23,"lng":41.22}]}


curl --insecure -iv -X GET https://127.0.0.1:8081/location/124abc

404 {"error":"No entities found!"}


curl --insecure -iv -X PUT https://127.0.0.1:8081/location/123abc -d '{"lat":11.23,"lng":42.22}'

200


curl --insecure -iv -X PUT https://127.0.0.1:8081/location/123abc -d '{"lat":11.23,"lng":43.22}'

200


curl --insecure -iv -X GET https://127.0.0.1:8081/location/123abc

200 {"order_id":"123abc","history":[{"lat":11.23,"lng":43.22},{"lat":11.23,"lng":42.22},{"lat":11.23,"lng":41.22}]}


curl --insecure -iv -X GET https://127.0.0.1:8081/location/123abc?max=1

200 {"order_id":"123abc","history":[{"lat":11.23,"lng":43.22}]}


curl --insecure -iv -X GET https://127.0.0.1:8081/location/123abc?max=2

200 {"order_id":"123abc","history":[{"lat":11.23,"lng":43.22},{"lat":11.23,"lng":42.22}]}


curl --insecure -iv -X GET https://127.0.0.1:8081/location/123abc?max=3

200 {"order_id":"123abc","history":[{"lat":11.23,"lng":43.22},{"lat":11.23,"lng":42.22},{"lat":11.23,"lng":41.22}]}


curl --insecure -iv -X GET https://127.0.0.1:8081/location/123abc?max=99

200 {"order_id":"123abc","history":[{"lat":11.23,"lng":43.22},{"lat":11.23,"lng":42.22},{"lat":11.23,"lng":41.22}]}


curl --insecure -iv -X GET https://127.0.0.1:8081/location/124abc

404 {"error":"No entities found!"}


curl --insecure -iv -X DELETE https://127.0.0.1:8081/location/123abc

200


curl --insecure -iv -X GET https://127.0.0.1:8081/location/123abc

404 {"error":"No entities found!"}


curl --insecure -iv -X GET https://127.0.0.1:8081/location/124abc

404 {"error":"No entities found!"}
```