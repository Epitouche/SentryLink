<script setup>
import { defineProps } from 'vue';

const props = defineProps({
    src: {
        type: Object,
        required: true,
    },
});

</script>

<template>
    <div class="card rounded-3xl p-3">
        <div class="card-content flex flex-col gap-8">
            <div class="card-title text-center text-5xl font-medium py-2 px-4">{{ src.name }}</div>
            <div class="card-info text-4xl font-normal rounded-2xl py-4 px-6">{{ src.info }}</div>
        </div>
    </div>
</template>

<style scoped>
.card {
    border: 3px solid #F2F4F5;
    box-shadow: 0 5px 0 rgba(0, 0, 0, 0.1);
}

.card-info {
    background-color: #F2F4F5;
    border: 2px solid #F2F4F5;
}
</style>
