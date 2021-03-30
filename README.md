# mock-api-server-go

Simple mock API server in go

## Schema for user struct
```
{
	"id": "60624180893d170927d32e27",
	"username": "john@example.com",
	"password": "EQWMJYq40spmT#g",
	"fullname": "John Doe",
	"mobile": "+91 9999999999",
	"createdAt": 1538919475135,
	"modifiedAt": 1599340945571,
	"blocked": false,
	"roles": [
		"ROLE_USER"
	]
}
```

## Exposed APIs
| Method |   Path      |
|:-------|:-----------:|
| GET    | /users      |
| POST   | /users      |
| GET    | /users/{id} |
| PUT    | /users/{id} |
| DELETE | /users/{id} |

## Flags
| Flag  | Default Value(s) | Description             |
| :---- |:-----------------|:------------------------|
| -p    | 8080             | Port for the API server |
| -h    | NA               | Print the help message  |


## Download the API server

### Docker
```
docker pull abhijitwakchaure/mock-api-server-go
```

### Linux, Windows or Mac
Download the latest binary from [**here**](https://github.com/abhijitWakchaure/mock-api-server-go/releases/latest) 


## How to Run the server

### Docker
```
docker run -it -p 8080:8080 abhijitwakchaure/mock-api-server-go
```

### Linux or Mac
Run the server with default port 8080

```
./mock-api-server-go-v1.0.0-linux-amd64
```

Specify the port (e.g. 9000)

```
./mock-api-server-go-v1.0.0-linux-amd64 -p 9000
```

### Windows
Simply run the exe file