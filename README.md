# GRPC apps for testing load balancing in kubernetes

## Build

### Build GRPC Server
```
docker build -t  your-registry/test-grpc-server -f grpc_server/Dockerfile .
docker push [to your registry]
```

### Build HTTP Server
```
docker build -t your-registry/test-http-server -f http_server/Dockerfile .
docker push [to your registry]
```

## Deploy servers

### Deploy GRPC server
Create your own helm values for deploy  `helm/my-grpc-values.yaml`
```
replicaCount: 3
image:
  repository: your-registry/test-grpc-server
```

Go to helm dir & start deploy
```
cd helm
helm install grpc-server --tiller-namespace your-tiller-namespace --namespace your-kube-namespace --name grpc-server-test -f my-grpc-values.yaml
```

### Deploy http server
Create your own helm values for deploy  `helm/my-http-values.yaml`
```
grpcServiceName: grpc-server-test:9000
replicaCount: 1
image:
  repository: your-registry/test-http-server
ingress:
  enabled: true
  paths:
    - /
  hosts:
    - http-server-test.yourdomain.net
  tls: []

```
Go to helm dir & start deploy
```
cd helm
helm install http-server --tiller-namespace your-tiller-namespace --namespace your-kube-namespace --name http-server-test -f my-http-values.yaml
```

## Test servers
```
curl http://http-server-test.yourdomain.net/
stdout>ServerName = grpc-server-test-56d9cc6b6d-4qptt
```

if the value doesn't change all the time, you do not have load balancing GRPC