# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine

WORKDIR api_gateway

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
COPY conf ./conf
COPY controllers ./controllers
COPY model ./model
COPY routers ./routers
COPY grpc ./grpc
COPY proto ./proto

RUN go mod download

COPY *.go ./

RUN go build api_gateway

EXPOSE 5000

CMD [ "./api_gateway" ]