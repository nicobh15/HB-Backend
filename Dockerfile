#Build Stage
FROM golang:1.21-alpine3.19 as builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/app/main.go
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz


#Run Stage
FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env . 
COPY start.sh .
COPY wait-for.sh .
COPY internal/db/migration ./migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]