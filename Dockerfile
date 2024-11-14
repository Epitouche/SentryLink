# pgAdmin
FROM dpage/pgadmin4:latest as pgadmin

USER root

RUN mkdir -p /var/lib/pgadmin

COPY ./pgadmin/pgadmin4.db /var/lib/pgadmin/pgadmin4.db

COPY ./pgadmin/config_local.py /pgadmin4/config_local.py

RUN chown -R 5050:5050 /var/lib/pgadmin

USER pgadmin

# Backend
FROM --platform=$BUILDPLATFORM golang:1.21.0 AS buildBackend
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=backend/go.sum,target=go.sum \
    --mount=type=bind,source=backend/go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=backend/,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .

FROM alpine:latest AS final

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

# Copy the executable from the "buildBackend" stage.
COPY --from=buildBackend /bin/server /bin/

ENV GIN_MODE=release

# Expose the port that the application listens on.
EXPOSE 8080

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]
