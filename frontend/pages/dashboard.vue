<script>
export default {
    name: 'DashboardPage',
}

const projects = [ // Get array of all of the user's projects: name, info, and link
    { name: 'Project Name 1', info: 'Information 1' },
    { name: 'Project Name 2', info: 'Information 2' },
    { name: 'Project Name 3', info: 'Information 3' },
    { name: 'Project Name 4', info: 'Information 4' },
    { name: 'Project Name 5', info: 'Information 5' },
    { name: 'Project Name 6', info: 'Information 6' },
    { name: 'Project Name 7', info: 'Information 7' },
    { name: 'Project Name 8', info: 'Information 8' }
]

const projectCardBatches = computed(() => {
    const numberToDisplayAtATime = 4;
    const batches = [];
    for (let i = 0; i < projects.length; i += numberToDisplayAtATime) {
        batches.push(projects.slice(i, i + numberToDisplayAtATime));
    }
    return batches;
})

</script>

<template>
    <div>
        <h1>Hello, Name</h1> <!-- Replace Name with logged in user's name -->
    </div>
    <!-- Temp for testing -->
    <UContainer>
        <UCarousel v-slot="{ projectCardBatch }" :items="projectCardBatches"
            :ui="{ project: 'basis-full', container: 'rounded-lg' }"
            :prev-button="{ color: 'gray', icon: 'i-heroicons-arrow-left-20-solid', class: '-start-12' }"
            :next-button="{ color: 'gray', icon: 'i-heroicons-arrow-right-20-solid', class: '-end-12' }" arrows
            class="w-64 mx-auto"
        >
            <div v-for="(project, index) in projectCardBatch" :key="index">
                <NuxtLink to=""> <!-- Link to project page -->
                    <ProjectPageCard :src="project" /> <!-- Create ProjectPageCard component -->
                </NuxtLink>
            </div>
        </UCarousel>
    </UContainer>
    <UContainer>
        <h2>Last Project: Project Name</h2>
        <!-- Replace Project Name with actual name of the users last worked on project -->
        <!-- Create my own list component to display list of a projects broken links. Each item will have the broken link name,
        number, status (maybe), image, url, error, rating, option to like/favorite, etc. -->
        <!-- Make list scrollable so we dont have to scroll through the page. -->
    </UContainer>
</template>
