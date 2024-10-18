<template>
    <div ref="graph" style="width: 100%; height: 600px;"></div>
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
        // Check if running in the client environment
        if (process.client) {
          try {
            // Dynamically import the library
            const ForceGraph2D = (await import('force-graph')).default;
  
            // Create the graph
            const Graph = ForceGraph2D()(this.$refs.graph)
              .graphData(this.graphData)
              .backgroundColor('#ffffff')
              .nodeLabel('name')
              .nodeAutoColorBy('group')
              .cameraPosition({ z: 200 }) // Adjust camera position if needed
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
    width: 100%;
    height: 100%;
  }
  </style>
  