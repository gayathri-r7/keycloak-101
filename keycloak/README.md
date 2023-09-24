# Keycloak setup

```sh
$ docker network create app-network

$ docker run -d --name keycloak -p 8088:8080 -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=admin --net app-network quay.io/keycloak/keycloak:22.0.1 start-dev
```

Access the Keycloak server's admin console via http://localhost:8088/

## Reference:-

1. Enabling the user profile - https://www.keycloak.org/docs/latest/server_admin/#enabling-the-user-profile
