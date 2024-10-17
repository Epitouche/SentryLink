<template>
  <div ref="graph" style="width: 100%; height: 600px;"></div>
</template>

<script>
import ForceGraph3D from '3d-force-graph';

export default {
  name: 'MindMapPage',
  mounted() {
    this.create3DGraph();
  },
  methods: {
    create3DGraph() {
      const graphData = {
        nodes: [
          { id: 1, name: 'Node 1' },
          { id: 2, name: 'Node 2' },
          { id: 3, name: 'Node 3' },
        ],
        links: [
          { source: 1, target: 2 },
          { source: 2, target: 3 },
        ],
      };

      const Graph = ForceGraph3D()(this.$refs.graph)
        .graphData(graphData)
        .nodeLabel('name')
        .nodeAutoColorBy('group');
    },
  },
};
</script>

<style scoped>
.graph-container {
  width: 100%;
  height: 100%;
}
</style>