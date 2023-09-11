# Cronupper

Schedule cron jobs to backup and upload automatically

## Presets and uploaders

Presets

- Custom command
- Docker volumes
- MongoDB
- Postgres

Uploaders

- S3

## Usage

Docker compose

```yaml
services:
  cronupper:
    image: "mducoli/cronupper:latest"
    restart: unless-stopped
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock:ro"
      - "./cronupper-config.yml:/etc/cronupper/config.yml"
```

### Configuration

Example

```yaml
jobs:
  postgres:
    cron: "0 3 * * *"
    preset: postgres
    config:
      container: "postgres"
      postgresuser: "postgresuser"
    upload:
      to: s3_1
      filename: postgres-$(date +%Y%m%dT%H%M).sql
      config:
        bucket: "backups"
        prefix: "postgres/"
  pgadmin:
    cron: "0 3 * * *"
    preset: docker-volume
    config:
      volume: "pgadmin_data"
    upload:
      to: s3_1
      filename: pgadmin-$(date +%Y%m%dT%H%M).tgz
      config:
        bucket: "backups"
        prefix: "pgadmin/"

uploaders:
  s3_1:
    provider: s3
    config:
      endpoint: "s3.eu-west-3.amazonaws.com"
      access_key: ${S3_ACCESS_KEY}
      secret_key: ${S3_SECRET_KEY}
```
