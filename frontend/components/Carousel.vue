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
    <div class="carousel flex justify-center items-center">
        <button @click="prev" :disabled="currentIndex === 0" class="carousel-arrow p-4 cursor-pointer"><Icon name="bytesize:chevron-left" size="40" /></button>
        <div class="carousel-items flex flex-row">
            <div v-for="(item, index) in displayedItems" :key="index" class="item-wrapper py-10 px-12">
                <component :is="props.componentType" :src="item" />
            </div>
        </div>
        <button @click="next" :disabled="currentIndex + displayNumber >= items.length"
            class="carousel-arrow p-4 cursor-pointer"><Icon name="bytesize:chevron-right" size="40" /></button>
    </div>
</template>

<style scoped>
.carousel-button:disabled {
    background-color: gray;
    cursor: not-allowed;
}
</style>
