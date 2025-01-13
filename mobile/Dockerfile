FROM reactnativecommunity/react-native-android AS node-install

WORKDIR /app/mobile/

RUN apt-get update

RUN apt-get install -y nodejs npm

COPY ./mobile/package.json ./mobile/package-lock.json /app/mobile/

RUN npm install --legacy-peer-deps


FROM node-install AS java-install

WORKDIR /app

RUN apt-get install -y default-jdk


FROM java-install AS build

WORKDIR /app

COPY ./mobile/ /app/mobile/
COPY ./.env /app/.env

WORKDIR /app/mobile/android

RUN ./gradlew clean :app:assembleRelease --stacktrace


FROM alpine:latest AS build-apk

WORKDIR /app/public/apk

COPY --from=build /app/mobile/android/app/build/outputs/apk/release/app-release.apk /app/public/apk/client.apk

CMD ["echo", "Mobile build complete"]


FROM java-install AS run

WORKDIR /app/mobile

COPY ./mobile /app/mobile

CMD ["npm", "start"]