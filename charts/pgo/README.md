# pgo

![Version: 5.3.0](https://img.shields.io/badge/Version-5.3.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 5.3.0](https://img.shields.io/badge/AppVersion-5.3.0-informational?style=flat-square)

Installer for PGO, the open source Postgres Operator from Crunchy Data

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| controllerImages.cluster | string | `"registry.developers.crunchydata.com/crunchydata/postgres-operator:ubi8-5.3.0-0"` |  |
| controllerImages.upgrade | string | `"registry.developers.crunchydata.com/crunchydata/postgres-operator-upgrade:ubi8-5.3.0-0"` |  |
| debug | bool | `true` |  |
| imagePullSecretNames | list | `[]` |  |
| relatedImages."postgres_13_gis_3.0".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-13.9-3.0-2"` |  |
| relatedImages."postgres_13_gis_3.1".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-13.9-3.1-2"` |  |
| relatedImages."postgres_14_gis_3.1".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-14.6-3.1-2"` |  |
| relatedImages."postgres_14_gis_3.2".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-14.6-3.2-2"` |  |
| relatedImages."postgres_14_gis_3.3".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-14.6-3.3-2"` |  |
| relatedImages."postgres_15_gis_3.3".image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-gis:ubi8-15.1-3.3-0"` |  |
| relatedImages.pgadmin.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-pgadmin4:ubi8-4.30-8"` |  |
| relatedImages.pgbackrest.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-pgbackrest:ubi8-2.41-2"` |  |
| relatedImages.pgbouncer.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-pgbouncer:ubi8-1.17-5"` |  |
| relatedImages.pgexporter.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres-exporter:ubi8-5.3.0-0"` |  |
| relatedImages.pgupgrade.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-upgrade:ubi8-5.3.0-0"` |  |
| relatedImages.postgres_13.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres:ubi8-13.9-2"` |  |
| relatedImages.postgres_14.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres:ubi8-14.6-2"` |  |
| relatedImages.postgres_15.image | string | `"registry.developers.crunchydata.com/crunchydata/crunchy-postgres:ubi8-15.1-0"` |  |
| resources.controller | object | `{}` |  |
| resources.upgrade | object | `{}` |  |
| singleNamespace | bool | `false` |  |

