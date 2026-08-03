FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /bin/web ./cmd/web

# Templates, CSS, JS and images are all embedded by //go:embed, so the final
# image needs nothing but the binary.
FROM alpine:3.23
COPY --from=builder /bin/web /bin/web
EXPOSE 4000
CMD ["/bin/web"]
