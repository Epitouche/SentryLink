<script setup>
import Sidebar from "@/components/Sidebar.vue";
</script>

<template>
  <div class="flex min-h-screen">
    <div class="fixed h-screen flex items-center justify-center">
      <aside class="sticky p-4">
        <Sidebar />
      </aside>
    </div>

    <main class="flex-1 p-4">
      <NuxtPage />
    </main>
  </div>
</template>