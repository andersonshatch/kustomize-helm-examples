#!/usr/bin/env bash

#take in chart contents from helm
cat <&0 > all.yaml

#kustomize the chart and remove generated file, helm will deploy the output
kubectl kustomize . && rm all.yaml
