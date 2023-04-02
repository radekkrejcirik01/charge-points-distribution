## To make the app work we need to run the database container and build the app image with these simple steps:
## Run database container
### 1) Build and start the container with init sql script
Details about postgres service can be find in `docker-compose.yaml` file
```shell
docker compose up -d
```
### 2) Check the running container
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
name     = "distribution"
```

## Run the application
### 1) Build the `charge-points-distribution` image
```shell
docker build -t charge-points-distribution .
```
### 2) Check the built image
```shell
docker images
```
### 3) Run the app
```shell
docker run -it -p 8080:8080 charge-points-distribution
```
### 4) Check that the app is running
```shell
docker ps
```
### 5) Open localhost with port 8080
```shell
localhost:8080
```

## Run unit tests
```shell
go test ./pkg/service -v
```