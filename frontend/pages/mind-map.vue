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
 * This Vue component renders a mind map using D3.js. It creates a graph with nodes and links,
 * and allows for interactive dragging of nodes with attraction forces applied to linked nodes.
 * 
 * Component Name: MindMapPage
 * 
 * Lifecycle Hooks:
 * - mounted: Calls the createGraph method to initialize the graph when the component is mounted.
 * 
 * Methods:
 * - createGraph: Initializes the graph with nodes and links, sets up the D3 simulation, and defines the ticked function to update positions.
 * - dragStarted: Handles the start of a drag event, setting fixed positions and adding attraction forces.
 * - dragged: Updates the position of the dragged node during the drag event.
 * - dragEnded: Handles the end of a drag event, removing fixed positions and attraction forces.
 * - addAttractionForce: Adds an attraction force to the simulation for nodes linked to the dragged node.
 * - removeAttractionForce: Removes the attraction force from the simulation.
 * 
 * Data:
 * - data: Array of node objects with id and group properties.
 * - links: Array of link objects with source and target properties.
 * 
 * D3.js Elements:
 * - svg: The SVG container for the graph.
 * - simulation: The D3 force simulation for the graph.
 * - link: The lines representing links between nodes.
 * - node: The circles representing nodes.
 * 
 * Styles:
 * - Scoped styles for the component, including styles for the h1 element and the graph container.
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
        .force('link', d3.forceLink(links).id(d => d.id).distance(100))
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
          .on('start', (event, d) => this.dragStarted(event, d, simulation, links))
          .on('drag', this.dragged)
          .on('end', (event, d) => this.dragEnded(event, d, simulation, links)));

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

    dragStarted(event, d, simulation, links) {
      if (!event.active) simulation.alphaTarget(0.3).restart();
      d.fx = d.x;
      d.fy = d.y;

      this.addAttractionForce(simulation, d, links);
    },

    dragged(event, d) {
      d.fx = event.x;
      d.fy = event.y;
    },

    dragEnded(event, d, simulation, links) {
      if (!event.active) simulation.alphaTarget(0);
      d.fx = null;
      d.fy = null;

      this.removeAttractionForce(simulation);
    },

    addAttractionForce(simulation, draggedNode, links) {
      const linkedNodes = links
        .filter(l => l.source.id === draggedNode.id || l.target.id === draggedNode.id)
        .map(l => (l.source.id === draggedNode.id ? l.target : l.source));

      const attractionForce = (alpha) => {
        linkedNodes.forEach(linkedNode => {
          linkedNode.vx += (draggedNode.x - linkedNode.x) * 0.1 * alpha;
          linkedNode.vy += (draggedNode.y - linkedNode.y) * 0.1 * alpha;
        });
      };

      simulation.force('attraction', attractionForce);
    },

    removeAttractionForce(simulation) {
      simulation.force('attraction', null);
    }
  },
};
</script>

<style scoped>
h1 {
  font-family: Arial, sans-serif;
  color: #333;
}

.graph-container {
  margin-top: 20px;
}
</style>
