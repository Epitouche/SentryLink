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
        <button @click="prev" :disabled="currentIndex === 0" class="carousel-arrow">Previous</button>
        <div class="carousel-items">
            <div v-for="(item, index) in displayedItems" :key="index" class="item-wrapper">
                <component :is="props.componentType" :src="item" />
            </div>
        </div>
        <button @click="next" :disabled="currentIndex + displayNumber >= items.length"
            class="carousel-arrow">Next</button>
    </div>
</template>

<style scoped>
.carousel {
    display: flex;
    align-items: center;
}

.carousel-arrow {
    background-color: #007BFF;
    color: white;
    border: none;
    padding: 10px 20px;
    cursor: pointer;
}

.carousel-button:disabled {
    background-color: gray;
    cursor: not-allowed;
}

.carousel-items {
    display: flex;
    overflow: hidden;
    flex-grow: 1;
}

.item-wrapper {
    flex: 0 0 25%;
    padding: 10px;
}
</style>
