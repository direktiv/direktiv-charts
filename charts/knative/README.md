# knative

knative for direktiv

![Version: 0.4.5](https://img.shields.io/badge/Version-0.4.5-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.5.0](https://img.shields.io/badge/AppVersion-1.5.0-informational?style=flat-square)

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
| activator.containers.activator.resources.limits.cpu | string | `"1"` |  |
| activator.containers.activator.resources.limits.memory | string | `"600Mi"` |  |
| activator.containers.activator.resources.requests.cpu | string | `"300m"` |  |
| activator.containers.activator.resources.requests.memory | string | `"60Mi"` |  |
| activator.containers.autoscaler.resources.limits.cpu | string | `"1"` |  |
| activator.containers.autoscaler.resources.limits.memory | string | `"1000Mi"` |  |
| activator.containers.autoscaler.resources.requests.cpu | string | `"100m"` |  |
| activator.containers.autoscaler.resources.requests.memory | string | `"100Mi"` |  |
| activator.containers.controller.resources.limits.cpu | string | `"1"` |  |
| activator.containers.controller.resources.limits.memory | string | `"1000Mi"` |  |
| activator.containers.controller.resources.requests.cpu | string | `"100m"` |  |
| activator.containers.controller.resources.requests.memory | string | `"100Mi"` |  |
| activator.containers.domainmapping.resources.limits.cpu | string | `"300m"` |  |
| activator.containers.domainmapping.resources.limits.memory | string | `"400Mi"` |  |
| activator.containers.domainmapping.resources.requests.cpu | string | `"30m"` |  |
| activator.containers.domainmapping.resources.requests.memory | string | `"40Mi"` |  |
| activator.containers.domainmappingwebhook.resources.limits.cpu | string | `"500m"` |  |
| activator.containers.domainmappingwebhook.resources.limits.memory | string | `"500Mi"` |  |
| activator.containers.domainmappingwebhook.resources.requests.cpu | string | `"100m"` |  |
| activator.containers.domainmappingwebhook.resources.requests.memory | string | `"100Mi"` |  |
| activator.containers.webhook.resources.limits.cpu | string | `"500m"` |  |
| activator.containers.webhook.resources.limits.memory | string | `"500Mi"` |  |
| activator.containers.webhook.resources.requests.cpu | string | `"100m"` |  |
| activator.containers.webhook.resources.requests.memory | string | `"100Mi"` |  |
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
| envoy.containers.envoy.resources.limits.cpu | string | `"500m"` |  |
| envoy.containers.envoy.resources.limits.memory | string | `"500Mi"` |  |
| envoy.containers.envoy.resources.requests.cpu | string | `"200m"` |  |
| envoy.containers.envoy.resources.requests.memory | string | `"200Mi"` |  |
| envoy.containers.shutdown-manager.resources.limits.cpu | string | `"400m"` |  |
| envoy.containers.shutdown-manager.resources.limits.memory | string | `"400Mi"` |  |
| envoy.containers.shutdown-manager.resources.requests.cpu | string | `"40m"` |  |
| envoy.containers.shutdown-manager.resources.requests.memory | string | `"40Mi"` |  |
| http_proxy | string | `""` | HTTP proxy information for knative |
| https_proxy | string | `""` | HTTPS proxy information for knative |
| netcontourcontroller.containers.controller.resources.limits.cpu | string | `"400m"` |  |
| netcontourcontroller.containers.controller.resources.limits.memory | string | `"400Mi"` |  |
| netcontourcontroller.containers.controller.resources.requests.cpu | string | `"40m"` |  |
| netcontourcontroller.containers.controller.resources.requests.memory | string | `"400Mi"` |  |
| no_proxy | string | `"localhost,127.0.0.1,10.0.0.0/8,.svc,.cluster.local"` | No proxy information for knative |
| replicas | int | `1` | Replicas for knative components |

