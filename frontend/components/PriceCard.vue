<script lang="ts" setup>

const props = defineProps<{
    plan: string;
    price: number;
    features: string[];
}>();

</script>
<template>
  <div class="flex flex-col justify-center items-center shadow-xl bg-cprimary-light rounded-3xl p-10">
    <div id="infos" class="backdrop-blur-xl">
        <h2 class="text-lg text-sky-500">{{ props.plan }}</h2>
        <h3>${{ props.price }}/m</h3>
    </div>
    <div id="features">
        <ul>
            <li v-for="feature in props.features" :key="feature">{{ feature }}</li>
        </ul>
    </div>
  </div>
</template>
