import { Component, OnInit } from '@angular/core';
import {
  select,
  forceSimulation,
  forceManyBody,
  forceCenter,
  forceLink,
  scaleOrdinal,
  schemeCategory10,
  drag,
  SimulationNodeDatum,
} from 'd3';

@Component({
  selector: 'app-mind-map',
  standalone: true,
  imports: [],
  templateUrl: './mind-map.component.html',
  styleUrl: './mind-map.component.css',
})
export class MindMapComponent implements OnInit {
  title = 'mind-map';
  private nodes = [
    { index: 0, name: '', group: 0 },
    { index: 1, name: 'Fruit', group: 1 },
    { index: 2, name: 'Vegetable', group: 2 },
    { index: 3, name: 'Orange', group: 1 },
    { index: 4, name: 'Apple', group: 1 },
    { index: 5, name: 'Banana', group: 1 },
    { index: 6, name: 'Peach', group: 1 },
    { index: 7, name: 'Bean', group: 2 },
    { index: 8, name: 'Pea', group: 2 },
    { index: 9, name: 'Carrot', group: 2 },
  ];
  private links = [
    { source: this.nodes[0], target: this.nodes[1] },
    { source: this.nodes[0], target: this.nodes[2] },
    { source: this.nodes[1], target: this.nodes[3] },
    { source: this.nodes[1], target: this.nodes[4] },
    { source: this.nodes[1], target: this.nodes[5] },
    { source: this.nodes[1], target: this.nodes[6] },
    { source: this.nodes[2], target: this.nodes[7] },
    { source: this.nodes[2], target: this.nodes[8] },
    { source: this.nodes[2], target: this.nodes[9] },
  ];
  private color = scaleOrdinal(schemeCategory10);

  ngOnInit(): void {
    const div: HTMLDivElement = document.querySelector('div');
    const svg = select('div').append('svg').attr('viewBox', '0 0 500 500');

    console.log(div.clientWidth);

    const link = svg
      .append('g')
      .attr('stroke', '#999')
      .attr('stroke-opacity', 0.6)
      .selectAll('line')
      .data(this.links)
      .join('line');

    const node = svg
      .append('g')
      .attr('class', 'nodes')
      .selectAll('g')
      .data(this.nodes)
      .enter()
      .append('g');

    const circles = node
      .append('circle')
      .attr('r', 10)
      .style('fill', (n) => this.color(n.group))
      .style('cursor', 'pointer')
      .on('dblclick', (e) => alert(e.srcElement.__data__.name))
      .call(
        drag()
          .on('start', (e, d) => dragstarted(e, d))
          .on('drag', (e, d) => dragged(e, d))
          .on('end', (e, d) => dragended(e, d))
      );

    const labels = node
      .append('text')
      .text((n) => n.name)
      .attr('x', 12)
      .attr('y', 3)
      .style('font-size', '12px')
      .style('color', (n) => this.color('' + n.group));

    node.append('title').text((n) => n.name);

    const simulation = forceSimulation(this.nodes)
      .force(
        'link',
        forceLink(this.links).id((d) => d.id)
        .distance(50)
      )
      .force('charge', forceManyBody().strength(-200))
      .force('center', forceCenter(div.clientWidth / 2, 200))
      .tick()
      .on('tick', () => {
        node.attr('transform', (n) => 'translate(' + n.x + ',' + n.y + ')');
        link
          .attr('x1', (l) => l.source.x)
          .attr('y1', (l) => l.source.y)
          .attr('x2', (l) => l.target.x)
          .attr('y2', (l) => l.target.y);
      });

    const dragstarted = (e: any, d: SimulationNodeDatum) => {
      if (!e.active) {
        simulation.alphaTarget(0.3).restart();
      }
      d.fx = d.x;
      d.fy = d.y;
    };

    const dragged = (e: any, d: SimulationNodeDatum) => {
      d.fx = e.x;
      d.fy = e.y;
    };

    const dragended = (e: any, d: SimulationNodeDatum) => {
      if (!e.active) {
        simulation.alphaTarget(0);
      }
      d.fx = null;
      d.fy = null;
    };
  }
}
