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