# Aim

- Deploy the [Contrast Agent Operator Helm chart](https://contrastsecurity.dev/helm-charts)
- Adjust the `deployment/contrast-agent-operator` affinity with an arbitrary example, removing all of `podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution`, replacing it with `podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution` so it should not run in the same host as an app with label `app=example-webgoat`

Before:
```yaml
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: kubernetes.io/os
            operator: In
            values:
            - linux
          - key: kubernetes.io/arch
            operator: In
            values:
            - amd64
    podAntiAffinity:
      preferredDuringSchedulingIgnoredDuringExecution:
      - podAffinityTerm:
          labelSelector:
            matchExpressions:
            - key: app.kubernetes.io/name
              operator: In
              values:
              - operator
            - key: app.kubernetes.io/part-of
              operator: In
              values:
              - contrast-agent-operator
          topologyKey: kubernetes.io/hostname
        weight: 100
```

After:
```yaml
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: kubernetes.io/os
            operator: In
            values:
            - linux
          - key: kubernetes.io/arch
            operator: In
            values:
            - amd64
    podAntiAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
      - labelSelector:
          matchLabels:
            app: example-webgoat
        topologyKey: kubernetes.io/hostname
```
