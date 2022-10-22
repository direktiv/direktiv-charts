# pgo

Installer for PGO, the open source Postgres Operator from Crunchy Data

![Version: 0.6.0](https://img.shields.io/badge/Version-0.6.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 5.2.0](https://img.shields.io/badge/AppVersion-5.2.0-informational?style=flat-square)

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
| controllerImages.cluster | string | `"registry.developers.crunchydata.com/crunchydata/postgres-operator:ubi8-5.2.0-0"` |  |
| controllerImages.upgrade | string | `"registry.developers.crunchydata.com/crunchydata/postgres-operator-upgrade:ubi8-5.2.0-0"` |  |
| debug | bool | `false` |  |
| imagePullSecretNames | list | `[]` |  |
| relatedImages."postgres_13_gis_3.0".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-13.8-3.0-1"` |  |
| relatedImages."postgres_13_gis_3.1".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-13.8-3.1-1"` |  |
| relatedImages."postgres_14_gis_3.1".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-14.5-3.1-1"` |  |
| relatedImages."postgres_14_gis_3.2".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-14.5-3.2-1"` |  |
| relatedImages.pgadmin.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-pgadmin4:ubi8-4.30-4"` |  |
| relatedImages.pgbackrest.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-pgbackrest:ubi8-2.40-1"` |  |
| relatedImages.pgbouncer.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-pgbouncer:ubi8-1.17-1"` |  |
| relatedImages.pgexporter.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-exporter:ubi8-5.2.0-0"` |  |
| relatedImages.pgupgrade.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-upgrade:ubi8-5.2.0-0"` |  |
| relatedImages.postgres_13.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres:ubi8-13.8-1"` |  |
| relatedImages.postgres_14.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres:ubi8-14.5-1"` |  |
| singleNamespace | bool | `true` |  |

