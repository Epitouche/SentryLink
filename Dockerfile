# pgAdmin
FROM dpage/pgadmin4:latest AS pgadmin

USER root

RUN mkdir -p /var/lib/pgadmin

COPY ./pgadmin/pgadmin4.db /var/lib/pgadmin/pgadmin4.db

COPY ./pgadmin/config_local.py /pgadmin4/config_local.py

RUN chown -R 5050:5050 /var/lib/pgadmin

USER pgadmin

# Backend
FROM --platform=$BUILDPLATFORM golang:1.22 AS build-backend
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=backend/go.sum,target=go.sum \
    --mount=type=bind,source=backend/go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=backend/,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .

FROM alpine:latest AS final-backend

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

COPY --from=build-backend /bin/server /bin/

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]

# Frontend
FROM node:20.12.2-alpine AS final-frontend

WORKDIR /app

COPY frontend/ .

RUN npm install

EXPOSE 3000

CMD ["npx",  "nuxt",  "dev"]

# Mobile

FROM reactnativecommunity/react-native-android AS node-install

WORKDIR /app

COPY mobile/package.json /app

RUN apt-get update 

RUN apt-get install -y nodejs npm

RUN npm install


FROM node-install AS java-install

WORKDIR /app

RUN apt-get install -y default-jdk


FROM java-install AS build

WORKDIR /app

COPY mobile/ /app

WORKDIR /app/android

RUN ./gradlew :app:assembleRelease

FROM alpine:latest AS final-mobile

WORKDIR /app/public/apk

COPY --from=build /app/android/app/build/outputs/apk/release/app-release.apk /app/public/apk/client.apk

CMD ["echo", "Mobile build complete"]