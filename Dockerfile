FROM golang:latest AS build
ENV GO111MODULE=on
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main
FROM scratch AS release
ENV GIN_MODE=release
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/ /app/
WORKDIR /app
CMD ["/app/main"]
