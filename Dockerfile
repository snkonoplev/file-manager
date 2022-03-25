FROM alpine:3.15 AS base
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.34-r0/glibc-2.34-r0.apk
RUN apk add glibc-2.34-r0.apk
WORKDIR /app
EXPOSE 8080

FROM golang:1.18 AS build_base
WORKDIR /app
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download
COPY src/. .
RUN CGO_ENABLED=0 go test ./...
RUN go build -o service

FROM base AS final
WORKDIR /app
COPY --from=build_base /app/service .
COPY --from=build_base /app/appsettings.yml .
COPY --from=build_base /app/migrations/ ./migrations/

VOLUME ["/data"]

CMD ["./service"]