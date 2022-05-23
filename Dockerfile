# syntax=docker/dockerfile:1
FROM golang:1.17-buster

# This is to initialize an seperate folder in docker container
WORKDIR /app

# move dependencies to the /app folder
COPY go.mod ./
COPY go.sum ./

# After moving run go download to the /app folder
RUN go mod download

# move the code now
COPY . ./

# we are basically outputting a binary docker-gs-ping with app logic
RUN go build -o /crane

EXPOSE 8080

# Execute the binary
CMD [ "/crane" ]