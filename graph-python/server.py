import sys
from typing import List

from bottle import route, run, request, BaseRequest

from graph import Graph

@route('/', method='POST')
def index():
    adjacency_matrix: List[List[int]] = request.json

    graph = Graph()

    for i, row in enumerate(adjacency_matrix):
        for j, weight in enumerate(row):
            if weight:
                graph.connect(i, j, weight)

    components = graph.find_connected_components()

    return { k: list(v) for k, v in components.items() }

def main():
    sys.setrecursionlimit(3000)

    KB = 1024
    MB = 1024 * KB
    BaseRequest.MEMFILE_MAX = 50 * MB

    run(host='0.0.0.0', port=8080)

if __name__ == '__main__':
    main()
