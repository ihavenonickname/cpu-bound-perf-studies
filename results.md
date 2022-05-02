# Results

In every test 36 identical requests were sent.

System information:

```none
$ wmic os get version
Version
10.0.19044

$ wmic cpu get name
Name
AMD Ryzen 5 3600 6-Core Processor

$ wmic memorychip get capacity, speed
Capacity    Speed
8589934592  2667
8589934592  2667
```

## .net

Version:

```none
$ dotnet --version
6.0.100
```

Compiled with:

```none
$ dotnet publish --configuration Release --runtime win-x64
```

### 1 worker

* Total elapsed: 42607 ms
* Average response time: 1182 ms
* Throughput: 50 requests/minute (1183 ms)

### 2 workers

* Total elapsed: 25762 ms
* Average response time: 1430 ms
* Throughput: 83 requests/minute (715 ms)

### 3 workers

* Total elapsed: 20225 ms
* Average response time: 1672 ms
* Throughput: 106 requests/minute (561 ms)

### 4 workers

* Total elapsed: 17553 ms
* Average response time: 1930 ms
* Throughput: 123 requests/minute (487 ms)

### 5 workers

* Total elapsed: 16754 ms
* Average response time: 2177 ms
* Throughput: 129 requests/minute (465 ms)

### 6 workers

* Total elapsed: 15256 ms
* Average response time: 2475 ms
* Throughput: 141 requests/minute (423 ms)

## Go

Version:

```none
$ go version
go version go1.17.3 windows/amd64
```

Compiled with:

```none
$ go build -ldflags "-s -w"
```

### 1 worker

* Total elapsed: 131196 ms
* Average response time: 3643 ms
* Throughput: 16 requests/minute (3644 ms)

### 2 workers

* Total elapsed: 72450 ms
* Average response time: 4024 ms
* Throughput: 29 requests/minute (2012 ms)

### 3 workers

* Total elapsed: 53485 ms
* Average response time: 4437 ms
* Throughput: 40 requests/minute (1485 ms)

### 4 workers

* Total elapsed: 42949 ms
* Average response time: 4739 ms
* Throughput: 50 requests/minute (1193 ms)

### 5 workers

* Total elapsed: 36117 ms
* Average response time: 4976 ms
* Throughput: 59 requests/minute (1003 ms)

### 6 workers

* Total elapsed: 33574 ms
* Average response time: 5265 ms
* Throughput: 64 requests/minute (932 ms)

## node.js

Version:

```none
$ node --version
v16.13.2
```

### 1 worker

* Total elapsed: 75236 ms
* Average response time: 2089 ms
* Throughput: 28 requests/minute (2089 ms)

### 2 workers

* Total elapsed: 67981 ms
* Average response time: 3722 ms
* Throughput: 31 requests/minute (1888 ms)

### 3 workers

* Total elapsed: 68394 ms
* Average response time: 5532 ms
* Throughput: 31 requests/minute (1899 ms)

### 4 workers

* Total elapsed: 68743 ms
* Average response time: 7310 ms
* Throughput: 31 requests/minute (1909 ms)

### 5 workers

* Total elapsed: 68528 ms
* Average response time: 8959 ms
* Throughput: 31 requests/minute (1903 ms)

### 6 workers

* Total elapsed: 68917 ms
* Average response time: 10648 ms
* Throughput: 31 requests/minute 1914 ms
