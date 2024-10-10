<script setup>
import { ref, computed } from 'vue';

const props = defineProps({
    componentType: {
        type: Object,
        required: true
    },
    items: {
        type: Array,
        required: true
    },
    displayNumber: {
        type: Number,
        required: true
    }
});

const currentIndex = ref(0);

const displayedItems = computed(() => {
    return props.items.slice(currentIndex.value, currentIndex.value + props.displayNumber);
});

const next = () => {
    if (currentIndex.value + props.displayNumber < props.items.length) {
        currentIndex.value += props.displayNumber;
    }
};

const prev = () => {
    if (currentIndex.value > 0) {
        currentIndex.value -= props.displayNumber;
    }
};
</script>

<template>
    <div class="carousel">
        <button @click="prev" :disabled="currentIndex === 0" class="carousel-arrow"><Icon name="bytesize:chevron-left" size="40" /></button>
        <div class="carousel-items">
            <div v-for="(item, index) in displayedItems" :key="index" class="item-wrapper">
                <component :is="props.componentType" :src="item" />
            </div>
        </div>
        <button @click="next" :disabled="currentIndex + displayNumber >= items.length"
            class="carousel-arrow"><Icon name="bytesize:chevron-right" size="40" /></button>
    </div>
</template>

<style scoped>
.carousel {
    display: flex;
    justify-content: center;
    align-items: center;
}

.carousel-arrow {
    padding: 1rem;
    cursor: pointer;
}

.carousel-button:disabled {
    background-color: gray;
    cursor: not-allowed;
}

.carousel-items {
    display: flex;
}

.item-wrapper {
    padding: 2rem;
}
</style>
