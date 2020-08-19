FROM golang:1.14-alpine AS builder
WORKDIR /app
COPY go.mod  go.sum /
RUN go mod download
COPY .. ..
RUN go build -o main cmd/*
FROM alpine
WORKDIR /app/
COPY --from=builder /build/main .
EXPOSE 8080
CMD [ "./main" ]
