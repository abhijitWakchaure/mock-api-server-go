# mock-api-server-go

Simple mock API server in go

## Flags
| Flag  | Default Value(s) | Description             |
| :---- |:-----------------|:------------------------|
| -p    | 8080             | Port for the API server |
| -h    | NA               | Print the help message  |



## How to Run the server

### Docker
```
docker run -it -p 8080:8080 abhijitwakchaure/mock-api-server-go
```

### Linux or Mac
Download the binary from here [**mock-api-server-go-v1.0.0-linux-amd64**](https://github.com/abhijitWakchaure/mock-api-server-go/releases/download/v1.0.0/mock-api-server-go-v1.0.0-linux-amd64) for linux or from here [**mock-api-server-go-v1.0.0-darwin-amd64**](https://github.com/abhijitWakchaure/mock-api-server-go/releases/download/v1.0.0/mock-api-server-go-v1.0.0-darwin-amd64) for mac

Run the server with default port 8080

```
./mock-api-server-go-v1.0.0-linux-amd64
```

Specify the port (e.g. 9000)

```
./mock-api-server-go-v1.0.0-linux-amd64 -p 9000
```

### Windows
Download the binary from here [**mock-api-server-go-v1.0.0-win-amd64.exe**](https://github.com/abhijitWakchaure/mock-api-server-go/releases/download/v1.0.0/mock-api-server-go-v1.0.0-win-amd64.exe)
