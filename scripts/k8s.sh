#!/bin/bash

make -C ..

cd airtable

kubectl apply -f secrets.yaml
helm upgrade --install airtable . -f values.yaml



# helm upgrade --install ingress-nginx ingress-nginx \
#   --repo https://kubernetes.github.io/ingress-nginx \
#   --namespace ingress-nginx \
#   --create-namespace




