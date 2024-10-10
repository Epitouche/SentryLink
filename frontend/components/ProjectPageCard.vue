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
    <div class="card">
        <div class="card-content">
            <div class="card-title">{{ src.name }}</div>
            <div class="card-info">{{ src.info }}</div>
        </div>
    </div>
</template>

<style scoped>
.card {
    width: 100%;
    background-color: red;
}
</style>
