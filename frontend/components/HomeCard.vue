<script setup lang="ts">
import { defineProps } from "vue";

const props = defineProps<{
  title: string;
  subtitle: string;
  content: string;
  image: string;
  isReversed: boolean;
}>();
</script>

<template>
  <div class="flex flex-col md:flex-row max-w-screen-xl gap-10">
    <!-- Si isReversed est vrai -->
    <div class="flex flex-col md:flex-row gap-10" v-if="isReversed">
      <img :src="image" alt="" class="w-full md:w-1/2 object-cover" />
      <div class="flex flex-col gap-5 justify-center p-4">
        <h1 class="text-2xl md:text-3xl font-bold">{{ title }}</h1>
        <h2 class="text-lg md:text-xl opacity-50">{{ subtitle }}</h2>
        <p>{{ content }}</p>
      </div>
    </div>

    <!-- Sinon, disposition normale -->
    <div class="flex flex-col md:flex-row gap-10" v-else>
      <div class="flex flex-col gap-5 justify-center w-full md:w-1/2 p-4">
        <h1 class="text-2xl md:text-3xl font-bold">{{ title }}</h1>
        <h2 class="text-lg md:text-xl opacity-50">{{ subtitle }}</h2>
        <p>{{ content }}</p>
      </div>
      <img :src="image" alt="" class="w-full md:w-1/2 object-cover" />
    </div>
  </div>
</template>
