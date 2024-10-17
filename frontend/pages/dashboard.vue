<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
import ProjectPageCard from '@/components/ProjectPageCard.vue';
import { useField, useForm } from 'vee-validate';
import * as yup from 'yup';

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
    ]
});

const columns = ref([
    { key: 'name', label: 'Name', sortable: true },
    { key: 'error', label: 'Value', sortable: true },
    { key: 'ping', label: 'Ping', sortable: true },
    { key: 'nbPageLink', label: 'Number', sortable: true }
]);

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

    return filtered;
});

function select(row) {
    const index = selected.value.findIndex((item) => item.id === row.id)
    if (index === -1) {
        selected.value.push(row)
    } else {
        selected.value.splice(index, 1)
    }
};

const selected = ref([]);

const itemsPerPage = 5;
const displayedPageLinks = ref([]);
const currentPage = ref(0);

function loadMoreItems() {
    const startIndex = currentPage.value * itemsPerPage;
    const nextPageLinks = lastProject.value.pageLinks.slice(startIndex, startIndex + itemsPerPage);

    displayedPageLinks.value = [...displayedPageLinks.value, ...nextPageLinks];

    currentPage.value++;
}

function onScroll() {
    const scrollableDiv = document.querySelector('.scrollable');
    const threshold = 1;  // You can customize this threshold

    if (scrollableDiv.scrollTop + scrollableDiv.clientHeight >= scrollableDiv.scrollHeight - threshold) {
        loadMoreItems();
    }
}

onMounted(() => {
    loadMoreItems();
    window.addEventListener('scroll', onScroll);
});

onBeforeUnmount(() => {
    window.removeEventListener('scroll', onScroll);
});

const popoverState = ref(false);

const state = ref({
    name: '',
    info: '',
});

// Define validation schema for the new project form
const projectSchema = yup.object({
    name: yup.string().required('Project name is required'),
    info: yup.string().required('Project info is required'),
});

// Initialize Vee-Validate form
const { handleSubmit, resetForm } = useForm({
    validationSchema: projectSchema
});

// Set up fields with validation
const { value: name, errorMessage: nameError } = useField('name');
const { value: info, errorMessage: infoError } = useField('info');

const onSubmit = (values) => {
    if (nameError.value || infoError.value) {
        console.log('Validation errors:', nameError.value, infoError.value);
        return;
    }
    console.log('Submitted values:', values);
    projects.value.push({ name: values.name, info: values.info });
    resetForm();
    popoverState.value = !popoverState.value;
};

const handleFormSubmit = handleSubmit(onSubmit);

</script>

<template>
    <div name="page-wrapper" class="flex flex-col gap-16 px-6 py-6 h-[100vh]">
        <div name="header" class="flex flex-row justify-between items-center px-20">
            <h1 class="text-8xl font-[750]">Hello, Name</h1>
            <div name="new-project-button-and-form">
                <UPopover overlay v-model:open="popoverState">
                    <UButton color="blue" label="New Project" icon="bytesize:plus" />
                    <template #panel>
                        <UForm :state="state" @submit="handleFormSubmit" class="">
                            <UFormGroup label="Project Name" name="name">
                                <UInput v-model="name" placeholder="Enter project name" />
                                <span v-if="nameError" class="text-red-600">{{ nameError }}</span>
                            </UFormGroup>

                            <UFormGroup label="Project Info" name="info">
                                <UInput v-model="info" placeholder="Enter project info" />
                                <span v-if="infoError" class="text-red-600">{{ infoError }}</span>
                            </UFormGroup>

                            <UButton type="submit">Submit</UButton>
                        </UForm>
                    </template>
                </UPopover>
            </div>
        </div>

        <div name="projects-carousel">
            <Carousel :component-type="ProjectPageCard" :items="projects" :display-number="5" />
        </div>

        <div name="last-project" class="flex flex-col gap-3">
            <h2 class="text-8xl font-[750] px-20">Last Project: {{ lastProject.name }}</h2>
            <div name="last-project-container" class="flex flex-col gap-2">
                <div name="link-search" class="flex justify-end px-2">
                    <UInput v-model="linkSearch" placeholder="Filter links..." color="white" variant="none" size="xl"
                        icon="bytesize:search" :trailing="true" class="w-1/4 border-2 border-black rounded-2xl" />
                </div>
                <div name="table" class="scrollable border-2 border-black rounded-2xl">
                    <UTable v-model="selected" :columns="columns" :rows="filteredRows" @select="select">
                        <template #name-data="{ row }">
                            <div class="flex flex-row gap-2">
                                <UAvatar :src="row.image" alt="image" />
                                <div class="flex flex-col">
                                    <p><strong>{{ row.name }}</strong></p>
                                    <p>{{ row.url }}</p>
                                </div>
                            </div>
                        </template>
                    </UTable>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.scrollable {
    /* flex: 1; */
    max-height: 40vh;
    overflow-y: auto;
}
</style>
