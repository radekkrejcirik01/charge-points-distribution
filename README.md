# Current distribution REST API

App for current distribution in groups of charge points

## Run requirements
- Docker
- Database management tool e.g. TablePlus (optional)

## Run the app step by step
`./docs/run.md`

## Run test locally

```shell
go test ./pkg/service -v
```

## REST API endpoints documentation

The documenation of each endpoint can be found in `./docs/api/openapi.spec.yml`
with success and error responses