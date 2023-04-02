# Current distribution REST API

This application distributes maximum current of charge group between charge points based on charging status of each charge point's connectors. If charge point has none connectors with `Charging` status, 0 current is allocated to the charge point.

## Run requirements
- Docker
- Database management tool e.g. TablePlus (optional)

## Running app step by step
To start the app locally, all we need is to run database container with init values and then build the image of the app and run it using docker. The app is then active on localhost:8080
`./docs/run.md`

## Run test locally

```shell
go test ./pkg/service -v
```

## REST API endpoints documentation

The documenation of each endpoint can be found in `./docs/api/openapi.spec.yml`
with success and error responses