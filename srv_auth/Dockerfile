FROM golang:alpine AS builder

WORKDIR /build
ADD go.mod .
COPY . .
COPY /internal/config /tempEnv
RUN go build -o . cmd/main.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main
COPY --from=builder /tempEnv /build/internal/config
CMD ["./main"]