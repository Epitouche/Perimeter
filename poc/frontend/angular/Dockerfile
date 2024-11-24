FROM node:20.12.2-alpine AS build

WORKDIR /app

COPY frontend/angular/project/package.json ./

RUN npm install

FROM build AS run

COPY frontend/angular/project/ ./

EXPOSE 4200

CMD ["npm", "start"]