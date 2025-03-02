# Base stage just for basic configuration
FROM golang:1.24-alpine3.20 AS base

ENV DIR /server
WORKDIR $DIR

# Build stage to compile the project
FROM base AS builder

COPY go.mod go.sum *.go ./

RUN go mod download

RUN mkdir bin

RUN go build -o bin/app .

# Production stage where the project runs
FROM scratch AS production

WORKDIR /server

COPY --from=builder /server/bin/app .

EXPOSE 8080

CMD ["./app"]