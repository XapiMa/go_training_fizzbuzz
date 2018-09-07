
FROM alpine

WORKDIR /app
COPY main_linux_amd64 /app

EXPOSE 8081
CMD ["./main_linux_amd64"]
