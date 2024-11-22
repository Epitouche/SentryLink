# POC React native

## React native

React Native is an open-source framework created by Facebook that enables developers to build mobile applications for iOS and Android using JavaScript and React, a popular library for building user interfaces.

Instead of rendering web components like React for the browser, React Native renders native components directly on the mobile platform. This approach allows developers to create apps with a native look and feel, leveraging platform-specific components, all while using a shared codebase for multiple platforms.

### Documentation

[React Native Official docs](https://reactnative.dev/docs/environment-setup)

### Install on linux

#### Install Node.js via your Linux package manager:

```bash
sudo apt install -y nodejs npm
```

#### Install openJDK
```bash
sudo apt install -y openjdk-11-jdk
```

#### Install Android Studio

1. Download the Android Studio package from [Android Studio](https://developer.android.com/studio)

2. Extract the tar.gz file:
```bash
tar -xvzf android-studio-*.tar.gz
sudo mv android-studio /opt/
```

3. Launch Android Studio:
```bash
/opt/android-studio/bin/studio.sh
```

#### Install Required SDK Components

 - Open Android Studio and complete the setup wizard.
 - Install the following SDK components via SDK Manager:
     - Android SDK
     - Android SDK Platform-Tools
     - Android Emulator
     - One or more system images for the emulator

#### Configure Environment variables

1. Edit your shell configuration file (~/.bashrc or ~/.zshrc):

```bash
export ANDROID_HOME=$HOME/Android/Sdk
export PATH=$PATH:$ANDROID_HOME/emulator:$ANDROID_HOME/tools:$ANDROID_HOME/tools/bin:$ANDROID_HOME/platform-tools
```

2. Apply the changes:

```bash
source ~/.bashrc
```

#### Install React Native CLI

```bash
npm install -g react-native-cli
```

### Project build and run

#### Build

```bash
npx react-native init MyReactNativeApp
```

#### Run

```bash
npm run
```
#### Build apk

```bash
cd android
./gradlew :app:assembleRelease
```