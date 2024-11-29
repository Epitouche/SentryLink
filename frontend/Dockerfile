FROM node:20.12.2-alpine AS build

WORKDIR /app

COPY frontend/package.json ./
COPY frontend/package-lock.json ./

RUN npm install

FROM build AS run

COPY frontend/ ./

EXPOSE 8081

CMD ["npx", "nuxt", "dev"]