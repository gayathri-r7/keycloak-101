# Todo Application

## How to Run?

### MongoDB Configuration

```sh
$ docker run --name mongodb --publish 27017:27017 -d mongo:4.0.4
```

### Backend service

```sh
$ cd todo-service
$ rm -rf go.mod go.sum vendor
$ go mod init todoservice
$ go mod tidy
$ go mod vendor
$ go run todo.go
```

### Frontend service

```sh
$ cd todo-ui
$ npm install
$ npm run serve
```

Access the todo application at http://localhost:8080/

### MongoDB Commands

```sh
$ docker exec -it <ContainerID> bash
$ mongo admin --host localhost:27017
$ use todo_db
$ show dbs
$ show tables
$ db.todos.find()
$ db.todos.insert({_id: "1", task: "Keycloak blog", description: "Create a blog post about keycloak."})
```
