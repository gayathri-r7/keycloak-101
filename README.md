# keycloak-101

Keycloak is an open source software product to allow single sign-on with identity and access management aimed at modern applications and services.

## Docker network

To make it easy to connect Keycloak to OpenLDAP create a user defined network:

```sh
$ docker network create app-network
```

## Keycloak server

```sh
docker run -d --name keycloak -p 8088:8080 -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=admin --net app-network quay.io/keycloak/keycloak:22.0.1 start-dev
```

Access the Keycloak server's admin console via http://localhost:8088/

## OpenLDAP server

```sh
$ docker run --detach --rm --name openldap \
  --user 0 \
  --net app-network \
  --publish 1389:1389 \
  --publish 1636:1636 \
  --env LDAP_ROOT=dc=example,dc=com \
  --env LDAP_CONFIG_ADMIN_PASSWORD=admin \
  --env LDAP_ADMIN_USERNAME=admin \
  --env LDAP_ADMIN_PASSWORD=admin \
  bitnami/openldap:2.6.4

$ docker exec -it <ContainerID> bash

# Add users
$ ldapadd -H ldap://localhost:1389 -x -W -D "cn=admin,dc=example,dc=com" -f users.ldif

# Search
ldapsearch -H ldap://localhost:1389 -x -D "cn=admin,dc=example,dc=com" -w admin -b "dc=example,dc=com"
```

## Todo application

A to-do list application is an application that allows users to create lists of tasks that can be sorted, edited, tracked, and completed. 

### MongoDB Configuration

```sh
$ docker run --name mongodb --publish 27017:27017 -d mongo:4.0.4
```

### Backend service

```sh
$ cd todo-application/todo-service
$ rm -rf go.mod go.sum vendor
$ go mod init todoservice
$ go mod tidy
$ go mod vendor
$ go run todo.go
```

### Frontend service

```sh
$ cd todo-application/todo-ui
$ npm install
$ npm run serve
```

Access the todo application at http://localhost:8080/
