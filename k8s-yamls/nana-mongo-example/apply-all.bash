#! /bin/bash

kubectl apply -f=mongo-secret.yaml
kubectl apply -f=mongo-configmap.yaml
kubectl apply -f=mongo-database.yaml
kubectl apply -f=mongo-express.yaml
