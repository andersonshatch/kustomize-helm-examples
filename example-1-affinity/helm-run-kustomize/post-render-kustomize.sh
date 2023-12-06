#!/usr/bin/env bash

cat <&0 > all.yaml

kubectl kustomize . && rm all.yaml
