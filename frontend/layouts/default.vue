<script setup>
import Sidebar from "@/components/Sidebar.vue";
</script>

<template>
  <div class="flex min-h-screen">
    <div class="relative">
      <aside class="sticky top-[40vh] p-4">
        <Sidebar />
      </aside>
    </div>
    <main class="flex-1 p-4 overflow-auto">
      <NuxtPage />
    </main>
  </div>
</template>
