FROM node:20-alpine AS dependencies

WORKDIR /app

COPY ./package.json ./
COPY ./package-lock.json ./

RUN npm install

FROM dependencies AS build

COPY ./ ./

RUN npm run build

FROM node:20-alpine AS run

WORKDIR /app

EXPOSE 8081

COPY --from=build /app/.output ./.output

CMD ["node", ".output/server/index.mjs"]