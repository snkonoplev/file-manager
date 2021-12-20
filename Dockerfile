FROM golang:latest AS base
WORKDIR /app
EXPOSE 8080

FROM golang:latest AS build_base
WORKDIR /app
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download
COPY src/. .
RUN go build -o service

FROM base AS final
WORKDIR /app
COPY --from=build_base /app/service .
COPY --from=build_base /app/appsettings.yml .
#COPY --from=build_base /app/data/manager.db ./data/
COPY --from=build_base /app/migrations/ ./migrations/

VOLUME ["/data"]

CMD ["./service"]