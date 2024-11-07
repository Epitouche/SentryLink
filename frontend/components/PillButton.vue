<template>
  <div class="flex w-24 overflow-hidden rounded-full text-white">
    <!-- First half of the button -->
    <a
      :href="link1"
      :class="[
        color1, 
        'w-1/2 px-1 py-2 text-center hover:opacity-90 transition'
      ]"
    >
      {{ name1 }}
    </a>

    <!-- Second half of the button -->
    <a
      :href="link2"
      :class="[
        color2, 
        'w-1/2 px-1 py-2 text-center hover:opacity-90 transition'
      ]"
    >
      {{ name2 }}
    </a>
  </div>
</template>

<script>
export default {
  name: 'PillButton',
  props: {
    link1: {
      type: String,
      required: true,
    },
    name1: {
      type: String,
      required: true,
    },
    link2: {
      type: String,
      required: true,
    },
    name2: {
      type: String,
      required: true,
    },
    color1: {
      type: String,
      required: true,
    },
    color2: {
      type: String,
      required: true,
    },
  },
};
</script>

<style scoped>
/* Optionally, add scoped styles if needed */
</style>
