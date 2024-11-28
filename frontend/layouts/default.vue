<script setup>
import Sidebar from "@/components/Sidebar.vue";
</script>

<template>
  <div class=""> <!-- relative -->
    <aside class="fixed z-50 top-1/2 left-0 transform -translate-y-1/2 p-4">
      <Sidebar />
    </aside>
    <!-- <div class="flex min-h-screen"> -->
      <!-- <main class="flex-1 p-4"> -->
        <main class="h-screen w-screen">
          <NuxtPage />
        </main>
      <!-- </main> -->
    <!-- </div> -->
  </div>
</template>
