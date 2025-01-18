FROM reactnativecommunity/react-native-android AS node-install

WORKDIR /app

COPY ./package.json ./package-lock.json /app/

RUN apt-get update

RUN apt-get install -y nodejs npm

RUN npm install --legacy-peer-deps


FROM node-install AS java-install

WORKDIR /app

RUN apt-get install -y default-jdk


FROM java-install AS build

WORKDIR /app

# Copy only files needed for Gradle to leverage caching effectively
COPY ./android/gradlew /app/android/gradlew
COPY ./android/gradle /app/android/gradle
COPY ./android/gradle.properties /app/android/gradle.properties
COPY ./android/settings.gradle /app/android/settings.gradle
COPY ./android/app/build.gradle /app/android/app/build.gradle
COPY ./android/build.gradle /app/android/build.gradle

# Run Gradle dependency resolution step to leverage caching
WORKDIR /app/android
RUN ./gradlew dependencies

COPY ./ /app

WORKDIR /app/android

RUN ./gradlew clean :app:assembleRelease --stacktrace


FROM alpine:latest AS build-apk

WORKDIR /app/public/apk

COPY --from=build /app/android/app/build/outputs/apk/release/app-release.apk /app/public/apk/client.apk

CMD ["echo", "Mobile build complete"]


FROM java-install AS run

WORKDIR /app

COPY ./ /app

CMD ["npm", "start"]