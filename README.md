# endgit-core

## Requirements

- Go 1.27+
- Docker (for local Postgres)

## Local database

Start db with Docker Compose:

```bash
docker compose up -d
```

Default database credentials are in [docker-compose.yml](docker-compose.yml):

- host: `localhost`
- port: `5432`
- database: `endgit`
- user: `endgit`
- password: `endgit`

Stop the database container:

```bash
docker compose down
```

Stop and clear data volume:

```bash
docker compose down -v
```

## Run

```bash
go mod download
go run .
```

Set a custom port with `PORT`:

```bash
PORT=3000 go run .
```

## Database environment variables

Copy [.env.example](.env.example) to `.env`:

```bash
cp .env.example .env
```

Supported variables:

- `DATABASE_URL` (preferred, full DSN)
- `DB_DRIVER` (default: `postgres`)
- `DB_HOST` (default: `localhost`)
- `DB_PORT` (default: `5432`)
- `DB_USER` (default: `endgit`)
- `DB_PASSWORD` (default: `endgit`)
- `DB_NAME` (default: `endgit`)
- `DB_SSLMODE` (default: `disable`)
- `DB_AUTO_MIGRATE` (default: `true`)

## Seed placeholder data

Run the seed script:

```bash
./scripts/seed-db.sh
```

This creates one maintainer user and a few sample plugins.

## API docs

Swagger UI is always on at:

- `http://localhost:8080/docs`

Regenerate docs after endpoint changes:

```bash
go generate ./...
```