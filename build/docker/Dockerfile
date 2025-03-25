FROM golang:alpine AS builder

ENV GOPROXY=https://goproxy.io,direct

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -v -a -ldflags "-w -s" -o ./build/p3o ./cmd/...

FROM alpine:latest

ENV TZ=Asia/Tehran

WORKDIR /app/

COPY --from=builder /app/build/p3o .
#COPY ./migrations/ ./migrations/
COPY ./assets/ ./assets/
COPY ./deploy/ ./deploy/

RUN chmod +x /app/p3o

CMD ["/app/p3o"]
