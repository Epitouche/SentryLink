<template>
  <div class="w-full">
    <div class="justify-end items-end flex">
      <PillButton link1="/mind-map/2D" link2="/mind-map/3D" name1="2D" name2="3D" color1="bg-red-500" color2="bg-blue-500"/>
    </div>
    <div>
      <Graph3D :graphData="graphData" />
    </div>
  </div>
</template>

<script>
import PillButton from '@/components/PillButton.vue';
import Graph3D from '@/components/3DGraph.vue'; // Adjust the import path as needed

export default {
  name: 'MindMapPage',
  components: {
    Graph3D,
    PillButton,
  },
  data() {
    return {
      graphData: {
        nodes: [
          { id: 1, name: 'Node 1', group: 1 },
          { id: 2, name: 'Node 2', group: 2 },
          { id: 3, name: 'Node 3', group: 1 },
          { id: 4, name: 'Node 4', group: 2 },
          { id: 5, name: 'Node 5', group: 1 },
        ],
        links: [
          { source: 1, target: 2 },
          { source: 2, target: 3 },
          { source: 3, target: 4 },
          { source: 4, target: 5 },
          { source: 5, target: 1 },
          { source: 1, target: 3 },
          { source: 2, target: 4 },
          { source: 3, target: 5 },
        ],
      },
    };
  },
};
</script>

<style scoped>
</style>
