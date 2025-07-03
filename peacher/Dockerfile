FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o peacher ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/peacher .
COPY --from=build /app/.env .env

RUN chmod +x peacher

EXPOSE 8080

CMD ["./peacher"]
