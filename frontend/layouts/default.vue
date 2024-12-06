<script setup>
import Sidebar from "@/components/Sidebar.vue";
</script>

<template>
  <div>
    <aside class="fixed z-50 top-1/2 left-0 transform -translate-y-1/2 p-4">
      <Sidebar />
    </aside>
        <main class="h-screen w-screen">
          <NuxtPage />
        </main>
  </div>
</template>
