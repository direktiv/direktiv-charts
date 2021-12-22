# pgo

Installer for PGO, the open source Postgres Operator from Crunchy Data

![Version: 0.2.3](https://img.shields.io/badge/Version-0.2.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 5.0.4](https://img.shields.io/badge/AppVersion-5.0.4-informational?style=flat-square)

## Additional Information

Installer for PGO, the open source Postgres Operator from Crunchy Data.
This chart installs <a href="https://www.crunchydata.com/">CrunchyData's postgres controller</a>.

## Installing the Chart

To install the chart with the name `pgo`:

```console
$ helm repo add direktiv https://chart.direktiv.io
$ helm install pgo direktiv/pgo
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| debug | bool | `true` |  Defaults to the value below. |
| image.image | string | `"registry.developers.crunchydata.com/crunchydata/postgres-operator:ubi8-5.0.4-0"` |  |
| relatedImages."postgres_13_gis_3.1".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:centos8-13.5-3.1-0"` |  |
| relatedImages."postgres_14_gis_3.1".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:centos8-14.1-3.1-0"` |  |
| relatedImages.pgbackrest.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-pgbackrest:centos8-2.36-0"` |  |
| relatedImages.pgbouncer.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-pgbouncer:centos8-1.16-0"` |  |
| relatedImages.pgexporter.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-exporter:ubi8-5.0.4-0"` |  |
| relatedImages.postgres_13.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres:centos8-13.5-0"` |  |
| relatedImages.postgres_14.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres:centos8-14.1-0"` |  |
| singleNamespace | bool | `false` |  false, PGO will watch for Postgres clusters in all namesapces Setting to "true" will instruct PGO to only watch for Postgres clusters in the namespace that it is installed in. Defaults to the value below. |

