using System.Collections.Generic;
using System.Linq;
using GraphCsharp.Domain;
using Microsoft.AspNetCore.Mvc;

namespace GraphCsharp.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class HomeController : ControllerBase
    {
        [HttpGet]
        public ISet<ISet<string>> Get()
        {
            var graph = new Graph<string>(8);

            graph.Connect("1", "2", 1);
            graph.Connect("2", "3", 1);
            graph.Connect("2", "4", 1);
            graph.Connect("3", "1", 1);
            graph.Connect("3", "4", 1);

            graph.Connect("4", "5", 1);
            graph.Connect("5", "4", 1);

            graph.Connect("6", "7", 1);
            graph.Connect("6", "5", 1);
            graph.Connect("7", "6", 1);
            graph.Connect("7", "8", 1);
            graph.Connect("8", "7", 1);
            graph.Connect("8", "5", 1);

            var components = graph.FindConnectedComponents();

            return new HashSet<ISet<string>>(components.Values);
        }

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
