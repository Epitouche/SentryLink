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
  <div class="flex flex-row max-w-screen-xl">
    <div class="flex flex-row gap-10" v-if="isReversed">
      <img :src="image" alt="" class="w-1/2 object-cover" />
      <div class="flex flex-col gap-5 justify-center p-4">
        <h1 class="text-3xl font-bold">{{ title }}</h1>
        <h2 class="text-xl opacity-50">{{ subtitle }}</h2>
        <p>{{ content }}</p>
      </div>
    </div>
    <div class="flex flex-row gap-10" v-else>
      <div class="flex flex-col gap-5 justify-center w-1/2 p-4">
        <h1 class="text-3xl font-bold">{{ title }}</h1>
        <h2 class="text-xl opacity-50">{{ subtitle }}</h2>
        <p>{{ content }}</p>
      </div>
      <img :src="image" alt="" class="w-1/2 object-cover" />
    </div>
  </div>
</template>
