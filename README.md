# certmon
Monitor and track TLS endpoints for expiration dates of certificates

## Usage

### From source

```
go install github.com/turbobytes/certmon/cmd/certmon
$GOPATH/bin/certmon -config $GOPATH/src/github.com/turbobytes/certmon/example_config.yaml -ui $GOPATH/src/github.com/turbobytes/certmon/assets/index.html
```

open http://127.0.0.1:8081/

### Docker

```
# Create config.yaml somewhere, look at example_config.yaml
docker run -p 8082:8082 -v /path/to/config.yaml:/config.yaml --rm -it turbobytes/certmon
```

open http://127.0.0.1:8082/

### Kubernetes

TODO


## TODO

- Prometheus exporter
- Meaningful logs, with error/warning level messages
- Add concurrency
- Retry failures n (configurable) times
