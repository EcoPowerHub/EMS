FROM golang:1.21 as builder

ARG GH_TOKEN

WORKDIR /app

RUN git config --global url."https://$GH_TOKEN@github.com/".insteadOf "https://github.com/"

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /bin/ems

LABEL org.opencontainers.image.source="github.com/EcoPowerHub/EMS"
