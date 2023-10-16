# Docker-redis-Dummy-

## Command to run build,push and run the image

### Build the Image

 docker-compose build

 ### Push the image to dockerhub

 docker-compose push

 ### Pull the image from dockerhub

 docker pull {imagename}.This code is given in dockerhub where you have pushed your image

### Create a custom network
docker network create my-network

 ### Start the Redis container and connect it to the custom network
docker run -d --network my-network --name redis redis:latest

 ### Start your application container, also connecting it to the custom network
docker run -d -p 8080:8080 --network my-network navneetshukla/goredis
