FROM golang:1.11.0-alpine3.8 AS build-env
WORKDIR /building
COPY main.go /building
RUN go build -o fizzbuzz_server main.go

FROM alpine:3.8
WORKDIR /app
COPY  --from=build-env /building/fizzbuzz_server /app/fizzbuzz_server
EXPOSE 8081
CMD ["./fizzbuzz_server"]
