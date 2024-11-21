FROM reactnativecommunity/react-native-android AS node-install

WORKDIR /app

COPY package.json /app

RUN apt-get update 

RUN apt-get install -y nodejs npm

RUN npm install


FROM node-install AS java-install

WORKDIR /app

RUN apt-get install -y default-jdk


FROM java-install AS build

WORKDIR /app

COPY . /app

WORKDIR /app/android

CMD ["./gradlew", ":app:assembleRelease"]