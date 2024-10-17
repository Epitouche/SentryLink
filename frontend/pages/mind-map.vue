<template>
  <div>
    <h1>Mind Map</h1>
    <Graph :graphData="graphData" />
  </div>
</template>

<script>
import Graph from '@/components/3DGraph.vue'; // Adjust the import path as needed

export default {
  name: 'MindMapPage',
  components: {
    Graph,
  },
  data() {
    return {
      graphData: {
        nodes: [
          { id: 1, name: 'Node 1', group: 1 },
          { id: 2, name: 'Node 2', group: 2 },
          { id: 3, name: 'Node 3', group: 1 },
        ],
        links: [
          { source: 1, target: 2 },
          { source: 2, target: 3 },
        ],
      },
    };
  },
};
</script>

<style scoped>
</style>
