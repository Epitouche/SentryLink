<script setup lang="ts">
async function getCurrentUser() {
  return await $fetch("http://backend:8080/api/v1/github/auth");
}
</script>

<template>
  <div>
    <UButton @click="getCurrentUser()">Button</UButton>
  </div>
</template>
