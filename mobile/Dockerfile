FROM node:16

WORKDIR /app

RUN npm install -g expo-cli eas-cli && \
    apt-get update && apt-get install -y jq

COPY package*.json ./

RUN npm install

COPY . .

# TODO
RUN git init

ARG EXPO_TOKEN
ENV EXPO_TOKEN=$EXPO_TOKEN

RUN eas whoami

RUN eas build --platform android --non-interactive

RUN BUILD_LIST=$(eas build:list) && \
    BUILD_URL=$(echo "$BUILD_LIST" | grep -o -m 1 'https://expo\.dev/artifacts/.*\.apk') && \
    echo "Build Artifact URL: ||$BUILD_URL||" > toto.txt && \
    curl -o /app/build.apk $BUILD_URL

CMD ["sleep", "infinity"]
