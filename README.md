# Meetrics

Basic Google Calendar meeting data tracking within a Google organization.

```
go run db_init.go       # Initialize the local database.
go run backfill.go      # Backfill the database with meeting data
```

## Docker

Build and push the docker image. Jenkins pulls the docker image each night and runs the script to populate the database with meeting information.

```bash
docker build -t chasdevs/meetrics .
```

You can run locally from the docker image:

```bash
docker run --rm -e DB_HOST="docker.for.mac.localhost" chasdevs/meetrics
```

Or against prod:
```bash
docker run --rm -e ENV=production chasdevs/meetrics
```