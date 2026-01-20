FROM golang:alpine AS builder

RUN apk --no-cache add gcc g++ make git bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify && go mod tidy

COPY . .

RUN GOOS=linux go build -ldflags="-s -w" -o webserver ./cmd/app/main.go

FROM alpine:3.17

RUN apk --no-cache add bash

WORKDIR /app

COPY --from=builder /app/webserver .

EXPOSE 8082

CMD ["./webserver"]