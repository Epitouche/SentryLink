<script setup lang="ts">
async function getCurrentUser() {
  return await $fetch('backend:8080/api/v1/github/auth')
}
</script>

<template>
  <div>
    <UButton to=getCurrentUser()>Button</UButton>
  </div>
</template>
