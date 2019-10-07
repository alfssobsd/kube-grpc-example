# Example app for deploy to kube (with grpc)

## Build

### GRPC Server
```
docker build -t your-registry/example-grpc-server -f grpc_server/Dockerfile .
docker push [to your registry]
```
### HTTP Server

Set correct grpc server `address` in `http_server/http_server.go`

Example:
```
const (
	address = "mynamespace-releasename-grpc-server:9000"
)
```

```
docker build -t your-registry/example-http-server -f http_server/Dockerfile .
docker push [to your registry]
```

## Use helm for deploy
Set correct path to registry in helm `helm/grpc-kube-example/values.yml`

Set correct ingress name in `helm/grpc-kube-example/values.yml`
