<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';

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

const responsiveDisplayNumber = ref(props.displayNumber);

const updateDisplayNumber = () => {
    if (window.innerWidth < 768) {
        responsiveDisplayNumber.value = 1;
    } else if (window.innerWidth < 1024) {
        responsiveDisplayNumber.value = 2;
    } else { 
        responsiveDisplayNumber.value = props.displayNumber;
    }
};

const displayedItems = computed(() => {
    return props.items.slice(currentIndex.value, currentIndex.value + responsiveDisplayNumber.value);
});

const next = () => {
    if (currentIndex.value + responsiveDisplayNumber.value < props.items.length) {
        currentIndex.value += responsiveDisplayNumber.value;
    }
};

const prev = () => {
    if (currentIndex.value > 0) {
        currentIndex.value -= responsiveDisplayNumber.value;
    }
};

onMounted(() => {
    updateDisplayNumber();
    window.addEventListener('resize', updateDisplayNumber);
});

onBeforeUnmount(() => {
    window.removeEventListener('resize', updateDisplayNumber);
});
</script>

<template>
    <div class="carousel flex justify-center items-center">
        <button @click="prev" :disabled="currentIndex === 0" class="carousel-arrow p-4 cursor-pointer">
            <Icon name="bytesize:chevron-left" size="40" />
        </button>
        <div class="carousel-items flex flex-row flex-shrink">
            <div v-for="(item, index) in displayedItems" :key="index" class="item-wrapper px-10">
                <component :is="props.componentType" :src="item" />
            </div>
        </div>
        <button @click="next" :disabled="currentIndex + responsiveDisplayNumber >= items.length"
            class="carousel-arrow p-4 cursor-pointer">
            <Icon name="bytesize:chevron-right" size="40" />
        </button>
    </div>
</template>

<style scoped>
.carousel {
    overflow: hidden;
}

.carousel-items {
    transition: transform 0.5s ease;
}

.item-wrapper {
    max-width: 100%;
}

.carousel-button:disabled {
    background-color: gray;
    cursor: not-allowed;
}

</style>
