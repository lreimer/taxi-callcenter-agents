FROM golang:1.24.3-bookworm AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o taxi-mcp-server -ldflags "-s -w -X main.version=$(date +%Y-%m-%dT%H:%M:%S%z)"

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/taxi-mcp-server /taxi-mcp-server

EXPOSE 8001
ENTRYPOINT [ "/taxi-mcp-server" ]
CMD ["-t", "sse", "-p", "8001"]
