FROM golang:1.17-buster as builder

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o application -v cmd/tasks.go

FROM alpine:3.15.4
WORKDIR /app
COPY --from=builder /app/application /app/application
COPY config/config.yaml ./config/
EXPOSE 3000
CMD ["/app/application"]