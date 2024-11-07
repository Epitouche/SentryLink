# SentryLink

## Table of Contents

- [SentryLink](#sentrylink)
  - [Table of Contents](#table-of-contents)
  - [General Info](#general-info)
  - [Technologies](#technologies)
    - [Application](#application)
      - [Application Frontend](#application-frontend)
      - [Application Backend](#application-backend)
      - [Application Database](#application-database)
    - [Monitoring](#monitoring)
      - [Monitoring Frontend](#monitoring-frontend)
      - [Monitoring Backend](#monitoring-backend)
      - [Monitoring Database](#monitoring-database)
  - [Installation](#installation)
    - [Dev](#dev)
      - [Dev Frontend](#dev-frontend)
      - [Dev Backend](#dev-backend)
    - [Production](#production)
  - [Overview](#overview)

## General Info

SentryLink is an web app that detect dead link of a web site.

## Technologies

### Application

#### Application Frontend

- Nuxt
  - TailwindCSS

#### Application Backend

- Golang
  - Gin Framework

#### Application Database

- Postgres

### Monitoring

#### Monitoring Frontend

#### Monitoring Backend

#### Monitoring Database

- PGAdmin

## Installation

### Dev

Install [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) in VSCode

#### Dev Frontend

open `backend` dev container with : `Dev Containers: Rebuild and Reopen in Container` -> `backend`
switch to `frontend` container with : `Dev Containers: Switch Container` -> `frontend`
start dev to frontend

#### Dev Backend

open `frontend` dev container with : `Dev Containers: Rebuild and Reopen in Container` -> `frontend`
switch to `backend` container with : `Dev Containers: Switch Container` -> `backend`
start dev to backend

### Production

```bash
docker compose up --build -d
```

## Overview

image
