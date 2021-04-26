#!/bin/sh

echo ======================
echo === Create Cluster ===
echo ======================

k3d cluster delete tykgw
k3d cluster create --no-lb --update-default-kubeconfig=false --wait tykgw
export KUBECONFIG=$(k3d kubeconfig write tykgw)
kubectl cluster-info

docker pull k8s.gcr.io/redis:e2e
docker pull tykio/tyk-gateway:v3.1.0

k3d image import --cluster tykgw k8s.gcr.io/redis:e2e
k3d image import --cluster tykgw tykio/tyk-gateway:v3.1.0

echo ======================
echo === Installing Tyk ===
echo ======================
kubectl apply -f deploy/tyk

echo "RUN:"
echo "export KUBECONFIG=$(k3d kubeconfig write tykgw)"
