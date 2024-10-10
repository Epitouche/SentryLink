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
        <div class="card-content flex flex-col gap-3">
            <div class="card-title"><p>{{ src.name }}</p></div>
            <div class="card-info"><p>{{ src.info }}</p></div>
        </div>
    </div>
</template>

<style scoped>
.card {
    border: 2px solid #DCDDDE;
    border-radius: 1.5rem;
    box-shadow: 0 4px 0 rgba(0, 0, 0, 0.1);
    padding: 0.8rem;
}

.card-title {
    text-align: center;
}

.card-content p {
    font-size: 3rem;
}

.card-title p {
    font-weight: 500;
}

.card-info p {
    font-weight: 400;
}

.card-info {
    background-color: #DCDDDE;
    border: 2px solid #DCDDDE;
    border-radius: 1.4rem;
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
    padding-left: 1rem;
    padding-right: 1rem;
}
</style>
