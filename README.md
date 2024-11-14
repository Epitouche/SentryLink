# SentryLink

## Techno

Nuxt

Golang

## Install / Run

put the database password in db/password.txt:

```bash
echo '[PASSWORD]' >> db/password.txt
```
****
### Dev

#### Frontend

open `backend` dev container
switch to `frontend` container
start dev to Frontend

#### Backend

open `frontend` dev container
switch to `backend` container
start dev to backend

### Production

```bash
docker compose up --build -d
```
