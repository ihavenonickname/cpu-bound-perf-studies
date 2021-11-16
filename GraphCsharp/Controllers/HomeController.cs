using System.Collections.Generic;
using GraphCsharp.Domain;
using Microsoft.AspNetCore.Mvc;

namespace GraphCsharp.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class HomeController : ControllerBase
    {
        [HttpPost]
        public ISet<ISet<int>> Post([FromBody] List<List<int>> adjacencyMatrix)
        {
            var graph = new Graph<int>(adjacencyMatrix.Count);

            for (var i = 0; i < adjacencyMatrix.Count; i += 1)
            {
                for (var j = 0; j < adjacencyMatrix[i].Count; j += 1)
                {
                    var weight = adjacencyMatrix[i][j];

                    if (weight != 0)
                    {
                        graph.Connect(i, j, weight);
                    }
                }
            }

            var components = graph.FindConnectedComponents();

            return new HashSet<ISet<int>>(components.Values);
        }
    }
}
