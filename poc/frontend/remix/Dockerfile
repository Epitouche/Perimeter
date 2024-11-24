FROM node:20.12.2-alpine AS build-manual

WORKDIR /app

COPY frontend/remix/project/manual/package.json ./

RUN npm install

FROM build-manual AS run-manual

COPY frontend/remix/project/manual/ ./

EXPOSE 3000

CMD ["npm", "run", "dev"]


FROM node:20.12.2-alpine AS build-auto

WORKDIR /app

COPY frontend/remix/project/automatic/package.json ./

RUN npm install

FROM build-auto AS run-auto

COPY frontend/remix/project/automatic/ ./

EXPOSE 5173

CMD ["npm", "run", "dev"]