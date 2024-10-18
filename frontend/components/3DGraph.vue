<template>
    <div ref="graph" style="width: 100%; height: 600px;"></div>
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
        // Check if running in the client environment
        if (process.client) {
          try {
            // Dynamically import the library
            const ForceGraph3D = (await import('3d-force-graph')).default;
  
            // Create the graph
            const Graph = ForceGraph3D()(this.$refs.graph)
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
  