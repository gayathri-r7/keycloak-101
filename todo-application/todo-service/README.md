# Todo-Service

## How to Run?

### MongoDB Configuration

```sh
$ docker run --name mongodb --publish 27017:27017 -d mongo:4.0.4
```

### Go Configuration

```sh
$ cd todo-service
$ rm -rf go.mod go.sum vendor
$ go mod init todoservice
$ go mod tidy
$ go mod vendor
$ go run todo.go
```

#### To check whether the todos are stored in the database:-

```sh
$ docker exec -it <ContainerID> bash
$ mongo admin --host localhost:27017
$ use todo_db
$ show dbs
$ show tables
$ db.todos.find()
$ db.todos.insert({_id: "1", task: "Keycloak blog", description: "Create a blog post about keycloak."})
```
