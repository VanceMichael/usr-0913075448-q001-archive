FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /out/archive-service ./cmd/server
FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/archive-service /app/archive-service
EXPOSE 8080
ENTRYPOINT ["/app/archive-service"]
