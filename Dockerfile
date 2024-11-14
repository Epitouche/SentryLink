# pgAdmin
FROM dpage/pgadmin4:latest as pgadmin

USER root

RUN mkdir -p /var/lib/pgadmin

COPY ./pgadmin/pgadmin4.db /var/lib/pgadmin/pgadmin4.db

COPY ./pgadmin/config_local.py /pgadmin4/config_local.py

RUN chown -R 5050:5050 /var/lib/pgadmin

USER pgadmin

# Backend
FROM --platform=$BUILDPLATFORM golang:1.22 AS buildBackend
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=backend/go.sum,target=go.sum \
    --mount=type=bind,source=backend/go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=backend/,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .

FROM alpine:latest AS finalBackend

WORKDIR /bin

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

COPY --from=buildBackend /bin/server /bin/

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]

# Frontend
FROM node:20.12.2-alpine as baseFrontend

WORKDIR /usr/src/app

COPY frontend/ .

RUN npm install

EXPOSE 3000

CMD ["npx",  "nuxt",  "dev"]