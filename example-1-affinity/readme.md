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

# Approaches

## [helm-run-kustomize](helm-run-kustomize) 
As the name suggests, has Helm run kustomize, as a post-render step.
Over the other methods, this has the advantage of Helm knowing about the deployment, allowing it to do upgrades, uninstalls etc.,

### Links
- https://helm.sh/docs/topics/advanced/#post-rendering
- https://github.com/thomastaylor312/advanced-helm-demos/tree/master/post-render
- https://austindewey.com/2020/07/27/patch-any-helm-chart-template-using-a-kustomize-post-renderer/

## [kustomize-rendered-helm](kustomize-rendered-helm)
Have Helm fetch and render the chart and then kustomize the changes onto it. Rendered chart should likely be added to version control for production use.

### Links
- https://github.com/kubernetes-sigs/kustomize/blob/master/examples/chart.md#performance

## [kustomize-run-helm](kustomize-run-helm)
Have kustomize run Helm to fetch and render the chart, then kustomize changes onto it. Chart is stored locally and should likely be added to version control for production use.

### Links
- https://github.com/kubernetes-sigs/kustomize/blob/master/examples/chart.md
