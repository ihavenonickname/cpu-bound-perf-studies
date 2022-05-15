import sys
from json import dumps
from typing import List

from bottle import route, run, request, BaseRequest, response

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

    response.content_type = 'application/json'

    return dumps([ list(x) for x in components.values() ])

def main():
    sys.setrecursionlimit(3000)

    KB = 1024
    MB = 1024 * KB
    BaseRequest.MEMFILE_MAX = 50 * MB

    run(host='0.0.0.0', port=8080)

if __name__ == '__main__':
    main()
