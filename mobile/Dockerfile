FROM reactnativecommunity/react-native-android AS node-install

WORKDIR /app

COPY mobile/package.json /app

RUN apt-get update 

RUN apt-get install -y nodejs npm

RUN npm install


FROM node-install AS java-install

WORKDIR /app

RUN apt-get install -y default-jdk


FROM java-install AS build

WORKDIR /app

COPY mobile/ /app

WORKDIR /app/android

RUN ./gradlew :app:assembleRelease

FROM alpine:latest AS final-mobile

WORKDIR /app/public/apk

COPY --from=build /app/android/app/build/outputs/apk/release/app-release.apk /app/public/apk/client.apk

CMD ["echo", "Mobile build complete"]