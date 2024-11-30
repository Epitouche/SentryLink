<!-- Use skeleten compononet from nuxt for when page is loading -->
<template>
  <UContainer class="max-w-full p-2 flex flex-row justify-between">
     <div>
      <UButton color="black" variant="link"><Icon name="bytesize:arrow-left"/></UButton>
     </div>
     <div class="flex gap-3">
      <UButton color="black" variant="link"><Icon name="bytesize:paperclip" /></UButton>
      <UButton color="black" variant="link"><Icon name="bytesize:calendar" /></UButton>
      <UButton color="black" variant="link"><Icon name="bytesize:ellipsis-vertical" /></UButton>
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
  name: "ResultsPage",
};

definePageMeta({
  layout: "default",
});
</script>

<style scoped>
/* h1 {
  font-family: Arial, sans-serif;
  color: #333;
} */
</style>
