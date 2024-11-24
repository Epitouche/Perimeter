FROM node:20.12.2-alpine AS build

WORKDIR /app

COPY frontend/nuxt/project/package.json ./

RUN npm install

FROM build AS run

COPY frontend/nuxt/project/ ./

EXPOSE 3000

CMD ["npm", "run", "dev", "--"]