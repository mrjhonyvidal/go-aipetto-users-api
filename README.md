### AIPETTO  🦴🐾🐈 🐕 🦮 🐻‍❄️ 😾 🐕‍🦺

### Run
```
go get -u github.com/gin-gonic/gin
go run main.go or Run on main.go directly from IDE.
```

### Users API
  
Check service
 ```
curl -X GET localhost:8080/ping -v
curl -X GET localhost:8080/users/123 -v
curl -X POST localhost:8080/users -d '{"id":123, "first_name": "GoPetto", "email": "go@aipetto.com"}' -v
```

Example of console enabled
```
[GIN-debug] GET    /ping                     --> github.com/aipetto/go-aipetto-users-api/controllers/ping.Ping (3 handlers)
[GIN-debug] GET    /users/:user_id           --> github.com/aipetto/go-aipetto-users-api/controllers/users.GetUser (3 handlers)
[GIN-debug] POST   /users                    --> github.com/aipetto/go-aipetto-users-api/controllers/users.CreateUser (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2020/11/14 - 23:35:40 | 404 |     123.153µs |             ::1 | GET      "/users/123"
[GIN] 2020/11/14 - 23:36:22 | 201 |     223.671µs |             ::1 | POST     "/users"
[GIN] 2020/11/14 - 23:36:25 | 200 |      14.335µs |             ::1 | GET      "/users/123"
```

#### Development

```
GOROOT: where golang is installed. 
    /usr/local/go
GOPATH: 
    /Users/{user}go
```

Set your environment by following the instructions in https://golang.org/ and git clone our services into GOPATH:
```
Example: /home/jvo/go/src/github.com/aipetto/go-aipetto-users-api
GOPATH    
/src
    github.com
        username
            repo
```

Get dependencies:
```
http framework Gin
go get -u github.com/gin-gonic/gin
```