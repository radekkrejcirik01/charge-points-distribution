# Current distribution REST API

This application distributes maximum current of charge group between charge points based on priority of the charge point. The priority is calculated between available charge points that have none connectors with `Charging` status, otherwise 0 current is allocated to the charge point.

## Run requirements
- Docker
- Database management tool e.g. TablePlus (optional)

## Running app step by step
To start the app locally, all we need is to run database container with init values and then build the image of the app and run it using docker. The app is then active on localhost:8080

All steps are described in subfolder `/docs/run.md`

## Run unit tests

```shell
go test ./pkg/service -v
```

## REST API endpoints documentation

The documenation of each endpoint can be found in OpenAPI `/docs/api/openapi.spec.yml`
with success and error responses

## Video presentation
https://drive.google.com/file/d/1XB021zDfmUCC4nACO7jqg79ARtUuegJX/view?usp=share_link
