# Run the app locally using Docker

## Run database container
## 1) Run the database container and migration
docker compose up -d
## 2) Check the running container
docker container ls
## 3) Build the app image `charge-points-distribution`
docker build -t charge-points-distribution .
## 4) Check the built image
docker images
## 5) Run the app
docker run -it -p 8080:8080 charge-points-distribution
## 6) Check that the app is running
docker ps
## 7) Open localhost with port 8080

## Run unit tests
go test ./pkg/service -v