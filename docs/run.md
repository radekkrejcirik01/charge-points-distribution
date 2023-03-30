## Run the app locally using Docker

# 1) Build the image `charge-points-distribution`
docker build -t charge-points-distribution .
# 2) Check the built image
docker images
# 3) Run the app
docker run -it -p 8080:8081 charge-points-distribution
# 4) Check that the app is running
docker ps
# 5) Open localhost with port 8080