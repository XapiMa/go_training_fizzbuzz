FROM golang:alpine AS build-env
WORKDIR /building
COPY main.go /building
RUN go build -o fizzbuzz_server main.go

FROM alpine
WORKDIR /app
COPY  --from=build-env /building/fizzbuzz_server /app/fizzbuzz_server
EXPOSE 8081
CMD ["./fizzbuzz_server"]
