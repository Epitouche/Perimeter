FROM node:20-alpine AS dependencies

WORKDIR /app

COPY ./package.json ./
COPY ./package-lock.json ./

RUN npm install

FROM dependencies AS run

COPY ./ ./

RUN npm run build

# FROM node:20-alpine AS run

# WORKDIR /app

EXPOSE 8081

# COPY --from=build /app/.output ./

# CMD ["node", "./server/index.mjs"]
CMD ["npx", "nuxt", "dev"]