<script setup>
import { ref } from 'vue';
import ProjectPageCard from '@/components/ProjectPageCard.vue';

const projects = ref([ // Get array of all of the user's projects: name, info, etc.
    { name: 'Project 1', info: 'Information 1' },
    { name: 'Project 2', info: 'Information 2' },
    { name: 'Project 3', info: 'Information 3' },
    { name: 'Project 4', info: 'Information 4' },
    { name: 'Project 5', info: 'Information 5' },
    { name: 'Project 6', info: 'Information 6' },
    { name: 'Project 7', info: 'Information 7' },
    { name: 'Project 8', info: 'Information 8' }
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

</script>

<template>
    <div class="header text-8xl font-[750] py-16 flex justify-end w-1/3">
        <h1>Hello, Name</h1> <!-- Replace Name with logged in user's name -->
    </div>

    <div class="carousel">
        <Carousel :component-type="ProjectPageCard" :items="projects" :display-number="4" />
    </div>

    <div class="last-project-container pt-10">
        <div class="last-project-header text-7xl font-[750] pl-28 pb-8">
            <h2>Last Project: {{ lastProject.name }}</h2>
        </div>
        <div class="px-10 overflow-y-auto max-h-[47vh]">
            <div v-for="(link, index) in lastProject.brokenLinks" :key="index" class="item">
                <div class="flex flex-row justify-between py-4 px-8 items-center text-2xl">
                    <div class="flex flex-row items-center gap-5">
                        <div class="pr-5">
                            <UCheckbox v-model="selectedCheckboxes[index]"  :ui="{base: 'h-7 w-7', rounded: 'rounded-none', border: 'border-2 border-black'}" name="checkbox" color="blue" />
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
</template>

<style scoped>
.item:nth-child(even) {
    background-color: #F2F4F5;
    border: 1px solid #F2F4F5;
    border-radius: 1rem;
}
</style>
