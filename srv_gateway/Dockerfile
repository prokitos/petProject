FROM golang:alpine AS builder

WORKDIR /build
ADD go.mod .
COPY . .
COPY /internal/config /tempEnv
COPY /internal/database/migrations /tempMigration
RUN go build -o . cmd/myService/main.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main
COPY --from=builder /tempEnv /build/internal/config
COPY --from=builder /tempMigration /build/internal/database/migrations
CMD ["./main"]