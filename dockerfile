FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /bin/ems

LABEL org.opencontainers.image.source="github.com/EcoPowerHub/EMS"
