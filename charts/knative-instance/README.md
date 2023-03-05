# knative-instance

Helm chart to install a pre-configured Knative instance for Direktiv

![Version: 1.9.2](https://img.shields.io/badge/Version-1.9.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v1.9.2](https://img.shields.io/badge/AppVersion-v1.9.2-informational?style=flat-square)

## Additional Information

This chart installs a Knative Service instance for the Knative Operator and Contour as the network component.

## Installing the Chart

To install the chart with the release name `knative`:

```console
$ helm repo add direktiv https://chart.direktiv.io
$ helm install knative direktiv/knative-instance
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| certificate | string | `""` | Custom certificate for controller. This needs to be a secret create before installation in the knative-serving namespace |
| http_proxy | string | `""` | Proxy settings for controller |
| https_proxy | string | `""` |  |
| no_proxy | string | `""` |  |
| replicas | int | `1` | Knative replicas |
| skip-digest | string | `"kind.local,ko.local,dev.local,localhost:5000,localhost:31212"` | Repositories which skip digest resolution |

