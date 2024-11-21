
# Expo Project Setup Guide for Linux

## What is Expo?

[Expo](https://expo.dev/) is an open-source platform for making universal native apps for Android, iOS, and the web with JavaScript and React. It simplifies the process of building, deploying, and testing mobile applications by providing a set of tools and services.

With Expo, you can quickly create a new React Native app and test it without the need for configuring native build tools like Xcode or Android Studio.

## Prerequisites

Before starting, ensure that you have the following installed on your system:

- **Node.js** (version 16 or higher)
- **npm** (Node Package Manager)

You can check if these are installed by running:

```bash
node -v
npm -v
```

If you don't have Node.js installed, you can install it by following the instructions at [Node.js official website](https://nodejs.org/).

### Installing Node.js on Linux (Fedora example)

```bash
sudo dnf install nodejs
```

## Installing Expo CLI

The Expo CLI is the command-line tool for creating and managing Expo projects.

To install Expo CLI globally, run the following command:

```bash
npm install -g expo-cli
```

### Verifying Expo CLI installation

After installation, you can verify that Expo CLI is installed correctly by running:

```bash
expo --version
```

You should see the version number of the Expo CLI.

## Creating a New Expo Project

Once you have Expo CLI installed, you can create a new Expo project by running the following command:

```bash
expo init my-new-project
```

This will prompt you to choose a template. You can select the "blank" template or any other template you prefer.

After the project is created, navigate into your project directory:

```bash
cd my-new-project
```

## Running the Expo Project on Linux

To run your Expo project locally, use the following command:

```bash
expo start
```

This will start a local development server and provide you with a QR code. You can scan this QR code using the Expo Go app (available on both Android and iOS) to view the app on your device.

Alternatively, you can run the project on a web browser using the browser option provided in the terminal.

### Additional commands:

- **To run the project on an Android emulator (if configured):**

```bash
expo start --android
```

- **To run the project on iOS simulator (if configured on macOS):**

```bash
expo start --ios
```

## Building the Expo Project for Production

Once you are ready to create a production build of your app, you can use Expo’s build service. This allows you to create a standalone app that can be distributed to the App Store or Google Play.

To create a production build, run:

### Building for Android:

```bash
expo build:android
```

### Building for iOS:

```bash
expo build:ios
```

Note: Building for iOS requires macOS and an Apple developer account.

Once the build completes, Expo will provide a link to download the `.apk` (for Android) or `.ipa` (for iOS) file.

## Troubleshooting

If you encounter any issues during the installation or running of your project, the following steps might help:

- Make sure your system is up to date and has all necessary dependencies.
- If the Expo CLI isn’t working as expected, try uninstalling and reinstalling it using `npm uninstall -g expo-cli` and then `npm install -g expo-cli`.
- You can also visit the [Expo forums](https://forums.expo.dev/) for further support.

## Conclusion

Now you have Expo installed and know how to create, run, and build a project on Linux. Expo simplifies mobile app development and testing, and with these steps, you can start building cross-platform apps with React Native right away!

For more details, check out the official [Expo documentation](https://docs.expo.dev/).