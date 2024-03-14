# HTTP-Status-On-Demand

This is a simple HTTP server that can be used to simulate different HTTP status codes.
It can be a easy hack during testing.


## Install

```bash
go install github.com/TheBigRoomXXL/http-status-on-demand@latest
```


## Use

```bash
# Start Server
http-status-on-demand       # Run on port 8000 by default
http-status-on-demand :4000 # Run on port 4000

# Send request
curl -i http://localhost:8000/200
curl -i http://localhost:8000/404
curl -i http://localhost:8000/efefwfzf # will return 400
```
