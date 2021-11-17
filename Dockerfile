# Multi-stage build setup (https://docs.docker.com/develop/develop-images/multistage-build/)
ARG TARGET=registry

# Stage 1 (to create a "build" image, ~850MB)
FROM golang:1.16 AS builder
RUN go version

WORKDIR /moar/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o moar cli/main.go

# Stage 2 (to create a downsized "container executable", ~7MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.7
#   RUN apk --no-cache add ca-certificates
#
#FROM scratch as moar-cli
FROM alpine:3.7 as moar-cli
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /moar/moar .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# enter cli and keep-alive
ENTRYPOINT ["tail", "-f", "/dev/null"]

FROM alpine:3.7 as moar-registry
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /moar/moar .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
EXPOSE 8000
CMD ["./moar", "up", "-d"]

# build requested image
FROM moar-${TARGET}