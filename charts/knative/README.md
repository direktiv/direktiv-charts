# knative

knative for direktiv

![Version: 0.4.4](https://img.shields.io/badge/Version-0.4.4-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.4.0](https://img.shields.io/badge/AppVersion-1.4.0-informational?style=flat-square)

## Additional Information

This chart installs Knative for Direktiv. It configures Knative with correct values in Direktiv's context and adds
 support to provide proxy values for corporate proxies.

### Changes in 0.4.5

- *Upgraded to 1.5.0*

### Changes in 0.4.4

- *Upgraded to 1.4.0*

### Changes in 0.4.3

- *Allowing node affinities for knative pods*

### Changes in 0.4.2

- *Downgrade to knative 1.1.0 because of isssues with private registries*

### Changes in 0.4.1

- *Disable auto-injection of linkerd for all components*
- *Inital scale default is 0, sidecar cpu request set to 50m*
- *SSL_CERT_FILE can be set with controller.ca*

## Installing the Chart

To install the chart with the release name `knative`:

```console
$ helm repo add direktiv https://chart.direktiv.io
$ helm install knative direktiv/knative
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autoscaler.allow_zero_initial_scale | string | `"true"` |  |
| autoscaler.grace_period | string | `"120s"` |  |
| autoscaler.initial_scale | string | `"0"` |  |
| autoscaler.max_scale | string | `"5"` |  |
| autoscaler.retention_period | string | `"120s"` |  |
| controller.ca | string | `"none"` | CA certifcate for self-signed certificate registries |
| defaults.max_timeout_seconds | string | `"7200"` | maximum timeout for knative functions in seconds |
| defaults.revision_cpu_request | string | `"50m"` | cpu requests for direktiv sidecar |
| defaults.timeout_seconds | string | `"900"` | default timeout for knative functions in seconds |
| deployment.skip_tag | string | `"kind.local,ko.local,dev.local,localhost:5000,localhost:31212"` |  |
| http_proxy | string | `""` | HTTP proxy information for knative |
| https_proxy | string | `""` | HTTPS proxy information for knative |
| no_proxy | string | `"localhost,127.0.0.1,10.0.0.0/8,.svc,.cluster.local"` | No proxy information for knative |
| replicas | int | `1` | Replicas for knative components |

