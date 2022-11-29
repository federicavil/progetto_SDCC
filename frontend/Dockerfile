# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine

WORKDIR frontend

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY model ./model
COPY routers ./routers
COPY static ./static
COPY views ./views
COPY conf ./conf
COPY controllers ./controllers



COPY *.go ./

RUN go build frontend

EXPOSE 8080

CMD [ "./frontend" ]
