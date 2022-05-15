This is a study on how different programming environments handle heavy CPU work.

https://en.wikipedia.org/wiki/Kosaraju%27s_algorithm

Some snippets to help me to run the tests.

```none
python test-single.thread.py --url "http://localhost:5000/" --max-workers 2 --requests 18
python test-single.thread.py --url "http://localhost:3000/" --max-workers 2 --requests 18
python test-single.thread.py --url "http://localhost:8080/" --max-workers 2 --requests 18
```

```json
[
  [ 0, 1, 0, 0, 0, 0, 0, 0 ],
  [ 0, 0, 1, 1, 0, 0, 0, 0 ],
  [ 1, 0, 0, 1, 0, 0, 0, 0 ],
  [ 0, 0, 0, 0, 1, 0, 0, 0 ],
  [ 0, 0, 0, 1, 0, 0, 0, 0 ],
  [ 0, 0, 0, 0, 1, 0, 1, 0 ],
  [ 0, 0, 0, 0, 0, 1, 0, 1 ],
  [ 0, 0, 0, 0, 1, 0, 1, 0 ]
]
```

```csharp
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
```
