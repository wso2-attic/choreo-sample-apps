# Choreo Samples

## Sample Go API

#### Use the following build config when creating this component in Choreo:

- Dockerfile: `go/rest-api/Dockerfile`
- Docker context: `go/rest-api/`
- Port: `8080` (or set env var `PORT`)
- OpenAPI filepath: `go/rest-api/docs/swagger.json`

#### Go build & run

```shell
go build main.go && go run main.go
```

#### Generate API definitions

Generated using Go annotations https://github.com/swaggo/swag

```shell
swag init
```

#### Load initial data ( optional )

1. Set env var in Choreo DevOps portal `INIT_DATA_PATH=<some_path>`
2. Mount the file contents of `configs/initial_data.json` in the path specified in step 1.
