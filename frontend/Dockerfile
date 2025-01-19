FROM node:22-alpine AS dependencies

WORKDIR /app

COPY ./package.json ./
COPY ./package-lock.json ./

ENV NODE_ENV=production

RUN npm install --production

FROM dependencies AS build

COPY ./ ./

RUN npm run build

FROM build AS run

EXPOSE 8081

CMD ["sh", "start.sh"]


## MOBILE DOCUMENTATION
FROM node:22-alpine AS typedoc-dependencies

WORKDIR /app

COPY ./package.json ./package-lock.json /app/

RUN npm install

COPY . .

RUN npx typedoc --out docs --entryPointStrategy expand

FROM alpine:3.21 AS documentation

WORKDIR /app

# Install http-server in the final stage
RUN apk add --no-cache nodejs npm \
    && npm install -g http-server@latest

COPY --from=typedoc-dependencies /app/docs .

EXPOSE 6061

CMD ["http-server", "-p", "6061"]