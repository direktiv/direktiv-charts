# knative

knative for direktiv

![Version: 0.4.0](https://img.shields.io/badge/Version-0.4.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.1.0](https://img.shields.io/badge/AppVersion-1.1.0-informational?style=flat-square)

## Additional Information

This chart installs Knative for Direktiv. It configures Knative with correct values in Direktiv's context and adds
 support to provide proxy values for corporate proxies.

## Installing the Chart

To install the chart with the release name `knative`:

```console
$ helm repo add direktiv https://charts.direktiv.io
$ helm install knative direktiv/knative
```

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://kubernetes.github.io/ingress-nginx | ingress-nginx | 4.0.13 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autoscaler.grace_period | string | `"120s"` |  |
| autoscaler.max_scale | string | `"5"` |  |
| autoscaler.retention_period | string | `"120s"` |  |
| defaults.max_timeout_seconds | string | `"7200"` | maximum timeout for knative functions in seconds |
| defaults.timeout_seconds | string | `"900"` | default timeout for knative functions in seconds |
| deployment.skip_tag | string | `"kind.local,ko.local,dev.local,localhost:5000,localhost:31212"` |  |
| http_proxy | string | `""` | HTTP proxy information for knative |
| https_proxy | string | `""` | HTTPS proxy information for knative |
| ingress-nginx | object | `{"controller":{"admissionWebhooks":{"patch":{"podAnnotations":{"linkerd.io/inject":"disabled"}}},"replicaCount":1,"service":{"ports":{"http":9090,"https":9443}}}}` | nginx ingress controller configuration |
| no_proxy | string | `"localhost,127.0.0.1,10.0.0.0/8,.svc,.cluster.local"` | No proxy information for knative |
| replicas | int | `1` | Replicas for knative components |

