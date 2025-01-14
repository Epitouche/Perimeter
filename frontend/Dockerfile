FROM node:22-alpine AS dependencies

WORKDIR /app

COPY ./package.json ./
COPY ./package-lock.json ./

RUN npm install

FROM dependencies AS build

COPY ./ ./

RUN npm run build

FROM build AS run

# WORKDIR /app

# COPY --from=build /app/dist ./dist
# COPY --from=build /app/.output ./.output

# ENV PORT=8081

EXPOSE 8081

# CMD [ "npx", "serve", ".public" ]
# CMD [ "npx", "serve", ".output/public" ]
# CMD ["node", "./server/index.mjs"]
CMD ["sh", "start.sh"]