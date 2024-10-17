<script setup>
import Sidebar from "@/components/Sidebar.vue";
</script>

<template>
  <div class="flex min-h-screen">
    <!-- Conteneur pour gérer le comportement sticky -->
    <div class="relative">
      <!-- Sidebar sticky à 50% de la vue -->
      <aside class="sticky top-[40vh] p-4">
        <Sidebar />
      </aside>
    </div>

    <!-- Contenu de la page -->
    <main class="flex-1 p-4 overflow-auto">
      <NuxtPage />
      <!-- Contenu long pour activer le scroll -->
      <div class="h-[200vh]">
      </div>
    </main>
  </div>
</template>
