# Results

## .net

Version:

```none
6.0.100
```

Compiled with:

```none
dotnet publish --configuration Release --runtime win-x64
```

### 1 worker

* Average per request: 1182 ms
* Total elapsed: 42607 ms
* Total throuput: 1183 ms

### 2 workers

* Average per request: 1430 ms
* Total elapsed: 25762 ms
* Total throuput: 715 ms

### 3 workers

* Average per request: 1672 ms
* Total elapsed: 20225 ms
* Total throuput: 561 ms

### 4 workers

* Average per request: 1930 ms
* Total elapsed: 17553 ms
* Total throuput: 487 ms

### 5 workers

* Average per request: 2177 ms
* Total elapsed: 16754 ms
* Total throuput: 465 ms

### 6 workers

* Average per request: 2475 ms
* Total elapsed: 15256 ms
* Total throuput: 423 ms

## Go

Version:

```none
go version go1.17.3 windows/amd64
```

Compiled with:

```none
go build -ldflags "-s -w"
```

### 1 worker

* Average per request: 3643 ms
* Total elapsed: 131196 ms
* Total throuput: 3644 ms

### 2 workers

* Average per request: 4024 ms
* Total elapsed: 72450 ms
* Total throuput: 2012 ms

### 3 workers

* Average per request: 4437 ms
* Total elapsed: 53485 ms
* Total throuput: 1485 ms

### 4 workers

* Average per request: 4739 ms
* Total elapsed: 42949 ms
* Total throuput: 1193 ms

### 5 workers

* Average per request: 4976 ms
* Total elapsed: 36117 ms
* Total throuput: 1003 ms

### 6 workers

* Average per request: 5265 ms
* Total elapsed: 33574 ms
* Total throuput: 932 ms

## node.js

Version:

```none
v16.13.2
```

### 1 worker

* Average per request: 2089 ms
* Total elapsed: 75236 ms
* Total throuput: 2089 ms

### 2 workers

* Average per request: 3722 ms
* Total elapsed: 67981 ms
* Total throuput: 1888 ms

### 3 workers

* Average per request: 5532 ms
* Total elapsed: 68394 ms
* Total throuput: 1899 ms

### 4 workers

* Average per request: 7310 ms
* Total elapsed: 68743 ms
* Total throuput: 1909 ms

### 5 workers

* Average per request: 8959 ms
* Total elapsed: 68528 ms
* Total throuput: 1903 ms

### 6 workers

* Average per request: 10648 ms
* Total elapsed: 68917 ms
* Total throuput: 1914 ms
