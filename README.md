# certmon
Monitor and track TLS endpoints for expiration dates of certificates

## Why?

Our TLS deployments span multiple certificates across servers and loadbalancers. It's currently PITA to keep track of each of them individually. Existing tools only check the endpoint they happen to resolve to at the time of the test.

This tool gives us a birds eye view of certificate status across endpoints.

PS: endpoint = anything (server or loadbalancer) that might be terminating TLS.

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

```
kubectl create configmap certmon --from-file=config.yaml
kubectl create -f dp.yaml
kubectl create -f svc.yaml
```

## TODO

- Prometheus exporter
- Meaningful logs, with error/warning level messages
- Add concurrency
- Retry failures n (configurable) times
