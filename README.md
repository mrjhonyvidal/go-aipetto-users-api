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
    Default /usr/local/go
    export GOROOT=/usr/local/go
    export PATH=$PATH:$GOROOT/bin
GOPATH: 
    /Users/{user}go
```

Set your environment by following the instructions in https://golang.org/ and git clone following the below rules:

GOPATH: Workspace for Go projects:
- Go < 1.13: Every Go project must be cloned INSIDE of your GOPATH
- Go > 1.13: Every Go project must be cloned OUTSIDE of your GOPATH
- Go 1.13 introduces "modules".

Clean: rm -fr /Users/..workspaces/golang/src/github.com/username/project-api

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
go get -u //github.com/gin-gonic/gin (http framework Gin)
go get -u github.com/gin-gonic/gin
docker build -t main .
```

#### Know-how and using Go Modules
Dependencies in Golang(check different behaviour depend if version is < 1.13 or > 1.13).
```
1. Look depedencies in myproject-api
                            main.go
                            vendor
                                github.com
                                    username
                                        project
                                            package

2. Look in GOPATH: Workspace for Go projects
        - Go < 1.13: Every Go project must be cloned INSIDE of your GOPATH
        - Go > 1.13: Every Go project must be cloned OUTSIDE of your GOPATH
        - Go 1.13 introduces "modules".

2. Using Go Modules
    cd $GOPATH
    pkg
        mod
            github.com
                username/company

3. Look in GOROOT: Where Go is installed. Default: /usr/local/go
        export GOROOT=/usr/local/go
        export PATH=$PATH:$GOROOT/bin                     
```
In the root of the project main.go run:

```
go mod init github.com/aipetto/go-aipetto-users-api
```

### Troubleshoot
```
fatal: could not read Username for 'https://github.com': terminal prompts disabled
Confirm the import path was entered correctly.
If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.

Check that the project has a tag, otherwise add one.

git config --global --add url."git@github.com:".insteadOf "https://github.com/" ---> ~/.gitconfig check for ssh always 
git config --global credential.helper cache

go env -w GOPRIVATE=github.com/aipetto/* ---> Make golang find dependencies in private repositories

go mod tidy ---> https://blog.golang.org/using-go-modules
The go mod tidy command cleans up these unused dependencies:

go mod init github.com/aipetto/go-aipetto-users-api
go clean -modcache

Inside src run: go run main.go //go get
```

#### know-how
- https://golang.org/doc/code.html?h=modcache

Module dependencies are automatically downloaded to the pkg/mod subdirectory of the directory indicated by the GOPATH environment variable. The downloaded contents for a given version of a module are shared among all other modules that require that version, so the go command marks those files and directories as read-only. To remove all downloaded modules, you can pass the -modcache flag to go clean:
`$ go clean -modcache`

#### Adding dependencies to our project on vendor folder
```
Create Gopkg.toml file
dep init
```