FROM golang:alpine AS builder

RUN adduser -D -g '' app-user

WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o proxy

FROM scratch

COPY --from=builder /proxy /
COPY --from=builder /etc/passwd /etc/passwd

USER app-user

EXPOSE 4000
ENTRYPOINT ["/proxy"]