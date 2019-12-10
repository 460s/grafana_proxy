FROM golang:alpine AS builder

WORKDIR /
COPY . .
RUN go build -o proxy

FROM alpine
COPY --from=builder /proxy /

EXPOSE 4000
ENTRYPOINT ["/proxy"]