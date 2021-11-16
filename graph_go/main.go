package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.POST("/", func(c *gin.Context) {
		var adjancencyMatrix [][]int
		c.BindJSON(&adjancencyMatrix)
		c.JSON(http.StatusOK, doStuff(adjancencyMatrix))
	})

	return r
}

func doStuff(adjacencyMatrix [][]int) [][]interface{} {
	graph := NewGraph(len(adjacencyMatrix))

	for i := 0; i < len(adjacencyMatrix); i += 1 {
		for j := 0; j < len(adjacencyMatrix[i]); j += 1 {
			weight := adjacencyMatrix[i][j]

			if weight != 0 {
				// graph.Connect(fmt.Sprintf("Vertex %d", i+1), fmt.Sprintf("Vertex %d", j+1), weight)
				graph.Connect(i, j, weight)
			}
		}
	}

	components := graph.FindConnectedComponents()
	componentsAsSlice := make([][]interface{}, 0, len(components))

	for root := range components {
		componentAsSlice := make([]interface{}, 0, len(components[root]))
		for vertex := range components[root] {
			componentAsSlice = append(componentAsSlice, vertex)
		}
		componentsAsSlice = append(componentsAsSlice, componentAsSlice)
	}

	return componentsAsSlice
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
