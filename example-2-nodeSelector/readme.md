# Aim

- Deploy the [Contrast Agent Operator Helm chart](https://contrastsecurity.dev/helm-charts)
- Add `nodeSelector` value(s) to the `deployment/contrast-agent-operator`.

# Approaches

## [terraform-run-helm-kustomize](terraform-run-helm-kustomize)
Terraform will run Helm to deploy the chart, executing a post-render step that adds the `nodeSelector` value(s) via kustomize.

## [terraform-run-helm-go](terraform-run-helm-go)
Terraform will run Helm to deploy the chart, executing a post-render step that adds the `nodeSelector` value(s) via a go script which can be built and stored in the repo.

### Links
- https://registry.terraform.io/providers/hashicorp/helm/latest/docs/resources/release
- https://helm.sh/docs/topics/advanced/#post-rendering

