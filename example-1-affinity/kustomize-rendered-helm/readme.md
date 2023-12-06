`helm template --repo https://contrastsecurity.dev/helm-charts contrast-agent-operator contrast-agent-operator > base/rendered.yaml`
`kubectl apply -k base`
`kubectl apply -k overlays/test`
