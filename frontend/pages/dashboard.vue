<script setup>
import { ref, computed } from 'vue';
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
    pageLinks: [
        { id: 1, image: 'https://picsum.photos/id/1000/600/400', name: 'Broken Link 1', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 2, image: 'https://picsum.photos/id/1010/600/400', name: 'Broken Link 2', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 3, image: 'https://picsum.photos/id/1020/600/400', name: 'Broken Link 3', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 4, image: 'https://picsum.photos/id/1001/600/400', name: 'Broken Link 4', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 5, image: 'https://picsum.photos/id/1040/600/400', name: 'Broken Link 5', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 6, image: 'https://picsum.photos/id/1050/600/400', name: 'Broken Link 6', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 7, image: 'https://picsum.photos/id/1060/600/400', name: 'Broken Link 7', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 8, image: 'https://picsum.photos/id/1070/600/400', name: 'Broken Link 8', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 9, image: 'https://picsum.photos/id/1080/600/400', name: 'Broken Link 9', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 10, image: 'https://picsum.photos/id/1001/600/400', name: 'Broken Link 10', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 11, image: 'https://picsum.photos/id/1009/600/400', name: 'Broken Link 11', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 12, image: 'https://picsum.photos/id/1006/600/400', name: 'Broken Link 12', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 13, image: 'https://picsum.photos/id/1005/600/400', name: 'Broken Link 13', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 14, image: 'https://picsum.photos/id/1004/600/400', name: 'Broken Link 14', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 15, image: 'https://picsum.photos/id/1003/600/400', name: 'Broken Link 15', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
        { id: 16, image: 'https://picsum.photos/id/1002/600/400', name: 'Broken Link 16', url: 'https://www.google.com', error: 'msg: 404 page not found', ping: '40ms', nbPageLink: 3 },
    ]
});

const columns = ref([
    { key: 'name', label: 'Name', sortable: true },
    { key: 'error', label: 'Value', sortable: true },
    { key: 'ping', label: 'Ping', sortable: true },
    { key: 'nbPageLink', label: 'Number', sortable: true }
]);

const page = ref(1);
const rowsPerPage = 5;

const linkSearch = ref('');

const filteredRows = computed(() => {
  let filtered = lastProject.value.pageLinks;

  if (linkSearch.value) {
    filtered = filtered.filter((pageLink) => {
      return Object.values(pageLink).some((value) => {
        return String(value).toLowerCase().includes(linkSearch.value.toLowerCase());
      });
    });
  }

  const start = (page.value - 1) * rowsPerPage;
  const end = start + rowsPerPage;
  return filtered.slice(start, end);
});

const pageCount = computed(() => {
  const filtered = lastProject.value.pageLinks.filter((pageLink) => {
    return Object.values(pageLink).some((value) => {
      return String(value).toLowerCase().includes(linkSearch.value.toLowerCase());
    });
  });

  return Math.ceil(filtered.length / rowsPerPage);
});

function select(row) {
    const index = selected.value.findIndex((item) => item.id === row.id)
    if (index === -1) {
        selected.value.push(row)
    } else {
        selected.value.splice(index, 1)
    }
};

const selected = ref([lastProject.pageLinks]);

</script>

<template>
    <div class="page-wrapper" :style="{ transform: `scale(${scaleValue})`, transformOrigin: 'top left' }">
        <div class="header text-8xl font-[750]">
            <h1>Hello, Name</h1>
        </div>

        <div class="carousel">
            <Carousel :component-type="ProjectPageCard" :items="projects" :display-number="5" />
        </div>

        <div>
            <div class="last-project-header text-8xl font-[750]">
                <h2>Last Project: {{ lastProject.name }}</h2>
            </div>

            <div>
                <UInput v-model="linkSearch" placeholder="Filter links..." />
                <UTable v-model="selected" :columns="columns" :rows="filteredRows" @select="select" />
                <UPagination v-model="page" :page-count="pageCount" :total="filteredRows.length" />
            </div>

        </div>
    </div>
</template>

<style scoped>
.page-wrapper {
    transition: transform 0.2s ease;
}

/* .item:nth-child(even) {
    background-color: #F2F4F5;
    border: 1px solid #F2F4F5;
    border-radius: 1rem;
} */
</style>
