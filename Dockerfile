FROM golang:alpine AS builder

ARG ENV
ARG PORT

WORKDIR /app
COPY . .
RUN apk add --no-cache make
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make build

FROM alpine AS runner
WORKDIR /app
COPY --from=builder /app/build/server .
COPY --from=builder /app/.env .env
EXPOSE $PORT
ENTRYPOINT ["./server"]