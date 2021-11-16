const express = require('express')
const { Graph } = require('./graph.js')
const app = express()
const port = 3000

app.use(express.json({ limit: '50mb' }))

app.post('/', (req, res) => {
  const adjacencyMatrix = req.body
  const graph = new Graph()

  for (let i = 0; i < adjacencyMatrix.length; i += 1) {
    for (let j = 0; j < adjacencyMatrix[i].length; j += 1) {
      const weight = adjacencyMatrix[i][j]

      if (weight !== 0) {
        graph.connect(i, j, weight)
      }
    }
  }

  const components = graph.findConnectedComponents()

  res.send(convertFromMap(components))
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})

function convertFromMap(components) {
  return [...components.values()].map(x => [...x])
}
