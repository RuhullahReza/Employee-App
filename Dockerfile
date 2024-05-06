FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/bin

FROM alpine

COPY --from=builder /app/bin /app/bin
COPY env.yaml .

RUN chmod 644 env.yaml

ENTRYPOINT ["/app/bin"]