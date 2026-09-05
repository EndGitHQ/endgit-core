# endgit-core

## Requirements

- Go 1.27+

## Run

```bash
go mod download
go run .
```

Server starts on `:8080` by default.

Set a custom port with `PORT`:

```bash
PORT=3000 go run .
```

## API docs

Swagger UI is always on at:

- `http://localhost:8080/docs`

Regenerate docs after endpoint changes:

```bash
go generate ./...
```