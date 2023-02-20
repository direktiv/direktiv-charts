# direktiv

direktiv helm chart

![Version: 0.1.16](https://img.shields.io/badge/Version-0.1.16-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v0.7.2](https://img.shields.io/badge/AppVersion-v0.7.2-informational?style=flat-square)

## Additional Information

This chart installs direktiv.

### Changes in 0.1.16

*Direktiv version upgrade*

### Changes in 0.1.15

* Deprecated unused function configuration items
* Added disk size to small, medium, large function settings
* Added permission to watch configmaps to function deployment
* Updated dependencies (Prometheus, Nginx)

### Changes in 0.1.14

* Added helm labels to function namespace
* Updated prometheus, nginx dependencies

### Changes in 0.1.13

* Added function namespace generation automatically
* Added update strategy for deployments
* Ingresses are optional, default true

### Changes in 0.1.12

*Version upgrade*

### Changes in 0.1.11

*Version upgrade*

### Changes in 0.1.10

*Version upgrade*
*Added API secret*

### Changes in 0.1.9
*Version fix*

### Changes in 0.1.8
*Version upgrade*

### Changes in 0.1.7
*Version upgrade*

### Changes in 0.1.6
*Flow filesystem is writable for git integration*

### Changes in 0.1.3

*creating of service accounts is optional*
*added `additional` for additional attribuites for db connections*
*make the cpu/mem limits for knative containers configurable*
*multiple replicas have now requiredDuringSchedulingIgnoredDuringExecution podAntiAffinity*

### Changes in 0.1.2

*Removed unnecessary environment variables in UI deployment*
*Fixed typo in opentelemetry config*

## Installing the Chart

To install the chart with the release name `direktiv`:

```console
$ kubectl create ns direktiv-services-direktiv
$ helm repo add direktiv https://chart.direktiv.io
$ helm install direktiv direktiv/direktiv
```

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://kubernetes.github.io/ingress-nginx | ingress-nginx | 4.4.0 |
| https://prometheus-community.github.io/helm-charts | prometheus | 15.18.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| api.affinity | object | `{}` |  |
| api.extraContainers | list | `[]` | extra container in api pod |
| api.extraContainers | list | `[]` |  |
| api.extraVolumeMounts | string | `nil` | extra volume mounts in api pod |
| api.extraVolumes | string | `nil` | extra volumes in api pod |
| api.image | string | `"direktiv/api"` | image for api pod |
| api.replicas | int | `1` |  |
| api.tag | string | `""` | image tag for api pod |
| apikey | string | `"none"` | enabled api key for API authentication with the `direktiv-token` header, key set in http-snippet in `ingress-nginx` none |
| database.additional | string | `""` | additional connection attributes, e.g. target_session_attrs |
| database.host | string | `"postgres-postgresql-ha-pgpool.postgres"` | database host |
| database.name | string | `"direktiv"` | database name, has to be created before installation |
| database.password | string | `"direktivdirektiv"` | database password |
| database.port | int | `5432` | database port |
| database.sslmode | string | `"require"` | sslmode for database |
| database.user | string | `"direktiv"` | database user |
| debug | bool | `false` | enable debug across all direktiv components |
| encryptionKey | string | `"01234567890123456789012345678912"` |  |
| eventing | object | `{"enabled":false}` | knative eventing enabled, requires knative setup and configuration |
| flow.affinity | object | `{}` | affinity for flow pods |
| flow.dbimage | string | `"direktiv/flow-dbinit"` | image for db update init container |
| flow.extraContainers | list | `[]` | extra container in flow pod |
| flow.extraVolumeMounts | string | `nil` | extra volume mounts in flow pod |
| flow.extraVolumes | string | `nil` | extra volumes in flow pod |
| flow.image | string | `"direktiv/flow"` | image for flow pod |
| flow.replicas | int | `1` | number of flow replicas |
| flow.tag | string | `""` | image tag for flow pod |
| fluentbit.extraConfig | string | `""` | postgres for direktiv services Append extra output to fluentbit configuration. There are two log types: application (system), functions (workflows) these can be matched to new outputs. |
| functions.affinity | object | `{}` |  |
| functions.extraContainers | list | `[]` | extra containers for tasks and knative pods |
| functions.extraContainersPod | list | `[]` | extra containers for function controller, e.g. database containers for google cloud or logging |
| functions.extraVolumes | list | `[]` | extra volumes for tasks and knative pods |
| functions.http_proxy | string | `""` | http_proxy injected as environment variable in functions |
| functions.https_proxy | string | `""` | https_proxy injected as environment variable in functions |
| functions.image | string | `"direktiv/functions"` |  |
| functions.ingressClass | string | `"contour.ingress.networking.knative.dev"` |  |
| functions.initPodImage | string | `"direktiv/direktiv-init-pod"` |  |
| functions.limits | object | `{"cpu":{"large":1,"medium":"500m","small":"250m"},"disk":{"large":4096,"medium":1024,"small":256},"memory":{"large":2048,"medium":1024,"small":512}}` | knative service limits  |
| functions.namespace | string | `"direktiv-services-direktiv"` |  |
| functions.netShape | string | `"10M"` | Egress/Ingress network limit for functions if supported by network |
| functions.no_proxy | string | `""` | no_proxy injected as environment variable in functions |
| functions.podCleaner | bool | `true` | Cleaning up tasks, Kubernetes < 1.20 does not clean finished tasks |
| functions.replicas | int | `1` | number of controller replicas |
| functions.runtime | string | `"default"` | runtime to use, e.g. gvisor on GCP |
| functions.sidecar | string | `"direktiv/sidecar"` |  |
| functions.tag | string | `""` |  |
| http_proxy | string | `""` | http proxy settings |
| https_proxy | string | `""` | https proxy settings |
| imagePullSecrets | list | `[]` |  |
| ingress-nginx | object | `{"controller":{"admissionWebhooks":{"patch":{"podAnnotations":{"linkerd.io/inject":"disabled"}}},"config":{"proxy-buffer-size":"16k"},"podAnnotations":{"linkerd.io/inject":"disabled"},"replicaCount":1},"install":true}` | nginx ingress controller configuration |
| ingress.additionalAnnotations | object | `{}` | Additional Annotations |
| ingress.certificate | string | `"none"` | TLS secret |
| ingress.class | string | `"nginx"` | ingress class |
| ingress.enabled | bool | `true` |  |
| ingress.host | string | `""` | host for external services, only required for TLS |
| logging | string | `"json"` | json or console logger |
| networkPolicies.db | string | `"0.0.0.0/0"` | CIDR for database, excempt from policies |
| networkPolicies.enabled | bool | `false` | adds network policies |
| networkPolicies.podCidr | string | `"0.0.0.0/0"` | CIDR for pods, excempt from policies |
| networkPolicies.serviceCidr | string | `"0.0.0.0/0"` | CIDR for services, excempt from policies |
| no_proxy | string | `""` | no proxy proxy settings |
| nodeSelector | object | `{}` |  |
| opentelemetry.address | string | `"localhost:4317"` | opentelemetry address where Direktiv is sending data to |
| opentelemetry.agentconfig | string | `"receivers:\n  otlp:\n    protocols:\n      grpc:\n      http:\nexporters:\n  otlp:\n    endpoint: \"192.168.1.113:14250\"\n    insecure: true\n    sending_queue:\n      num_consumers: 4\n      queue_size: 100\n    retry_on_failure:\n      enabled: true\n  logging:\n    loglevel: debug\nprocessors:\n  batch:\n  memory_limiter:\n    # Same as --mem-ballast-size-mib CLI argument\n    ballast_size_mib: 165\n    # 80% of maximum memory up to 2G\n    limit_mib: 400\n    # 25% of limit up to 2G\n    spike_limit_mib: 100\n    check_interval: 5s\nextensions:\n  zpages: {}\nservice:\n  extensions: [zpages]\n  pipelines:\n    traces:\n      receivers: [otlp]\n      processors: [memory_limiter, batch]\n      exporters: [logging, otlp]\n"` | config for sidecar agent |
| opentelemetry.enabled | bool | `false` | installs opentelemtry agent as sidecar in flow |
| prometheus.alertmanager.enabled | bool | `false` |  |
| prometheus.backendName | string | `"prom-backend-server"` |  |
| prometheus.global.evaluation_interval | string | `"1m"` |  |
| prometheus.global.scrape_interval | string | `"1m"` |  |
| prometheus.install | bool | `true` |  |
| prometheus.kubeStateMetrics.enabled | bool | `false` |  |
| prometheus.nodeExporter.enabled | bool | `false` |  |
| prometheus.pushgateway.enabled | bool | `false` |  |
| prometheus.server.persistentVolume.enabled | bool | `false` |  |
| prometheus.server.retention | string | `"96h"` |  |
| prometheus.serviceAccounts.alertmanager.create | bool | `false` |  |
| prometheus.serviceAccounts.nodeExporter.create | bool | `false` |  |
| prometheus.serviceAccounts.pushgateway.create | bool | `false` |  |
| prometheus.serviceAccounts.server.create | bool | `true` |  |
| pullPolicy | string | `"Always"` |  |
| registry | string | `"docker.io"` |  |
| secrets | object | `{"db":"","extraVolumeMounts":[],"image":"direktiv/secrets","tag":""}` | secrets sidecar in flow pod |
| serviceAccount | object | `{"annotations":{},"create":true,"name":""}` | service account for components. If preconfigured serviceaccounts are used the name ise the base  and two additional service accounts are needed, e.g. service account name is myaccount then another two  acounts are needed: myaccount-functions and myaccount-functions-pod |
| timeout | int | `7200` | max request timeouts in seconds |
| tolerations | list | `[]` |  |
| ui | object | `{"affinity":{},"certificate":"none","extraContainers":[],"image":"direktiv/ui","replicas":1,"tag":""}` | UI configuration |

