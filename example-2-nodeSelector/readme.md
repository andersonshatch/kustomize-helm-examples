# Aim

- Deploy the [Contrast Agent Operator Helm chart](https://contrastsecurity.dev/helm-charts)
- Add `nodeSelector` value(s) to the `deployment/contrast-agent-operator`.

# Approaches

## [terraform-run-helm](terraform-run-helm)
Terraform will run Helm to deploy the chart, executing a post-render step that adds the `nodeSelector` value(s).

### Links
- https://registry.terraform.io/providers/hashicorp/helm/latest/docs/resources/release
- https://helm.sh/docs/topics/advanced/#post-rendering

