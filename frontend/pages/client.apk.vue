<template>
    <div>
      <p>Downloading your APK...</p>
    </div>
  </template>
  
  <script>
  export default {
    mounted() {
      const apkUrl = '/apk/client.apk';
  
      const link = document.createElement('a');
      link.href = apkUrl;
      link.download = 'client.apk';
      link.click();
    }
  };
  </script>
  