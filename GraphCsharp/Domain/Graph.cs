using System.Collections.Generic;

namespace GraphCsharp.Domain
{
    public class Graph<T>
    {
        private readonly Dictionary<T, Dictionary<T, int>> Edges;
        private readonly Dictionary<T, ISet<T>> InNeighbors;
        private readonly int ExpectedNodes;

        public Graph(int expectedNodes)
        {
            ExpectedNodes = expectedNodes;
            Edges = new Dictionary<T, Dictionary<T, int>>(ExpectedNodes);
            InNeighbors = new Dictionary<T, ISet<T>>(ExpectedNodes);
        }

        public void Connect(T vertexFrom, T vertexTo, int weight)
        {
            if (!Edges.ContainsKey(vertexFrom))
            {
                Edges[vertexFrom] = new Dictionary<T, int>(ExpectedNodes);
                InNeighbors[vertexFrom] = new HashSet<T>(ExpectedNodes);
            }

            if (!Edges.ContainsKey(vertexTo))
            {
                Edges[vertexTo] = new Dictionary<T, int>(ExpectedNodes);
                InNeighbors[vertexTo] = new HashSet<T>(ExpectedNodes);
            }

            Edges[vertexFrom][vertexTo] = weight;
            InNeighbors[vertexTo].Add(vertexFrom);
        }

        public ISet<T> FindAllOutNeighbors(T vertex)
        {
            return new HashSet<T>(Edges[vertex].Keys);
        }

        public ISet<T> FindAllInNeighbors(T vertex)
        {
            return new HashSet<T>(InNeighbors[vertex]);
        }

        public ISet<T> FindAllVertices()
        {
            return new HashSet<T>(Edges.Keys);
        }

        public Dictionary<T, ISet<T>> FindConnectedComponents()
        {
            // For each vertex u of the graph, mark u as unvisited
            var unvisited = FindAllVertices();

            // Let L be empty
            var L = new LinkedList<T>();

            // For each vertex u of the graph do Visit(u)
            foreach (var vertex in FindAllVertices())
            {
                Visit(vertex, unvisited, L);
            }

            var components = new Dictionary<T, ISet<T>>();

            // For each element u of L in order, do Assign(u,u)
            foreach (var vertex in L)
            {
                Assign(vertex, vertex, components);
            }

            return components;
        }

        private void Visit(T vertex, ISet<T> unvisited, LinkedList<T> L)
        {
            // If u is unvisited then
            if (unvisited.Contains(vertex))
            {
                // Mark u as visited
                unvisited.Remove(vertex);

                // For each out-neighbour v of u, do Visit(v)
                foreach (var neighbor in FindAllOutNeighbors(vertex))
                {
                    Visit(neighbor, unvisited, L);
                }

                // Prepend u to L
                L.AddFirst(vertex);
            }
        }

        private void Assign(T vertex, T root, Dictionary<T, ISet<T>> components)
        {
            var hasBeenAssigned = false;

            foreach (var component in components.Values)
            {
                if (component.Contains(vertex))
                {
                    hasBeenAssigned = true;
                    break;
                }
            }

            // If u has not been assigned to a component then
            if (!hasBeenAssigned)
            {
                if (!components.ContainsKey(root))
                {
                    components[root] = new HashSet<T>();
                }

                // Assign u as belonging to the component whose root is root
                components[root].Add(vertex);

                // For each in-neighbour v of u, do Assign(v,root)
                foreach (var neighbor in FindAllInNeighbors(vertex))
                {
                    Assign(neighbor, root, components);
                }
            }
        }
    }
}
