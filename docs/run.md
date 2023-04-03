## To make the app work we need to run the database and app containers with these simple steps:

### 1) Build and start the containers with init sql script
Details about postgres and app services can be found in `docker-compose.yaml` file
```shell
docker-compose up
```
### 2) Check the running containers
```shell
docker container ls
```
### 3) Connect locally to database management tool e.g. TablePlus (optional)
All tables and records are directly in database after every change.
Use database management tool of your choice and create new connection to
postgres with these values:
```shell
host     = "localhost"
port     = "5432"
user     = "user"
password = "userpassword"
database = "distribution"
```
### 4) Open localhost with port 8080
```shell
localhost:8080
```
### 5) Perform API post and put requests using API platform e.g. Postman (optional)
To successfully call `POST` and `PUT` requests for creating and updating groups, charge points and connectors
we need to open the API platfom tool and select either `POST` or `PUT` method

Example:
```shell
`PUT` localhost:8080/charge-point
```
The request body will contain `id`, `groupId` and `priority` fields

Detailed description of the endpoints can be found in `docs/api/openapi.spec.yml` file

## Run unit tests
```shell
go test ./pkg/service -v
```