<template>
  <div ref="graph" class="graph-container"></div>
</template>

<script>
export default {
  name: '3DGraph',
  props: {
    graphData: {
      type: Object,
      required: true,
    },
  },
  mounted() {
    this.create3DGraph();
  },
  methods: {
    async create3DGraph() {
      if (process.client) {
        try {
          const ForceGraph3D = (await import('3d-force-graph')).default;

          const Graph = ForceGraph3D()(this.$refs.graph)
            .graphData(this.graphData)
            .backgroundColor('#ffffff')
            .nodeLabel('name')
            .nodeAutoColorBy('group')
            .cameraPosition({ z: 200 })
            .enableNodeDrag(true);
        } catch (error) {
          console.error('Error creating 3D graph:', error);
        }
      }
    },
  },
};
</script>

<style scoped>
.graph-container {
  max-width: 100%;
  height: 600px; /* Make sure height is fixed to what you want */
  overflow: hidden;
}
</style>
