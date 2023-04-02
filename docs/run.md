## Run the app locally using Docker

### Run database container
### 1) Run the database container and migration
```shell
docker compose up -d
```
### 2) Check the running container
```shell
docker container ls
```
### 3) Build the app image `charge-points-distribution`
```shell
docker build -t charge-points-distribution .
```
### 4) Check the built image
```shell
docker images
```
### 5) Run the app
```shell
docker run -it -p 8080:8080 charge-points-distribution
```
### 6) Check that the app is running
```shell
docker ps
```
### 7) Open localhost with port 8080
```shell
localhost:8080
```

## Run unit tests
```shell
go test ./pkg/service -v
```