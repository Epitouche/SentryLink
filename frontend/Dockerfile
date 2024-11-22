FROM node:20.12.2-alpine AS final-frontend

WORKDIR /app

COPY frontend/ .

RUN npm install

EXPOSE 3000

CMD ["npx",  "nuxt",  "dev"]