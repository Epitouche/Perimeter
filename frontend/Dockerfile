FROM node:20.12.2-alpine AS build

WORKDIR /app

COPY ./package.json ./
COPY ./package-lock.json ./

RUN npm install

FROM build AS run

COPY ./ ./

EXPOSE 8081

CMD ["npx", "nuxt", "dev"]