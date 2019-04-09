#!/usr/bin/env bash
kubectl apply -f namespace.yaml
kubectl label namespace hello istio-injection=enabled
kubectl apply -f helloManifest.yaml
#kubectl apply -f destRules.yaml
kubectl apply -f gatewayNamespaceHello.yaml

echo "*** wait 30 seconds for pods to start ***"
sleep 30

echo "*** Now test: ***"
curl -o /dev/null -s -w "%{http_code}\n" http://0.0.0.0:31380/hello/hello