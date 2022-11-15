# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine

WORKDIR path_manager

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
COPY conf ./conf
COPY controller ./controller
COPY model ./model

RUN go mod download

COPY *.go ./

RUN go build pathManager

EXPOSE 8081

CMD [ "./pathManager" ]