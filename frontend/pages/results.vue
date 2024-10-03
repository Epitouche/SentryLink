<!-- Use skeleten compononet from nuxt for when page is loading -->
<template>
  <UContainer class="bg-red-400 w-full p-2 flex flex-row justify-between">
     <div>
      <UButton><Icon name="bytesize:arrow-left" /></UButton>
     </div>
     <div>
      <UButton><Icon name="bytesize:paperclip" /></UButton>
      <UButton><Icon name="bytesize:calendar" /></UButton>
      <UButton><Icon name="bytesize:ellipsis-vertical" /></UButton>
     </div>
  </UContainer>
  <UContainer>
    <!-- Top bar thing with project name and stuff -->
  </UContainer>
  <UContainer>
    <!-- Table/Mindmap display buttons: use the tabs component -->
  </UContainer>
  <UContainer>
    <!-- Table: use table component -->
  </UContainer>
  <UContainer>
    <!-- Mindmap -->
  </UContainer>
</template>

<script>
export default {
  name: 'ResultsPage',
}
</script>

<style scoped>
/* h1 {
  font-family: Arial, sans-serif;
  color: #333;
} */
</style>
