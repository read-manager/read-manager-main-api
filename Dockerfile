FROM golang:1.21.6-alpine3.19 as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

FROM alpine:3.16
WORKDIR /
COPY --from=build /app/api ./
EXPOSE 3000
CMD ["./api"]
