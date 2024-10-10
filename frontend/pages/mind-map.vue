<template>
  <div>
    <h1>Hello World</h1>
    <div ref="graph" class="graph-container"></div>
  </div>
</template>

<script>
import * as d3 from 'd3';

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

      const width = 400;
      const height = 400;

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
        .attr('fill', d => (d.group === 'A' ? '#ff6347' : d.group === 'B' ? '#4682b4' : '#32cd32'));

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
