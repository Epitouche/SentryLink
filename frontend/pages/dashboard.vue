<script setup>
import { ref } from 'vue';
import ProjectPageCard from '@/components/ProjectPageCard.vue';

const selected = ref(true);

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
    name: 'Project 8',
    brokenLinks: [
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 1', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 2', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 3', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 4', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 5', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 6', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 7', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 8', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 9', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 10', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 11', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 12', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 13', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 14', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 15', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 16', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
    ]
});

</script>

<template>
    <div class="header text-[7rem] font-[750] py-20 px-72">
        <h1>Hello, Name</h1> <!-- Replace Name with logged in user's name -->
    </div>

    <div class="carousel px-5">
        <Carousel :component-type="ProjectPageCard" :items="projects" :display-number="5" />
    </div>

    <div class="last-project-container pt-20">
        <div class="last-project-header text-[5rem] font-[750] px-32">
            <h2>Last Project: {{ lastProject.name }}</h2>
        </div>
        <div class="list">
            <div v-for="(link, index) in lastProject.brokenLinks" :key="index">
                <div class="list flex flex-row justify-between p-5 items-center">
                    <div class="flex flex-row items-center gap-4">
                        <div class="pr-5">
                            <UCheckbox v-model="selected" :ui="{rounded: 'rounded-none', border: 'border-2 border-black'}" name="checkbox" color="blue" />
                        </div>
                        <UAvatar :src="link.image" alt="temp" />
                        <div class="flex flex-col">
                            <h3>{{ link.name }}</h3>
                            <p>{{ link.url }}</p>
                        </div>
                    </div>
                    <div class="flex flex-row">
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
</style>
