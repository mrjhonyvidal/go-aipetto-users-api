### AIPETTO  🦴🐾🐈 🐕 🦮 🐻‍❄️ 😾 🐕‍🦺

### Run
Set your DB environments, in shell type:
```
MYSQL_HOST=0.0.0.0:3306
MYSQL_DATABASE=aipetto_users_db
MYSQL_ROOT_PASSWORD=GhJkLBnM1029
MYSQL_USER=aipetto01
MYSQL_PASSWORD=GhJkLBnM1029

OR

export MYSQL_HOST=0.0.0.0:3306
export MYSQL_DATABASE=aipetto_users_db
export MYSQL_ROOT_PASSWORD=
export MYSQL_USER=aipetto01
export MYSQL_PASSWORD=
```

Check exported variables of your system:
```
export -p
export -p | grep  something
export -p | less
export -p | more
```

```
sudo docker-compose up --build
sudo docker-compose up -d (daemon mode)
sudo docker-compose up --remove-orphans


```
The result should be something like this:
```
aipetto-users-mysql     | 2021-01-18T23:21:30.930957Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '8.0.22'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server - GPL.
go-users-api            | 2021/01/18 23:21:32 Database successfully configured
go-users-api            | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
go-users-api            |
go-users-api            | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
go-users-api            |  - using env: export GIN_MODE=release
go-users-api            |  - using code:        gin.SetMode(gin.ReleaseMode)
go-users-api            |
go-users-api            | [GIN-debug] GET    /ping                     --> github.com/aipetto/go-aipetto-users-api/src/controllers/ping.Ping (3 handlers)
go-users-api            | [GIN-debug] GET    /users/:user_id           --> github.com/aipetto/go-aipetto-users-api/src/controllers/users.GetUser (3 handlers)
go-users-api            | [GIN-debug] POST   /users                    --> github.com/aipetto/go-aipetto-users-api/src/controllers/users.CreateUser (3 handlers)
go-users-api            | [GIN-debug] Listening and serving HTTP on :8081
go-users-api            | [GIN] 2021/01/18 - 23:21:39 | 200 |      13.708µs |      172.22.0.1 | GET      "/ping"
go-users-api            | [GIN] 2021/01/18 - 23:21:45 | 201 |     178.364µs |      172.22.0.1 | POST     "/users"
go-users-api            | [GIN] 2021/01/18 - 23:21:48 | 200 |     114.499µs |      172.22.0.1 | GET      "/users/123"
go-users-api            | [GIN] 2021/01/18 - 23:21:52 | 404 |     152.793µs |      172.22.0.1 | GET      "/users/456"
```

Check containers with `docker ps`:
```
4a686766c72a   go-aipetto-users-api_go-users-service   "./go-users-api"         35 seconds ago   Up 34 seconds                   0.0.0.0:8081->8081/tcp              go-users-api
4a80fefac883   mysql:latest                            "docker-entrypoint.s…"   35 seconds ago   Up 34 seconds                   0.0.0.0:3306->3306/tcp, 33060/tcp   mysql-container
```

### Run Dockerfile individually
```
cd client
docker build -f Dockerfile.dev .
docker run -it -p 3000:3000
```

### Running our Docker container
```
go get -u github.com/gin-gonic/gin
go run main.go or Run on main.go directly from IDE.
sudo docker run -p 8081:8081 users-api:latest
sudo docker run -p 8081:8081 -p 9200:9200 (ElasticSearch) users-api:latest
docker run --name aipetto-mysql -e MYSQL_ROOT_PASSWORD=PASSWORD -d mysql:latest
docker run -p 3306:3306 aipetto-users-mysql (must add the Dockerfile only for Mysql - work in progress)
docker run user_service
```

### DB
```
docker exec -it aipetto-mysql mysql -uroot -p
sudo docker logs (check password installed for first time)
```

### Users API
  
Check service
 ```
curl -X GET localhost:8080/ping -v
curl -X GET localhost:8080/users/123 -v
curl -X POST localhost:8080/users -d '{"id":123, "first_name": "GoPetto", "email": "go@aipetto.com"}' -v
```

#### Getting Access Token from our OAuth Go Service when Public APIs
 ```
 /users/1?access_token_id=241dsad124121d21d2141
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

#### GO know-how

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
github.com/go-sql-driver/mysql
docker build -t main .
```

#### Using Go Module
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

Private repos access, example: go: github.com/aipetto/go-aipetto-oauth-library@v0.2.1: reading github.com/aipetto/go-aipetto-oauth-library/go.mod at revision v0.2.1: unknown revision v0.2.1

go env -w GO111MODULE=on
go help module-private
go env -w GOPRIVATE=github.com/aipetto/*
git config --global url."ssh://git@github.com:aipetto".insteadOf "https://github.com/aipetto"
git config --global url."https://:x-oauth-basic@github.com:aipetto".insteadOf "https://github.com/aipetto"
```


### Extra
- https://golang.org/doc/code.html?h=modcache

Module dependencies are automatically downloaded to the pkg/mod subdirectory of the directory indicated by the GOPATH environment variable. The downloaded contents for a given version of a module are shared among all other modules that require that version, so the go command marks those files and directories as read-only. To remove all downloaded modules, you can pass the -modcache flag to go clean:
`$ go clean -modcache`

#### Adding dependencies to our project on vendor folder
```
Create Gopkg.toml file
dep init
```


###DB MySQL General Troubleshoot Log 

Not accessing from MySQL workbench or other resource or container, add or uncoment bind-address from mysql configs:
```
RUN sed -Ei 's/^(bind-address|log)/#&/' /etc/mysql/my.cnf
sed -i -e "/^bind-address/d" /etc/mysql/my.cnf
sed -ie "s/^bind-address\s*=\s*127\.0\.0\.1$/#bind-address = 0.0.0.0/"

In my case the folder where it was located inside the container was /etc/my.cnf and the property bind-address did not exist.
So I added manually below the mysqld and copy from my local machine after had copied from the container.
sudo docker cp my.cnf aipetto-mysql:/etc/my.cnf
sudo docker aipetto-mysql stop
sudo docker aipetto-mysql start
```

#### Considering / TODO
Should we add a explicit network between this two containers? Let's check Kubernetes and overall infra.
```
networks:
  aipettonet:
    driver: bridge
    
In each service add:
    expose:
     ...
    networks:
      - aipettonet
    volumes_from:
      ...    
```

#### Crypo Bcrypt Know-how
A Bcrypt hash can be stored in a BINARY(40) column.

BINARY(60), as the other answers suggest, is the easiest and most natural choice, but if you want to maximize storage efficiency, you can save 20 bytes by losslessly deconstructing the hash. I've documented this more thoroughly on GitHub: https://github.com/ademarre/binary-mcf

Bcrypt hashes follow a structure referred to as modular crypt format (MCF). Binary MCF (BMCF) decodes these textual hash representations to a more compact binary structure. In the case of Bcrypt, the resulting binary hash is 40 bytes.

Gumbo did a nice job of explaining the four components of a Bcrypt MCF hash:

$<id>$<cost>$<salt><digest>
Decoding to BMCF goes like this:

$<id>$ can be represented in 3 bits.
<cost>$, 04-31, can be represented in 5 bits. Put these together for 1 byte.
The 22-character salt is a (non-standard) base-64 representation of 128 bits. Base-64 decoding yields 16 bytes.
The 31-character hash digest can be base-64 decoded to 23 bytes.
Put it all together for 40 bytes: 1 + 16 + 23

BCrypt generates an implementation-dependent 448-bit hash value. You might need CHAR(56), CHAR(60), CHAR(76), BINARY(56) or BINARY(60)