<template>
  <div>
    <h1>Mind Map</h1>
    <PillButton link1="/mind-map" link2="/mind-map" name1="2D" name2="3D" color1="bg-red-500" color2="bg-blue-500"/>
    <Graph3D :graphData="graphData" />
    <!-- <Graph2D :graphData="graphData" /> -->
  </div>
</template>

<script>
import PillButton from '@/components/PillButton.vue';
import Graph3D from '@/components/3DGraph.vue'; // Adjust the import path as needed
// import Graph2D from '@/components/2DGraph.vue'; // Adjust the import path as needed

export default {
  name: 'MindMapPage',
  components: {
    Graph3D,
    PillButton,
    // Graph2D,
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
