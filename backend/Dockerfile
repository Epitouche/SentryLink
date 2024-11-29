FROM golang:1.23-alpine AS dependencies

WORKDIR /app

COPY backend/go.mod .

COPY backend/go.sum .

RUN go mod download


FROM dependencies AS build

COPY backend/ .

RUN go build -o /app


FROM dependencies AS test

COPY backend/ .

CMD ["go", "test", "-v"]

FROM alpine:3.15 AS run

WORKDIR /app

COPY --from=build /app/SentryLink .

EXPOSE 8080

CMD ["/app/SentryLink"]