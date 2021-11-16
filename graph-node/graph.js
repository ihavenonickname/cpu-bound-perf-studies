// class Graph {
//   constructor() {
//     this._edges = {}
//     this._inNeighbors = {}
//   }

//   connect(vertexFrom, vertexTo, weight) {
//     vertexFrom = vertexFrom.toString()
//     vertexTo = vertexTo.toString()

//     if (!this._edges.hasOwnProperty(vertexFrom)) {
//       this._edges[vertexFrom] = {}
//       this._inNeighbors[vertexFrom] = new Set()
//     }

//     if (!this._edges.hasOwnProperty(vertexTo)) {
//       this._edges[vertexTo] = {}
//       this._inNeighbors[vertexTo] = new Set()
//     }

//     this._edges[vertexFrom][vertexTo] = weight
//     this._inNeighbors[vertexTo].add(vertexFrom)
//   }

//   findOutNeighbors(vertex) {
//     return new Set(Object.keys(this._edges[vertex]))
//   }

//   findInNeighbors(vertex) {
//     return new Set(this._inNeighbors[vertex])
//   }

//   findVertices() {
//     return new Set(Object.keys(this._edges))
//   }

//   findConnectedComponents() {
//     const unvisited = this.findVertices()

//     const L = []

//     for (const vertex of this.findVertices()) {
//       this.visit(vertex, unvisited, L)
//     }

//     const components = {}

//     for (let i = L.length - 1; i >= 0; i -= 1) {
//       this.assign(L[i], L[i], components)
//     }

//     return components
//   }

//   visit(vertex, unvisited, L) {
//     if (unvisited.has(vertex)) {
//       unvisited.delete(vertex)

//       for (const neighbor of this.findOutNeighbors(vertex)) {
//         this.visit(neighbor, unvisited, L)
//       }

//       L.push(vertex)
//     }
//   }

//   assign(vertex, root, components) {
//     let hasBeenAssigned = false

//     for (const component of Object.values(components)) {
//       if (component.has(vertex)) {
//         hasBeenAssigned = true
//         break
//       }
//     }

//     if (!hasBeenAssigned) {
//       if (!components.hasOwnProperty(root)) {
//         components[root] = new Set()
//       }

//       components[root].add(vertex)

//       for (const neighbor of this.findInNeighbors(vertex)) {
//         this.assign(neighbor, root, components)
//       }
//     }
//   }
// }

class Graph {
  constructor() {
    this._edges = new Map()
    this._inNeighbors = new Map()
  }

  connect(vertexFrom, vertexTo, weight) {
    if (!this._edges.has(vertexFrom)) {
      this._edges.set(vertexFrom, new Map())
      this._inNeighbors.set(vertexFrom, new Set())
    }

    if (!this._edges.has(vertexTo)) {
      this._edges.set(vertexTo, new Map())
      this._inNeighbors.set(vertexTo, new Set())
    }

    this._edges.get(vertexFrom).set(vertexTo, weight)
    this._inNeighbors.get(vertexTo).add(vertexFrom)
  }

  findOutNeighbors(vertex) {
    return new Set(this._edges.get(vertex).keys())
  }

  findInNeighbors(vertex) {
    return new Set(this._inNeighbors.get(vertex))
  }

  findVertices() {
    return new Set(this._edges.keys())
  }

  findConnectedComponents() {
    const unvisited = this.findVertices()

    const L = []

    for (const vertex of this.findVertices()) {
      this.visit(vertex, unvisited, L)
    }

    const components = new Map()

    for (let i = L.length - 1; i >= 0; i -= 1) {
      this.assign(L[i], L[i], components)
    }

    return components
  }

  visit(vertex, unvisited, L) {
    if (unvisited.has(vertex)) {
      unvisited.delete(vertex)

      for (const neighbor of this.findOutNeighbors(vertex)) {
        this.visit(neighbor, unvisited, L)
      }

      L.push(vertex)
    }
  }

  assign(vertex, root, components) {
    let hasBeenAssigned = false

    for (const component of components.values()) {
      if (component.has(vertex)) {
        hasBeenAssigned = true
        break
      }
    }

    if (!hasBeenAssigned) {
      if (!components.has(root)) {
        components.set(root, new Set())
      }

      components.get(root).add(vertex)

      for (const neighbor of this.findInNeighbors(vertex)) {
        this.assign(neighbor, root, components)
      }
    }
  }
}

module.exports = { Graph }

// const adjacenceMatrix = [
//     [0, 1, 0, 0, 0, 0, 0, 0],
//     [0, 0, 1, 1, 0, 0, 0, 0],
//     [1, 0, 0, 1, 0, 0, 0, 0],
//     [0, 0, 0, 0, 1, 0, 0, 0],
//     [0, 0, 0, 1, 0, 0, 0, 0],
//     [0, 0, 0, 0, 1, 0, 1, 0],
//     [0, 0, 0, 0, 0, 1, 0, 1],
//     [0, 0, 0, 0, 1, 0, 1, 0],
// ]
