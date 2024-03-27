# Multi-stage build setup (https://docs.docker.com/develop/develop-images/multistage-build/)
ARG TARGET=registry

FROM golang:1.21 as base

# Stage 1 (to create a "build" image, ~850MB)
FROM base AS builder
RUN go version

WORKDIR /moar/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o moar cli/main.go

FROM base as dev

# Install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Run the air command in the directory where our code will live
WORKDIR /moar/

CMD ["air"]

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
