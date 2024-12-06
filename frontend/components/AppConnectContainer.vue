<script setup lang="ts">
defineProps<{
  apps: string[]
}>();

</script>

<template>
  <UContainer :ui="{ padding: 'px-0' }" class="min-w-full flex flex-wrap justify-between border-t-2 border-black pt-8">
    <UButton variant="ghost" v-for="(app, index) in apps" :key="index" :icon="app" class="app_button basis-1/3 flex justify-center" />
  </UContainer>
</template>

<style scoped>
:deep(.app_button span) {
  height: 5rem;
  width: 5rem;
}
</style>
