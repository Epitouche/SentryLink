<template>
  <div>
    <h1>Hello MindMapPage</h1>
    <div ref="graph" class="graph-container"></div>
  </div>
</template>

<script>
import * as d3 from 'd3';

/**
 * MindMapPage Component
 * 
 * This Vue component renders a mind map using D3.js. It creates a force-directed graph
 * with nodes and links, and allows for interactive dragging of nodes.
 * 
 * Component Name: MindMapPage
 * 
 * Lifecycle Hooks:
 * - mounted: Calls the createGraph method to initialize the graph when the component is mounted.
 * 
 * Methods:
 * - createGraph: Initializes the D3.js graph, sets up the simulation, and appends SVG elements for nodes and links.
 * - dragStarted: Handles the start of a drag event on a node, setting fixed positions and restarting the simulation.
 * - dragged: Updates the fixed positions of a node during a drag event.
 * - dragEnded: Handles the end of a drag event on a node, releasing fixed positions and stopping the simulation.
 * 
 * D3.js Configuration:
 * - Data: An array of nodes with id and group properties.
 * - Links: An array of links with source and target properties.
 * - SVG: Appends an SVG element to the component's graph reference with specified width and height.
 * - Simulation: Configures the force simulation with link, charge, and center forces.
 * - Nodes: Appends circle elements for each node, with colors based on their group and drag event handlers.
 * - Links: Appends line elements for each link, with specified stroke width and color.
 * - Ticked: Updates the positions of nodes and links on each tick of the simulation.
 */
export default {
  name: 'MindMapPage',
  mounted() {
    this.createGraph();
  },
  methods: {
    createGraph() {
      const data = [
        { id: 1, group: 'A' },
        { id: 2, group: 'A' },
        { id: 3, group: 'B' },
        { id: 4, group: 'B' },
        { id: 5, group: 'C' },
      ];

      const links = [
        { source: 1, target: 2 },
        { source: 1, target: 3 },
        { source: 3, target: 4 },
        { source: 4, target: 5 },
      ];

      const width = 1920;
      const height = 1080;

      const svg = d3.select(this.$refs.graph)
        .append('svg')
        .attr('width', width)
        .attr('height', height);

      const simulation = d3.forceSimulation(data)
        .force('link', d3.forceLink(links).id(d => d.id).distance(50))
        .force('charge', d3.forceManyBody().strength(-300))
        .force('center', d3.forceCenter(width / 2, height / 2));

      const link = svg.append('g')
        .selectAll('line')
        .data(links)
        .enter()
        .append('line')
        .attr('stroke-width', 2)
        .attr('stroke', '#999');

      const node = svg.append('g')
        .selectAll('circle')
        .data(data)
        .enter()
        .append('circle')
        .attr('r', 10)
        .attr('fill', d => (d.group === 'A' ? '#ff6347' : d.group === 'B' ? '#4682b4' : '#32cd32'))
        .call(d3.drag()
          .on('start', (event, d) => this.dragStarted(event, d, simulation))
          .on('drag', this.dragged)
          .on('end', (event, d) => this.dragEnded(event, d, simulation)));

      const ticked = () => {
        link
          .attr('x1', d => d.source.x)
          .attr('y1', d => d.source.y)
          .attr('x2', d => d.target.x)
          .attr('y2', d => d.target.y);

        node
          .attr('cx', d => d.x)
          .attr('cy', d => d.y);
      };

      simulation.on('tick', ticked);
    },
    dragStarted(event, d, simulation) {
      if (!event.active) simulation.alphaTarget(0.3).restart();
      d.fx = d.x;
      d.fy = d.y;
    },
    dragged(event, d) {
      d.fx = event.x;
      d.fy = event.y;
    },
    dragEnded(event, d, simulation) {
      if (!event.active) simulation.alphaTarget(0);
      d.fx = null;
      d.fy = null;
    },
  },
};
</script>

<style scoped>
h1 {
  font-family: Arial, sans-serif;
  color: #333;
}

.graph-container {
  margin-top: 10px;
}
</style>
