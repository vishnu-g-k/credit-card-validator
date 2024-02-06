# syntax=docker/dockerfile:1

FROM golang:1.21.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /credit-card-validator

EXPOSE 8080

CMD ["/credit-card-validator"]