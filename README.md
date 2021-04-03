[![GitHub license](https://img.shields.io/github/license/abhijitWakchaure/mock-api-server-go?style=for-the-badge)](https://github.com/abhijitWakchaure/mock-api-server-go/blob/master/LICENSE) 
![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/abhijitwakchaure/mock-api-server-go?style=for-the-badge)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/abhijitWakchaure/mock-api-server-go?style=for-the-badge)

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

## Exposed APIs (Default)
| Method |   Path          |
|:-------|:----------------|
| GET    | /api/users      |
| POST   | /api/users      |
| GET    | /api/users/{id} |
| PUT    | /api/users/{id} |
| DELETE | /api/users/{id} |

## Flags
| Flag      | Default Value(s) | Description              |
| :---------|:-----------------|:-------------------------|
| -endpoint | /api/users/      | Endpoint for API server  |
| -port     | 8080             | Port for the API server  |
| -version  | NA               | Print the server version |
| -help     | NA               | Print the help message   |


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