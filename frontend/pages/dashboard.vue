<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';
import ProjectPageCard from '@/components/ProjectPageCard.vue';

const projects = ref([ // Get array of all of the user's projects: name, info, etc.
    { name: 'Project 1', info: 'Information 1' },
    { name: 'Project 2', info: 'Information 2' },
    { name: 'Project 3', info: 'Information 3' },
    { name: 'Project 4', info: 'Information 4' },
    { name: 'Project 5', info: 'Information 5' },
    { name: 'Project 6', info: 'Information 6' },
    { name: 'Project 7', info: 'Information 7' },
    { name: 'Project 8', info: 'Information 8' },
    { name: 'Project 9', info: 'Information 9' },
    { name: 'Project 10', info: 'Information 10' }
]);

const lastProject = ref({ // Get the last project the user worked on: name, broken links: image, name, url, etc.
    name: 'Project Name',
    brokenLinks: [
        { image: 'https://picsum.photos/id/1000/600/400', name: 'Broken Link 1', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1010/600/400', name: 'Broken Link 2', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 3', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1001/600/400', name: 'Broken Link 4', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1040/600/400', name: 'Broken Link 5', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1050/600/400', name: 'Broken Link 6', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1060/600/400', name: 'Broken Link 7', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1070/600/400', name: 'Broken Link 8', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1080/600/400', name: 'Broken Link 9', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1001/600/400', name: 'Broken Link 10', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1009/600/400', name: 'Broken Link 11', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1006/600/400', name: 'Broken Link 12', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1005/600/400', name: 'Broken Link 13', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1004/600/400', name: 'Broken Link 14', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1003/600/400', name: 'Broken Link 15', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1002/600/400', name: 'Broken Link 16', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
    ]
});

const selectedCheckboxes = ref(lastProject.value.brokenLinks.map(() => false));

const scaleValue = ref(1);
const minScale = 1;
const maxScale = 1;

const updateScale = () => {
    if (typeof window !== 'undefined') {
        const widthScale = window.innerWidth / 2560;
        const heightScale = window.innerHeight / 1600;
        const newScale = Math.min(widthScale, heightScale);
        scaleValue.value = Math.max(minScale, Math.min(newScale, maxScale)); // Clamp between min and max scale
    }
};

onMounted(() => {
    updateScale();
    window.addEventListener('resize', updateScale);
});

onBeforeUnmount(() => {
    window.removeEventListener('resize', updateScale);
});
</script>

<template>
    <div class="page-wrapper" :style="{ transform: `scale(${scaleValue})`, transformOrigin: 'top left' }">
        <div class="header font-[750] flex justify-end w-1/3">
            <h1>Hello, Name</h1>
        </div>

        <div class="carousel">
            <Carousel :component-type="ProjectPageCard" :items="projects" :display-number="5" />
        </div>

        <div class="last-project-container">
            <div class="last-project-header font-[750] pl-28 pb-8">
                <h2>Last Project: {{ lastProject.name }}</h2>
            </div>

            <div class="px-10 overflow-y-auto max-h-[50vh]">
                <div v-for="(link, index) in lastProject.brokenLinks" :key="index" class="item">
                    <div class="flex flex-row justify-between py-4 px-8 items-center text-2xl">
                        <div class="flex flex-row items-center gap-5">
                            <div class="pr-5">
                                <UCheckbox v-model="selectedCheckboxes[index]"
                                    :ui="{ base: 'h-7 w-7', rounded: 'rounded-none', border: 'border-2 border-black' }"
                                    name="checkbox" color="blue" />
                            </div>
                            <UAvatar size="3xl" :src="link.image" alt="temp" />
                            <div class="flex flex-col gap-2">
                                <p><strong>{{ link.name }}</strong></p>
                                <p>url:{{ link.url }}</p>
                            </div>
                        </div>
                        <div>
                            <p>{{ link.error }}</p>
                        </div>
                        <div>
                            <p>{{ link.ping }}</p>
                        </div>
                        <div>
                            <p>{{ link.nbPageLink }}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.page-wrapper {
    transition: transform 0.2s ease;
}

.item:nth-child(even) {
    background-color: #F2F4F5;
    border: 1px solid #F2F4F5;
    border-radius: 1rem;
}

/* Responsive Typography */
h1 {
    font-size: clamp(1rem,56vw, 6rem);
}

h2 {
    font-size: clamp(0.5rem, 4vw, 4.5rem);
}

.text-2xl {
    font-size: clamp(1rem, 3vw, 1.5rem);
}

.header {
    padding: 2rem;
}

.carousel {
    padding: 0.5rem;
}

.last-project-container {
    padding-top: 2rem;
}
</style>

<!-- Mine
2560 x 1600

Desktops
1920 x 1080
1366 x 768
1280 x 1024
1024 x 768

Tablets
1024 x 768
768 x 1024
962 x 601
601 x 962

Mobiles
375 x 667
414 x 736
360 x 800
390 x 844 -->
