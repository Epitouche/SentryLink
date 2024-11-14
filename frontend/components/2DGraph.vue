<template>
  <div ref="graph" class="graph-container"></div>
</template>

<script>
export default {
  name: '2DGraph',
  props: {
    graphData: {
      type: Object,
      required: true,
    },
  },
  mounted() {
    this.create2DGraph();
  },
  methods: {
    async create2DGraph() {
      // Ensure we're running in a client-side environment
      if (process.client) {
        try {
          // Dynamically import ForceGraph2D
          const ForceGraph2D = (await import('force-graph')).default;

          // Create and initialize the graph
          const Graph = ForceGraph2D()(this.$refs.graph)
            .graphData(this.graphData)
            .backgroundColor('#ffffff') // White background
            .nodeLabel('name')
            .nodeAutoColorBy('group')
            .enableNodeDrag(true); // Enable dragging of nodes
        } catch (error) {
          console.error('Error creating 2D graph:', error);
        }
      }
    },
  },
};
</script>

<style scoped>
.graph-container {
  max-width: 100%;
  height: 600px; /* Fixed height */
  overflow: hidden;
}
</style>
