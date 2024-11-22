<template>
    <div>
      <p>Downloading your APK...</p>
    </div>
  </template>
  
  <script>
  export default {
    mounted() {
      // Path to the APK file in the static folder
      const apkUrl = '/apk/client.apk'; // Updated to include the subfolder
  
      // Create an invisible link and trigger the download
      const link = document.createElement('a');
      link.href = apkUrl;
      link.download = 'client.apk'; // Optional: suggest a file name
      link.click();
  
      // Optionally redirect after download
      // this.$router.push('/');
    }
  };
  </script>
  