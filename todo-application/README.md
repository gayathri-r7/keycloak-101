# Todo Application

A to-do list application is an application that allows users to create lists of tasks that can be sorted, edited, tracked, and completed. 

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

### Rest API calls

```sh
POST - http://localhost:3000/todo
REQUEST BODY:
{
"task": "Kecycloak blog",
"description": "Create a blog post about keycloak."
}
RESPONSE CODE: 201
RESPONSE:
{
  "message": "Todo created successfully.",
  "todoId": "tsk209f541be1a84973a535f0ad19574558"
}
---------------------------------------------------

GET - http://localhost:3000/todo/<TaskID>
RESPONSE CODE: 200
RESPONSE:
{
  "id": "tsk209f541be1a84973a535f0ad19574558",
  "task": "Kecycloak blog",
  "description": "Create a blog post about keycloak.",
  "status": "In-Progress"
}
---------------------------------------------------

PUT - http://localhost:3000/todo/<TaskID>
REQUEST BODY:
{
"task": "Kecycloak blog",
"description": "Create a blog post about keycloak.",
"status": "Completed"
}
RESPONSE CODE: 200
RESPONSE:
{
  "message": "Todo updated successfully."
}
---------------------------------------------------

GET - http://localhost:3000/todos
RESPONSE CODE: 200
RESPONSE:
[
    {
        "id": "tsk209f541be1a84973a535f0ad19574558",
        "task": "Kecycloak blog",
        "description": "Create a blog post about keycloak.",
        "status": "In-Progress"
    }
]
---------------------------------------------------

DELETE - http://localhost:3000/todo/<TaskID>
RESPONSE CODE: 200
RESPONSE:
{
  "message": "Todos deleted successfully"
}
---------------------------------------------------
```

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
