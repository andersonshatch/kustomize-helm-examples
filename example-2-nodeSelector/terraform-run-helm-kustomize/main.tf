provider "helm" {
  #required to ensure that changes to the chart or post-render step are picked up by terraform. See below helm-provider issues:
  # https://github.com/hashicorp/terraform-provider-helm/issues/675
  # https://github.com/hashicorp/terraform-provider-helm/issues/346
  experiments {
    manifest = true
  }
}


resource "helm_release" "contrast_operator" {
  name       = "contrast-agent-operator"
  repository = "https://contrastsecurity.dev/helm-charts"
  chart      = "contrast-agent-operator"
  values     = ["${file("values.yaml")}"]
  postrender {
    binary_path = "./post-render-kustomize.sh"
  }
}
