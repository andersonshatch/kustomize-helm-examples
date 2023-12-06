# Run steps
```bash
kubectl kustomize base --enable-helm | kubectl apply -f -
kubectl kustomize overlays/test/ --enable-helm | kubectl apply -f -
```
