from typing import Dict, Any, Set, List
from collections import deque

class Graph():
    _edges: Dict[Any, Dict[Any, int]]
    _in_neighbors: Dict[Any, Set[Any]]

    def __init__(self):
        self._edges = {}
        self._in_neighbors = {}

    def connect(self, vertex_from: Any, vertex_to: Any, weight: int):
        if vertex_from not in self._edges:
            self._edges[vertex_from] = {}
            self._in_neighbors[vertex_from] = set()

        if vertex_to not in self._edges:
            self._edges[vertex_to] = {}
            self._in_neighbors[vertex_to] = set()

        self._edges[vertex_from][vertex_to] = weight
        self._in_neighbors[vertex_to].add(vertex_from)

    def find_all_out_neighbors(self, vertex):
        return set(self._edges[vertex].keys())

    def find_all_in_neighbors(self, vertex):
        return set(self._in_neighbors[vertex])

    def find_all_vertices(self):
        return set(self._edges.keys())

    def find_connected_components(self):
        unvisited = self.find_all_vertices()

        L = deque()

        for vertex in self.find_all_vertices():
            self._visit(vertex, unvisited, L)

        components: Dict[Any, Set[Any]] = {}

        for vertex in L:
            self._assign(vertex, vertex, components)

        return components

    def _visit(self, vertex: Any, unvisited: Set[Any], L: deque):
        if vertex in unvisited:
            unvisited.remove(vertex)

            for neighbor in self.find_all_out_neighbors(vertex):
                self._visit(neighbor, unvisited, L)

            L.appendleft(vertex)

    def _assign(self, vertex: Any, root: Any, components: Dict[Any, Set[Any]]):
        has_been_assigned = False

        for component in components.values():
            if vertex in component:
                has_been_assigned = True
                break

        if not has_been_assigned:
            if root not in components:
                components[root] = set()

            components[root].add(vertex)

            for neighbor in self.find_all_in_neighbors(vertex):
                self._assign(neighbor, root, components)

def main():
    graph = Graph()

    graph.connect('1', '2', 1)
    graph.connect('2', '3', 1)
    graph.connect('2', '4', 1)
    graph.connect('3', '1', 1)
    graph.connect('3', '4', 1)

    graph.connect('4', '5', 1)
    graph.connect('5', '4', 1)

    graph.connect('6', '7', 1)
    graph.connect('6', '5', 1)
    graph.connect('7', '6', 1)
    graph.connect('7', '8', 1)
    graph.connect('8', '7', 1)
    graph.connect('8', '5', 1)

    print(graph.find_connected_components())

if __name__ == '__main__':
    main()
