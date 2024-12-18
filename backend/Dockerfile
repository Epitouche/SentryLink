FROM golang:1.23-alpine AS dependencies

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download


FROM dependencies AS build

COPY . /app/

RUN go build -o /app


FROM dependencies AS test

COPY backend /app/

CMD ["go", "test", "-v"]

FROM alpine:3.15 AS run

WORKDIR /app

COPY docs /app/docs

COPY --from=build /app/SentryLink .

EXPOSE 8080

CMD ["/app/SentryLink"]
