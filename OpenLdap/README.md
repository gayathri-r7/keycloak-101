# OpenLDAP

OpenLDAP is the open-source solution for LDAP (Lightweight Directory Access Protocol). It is a protocol used to store and retrieve data from a hierarchical directory structure such as in databases.

```sh
$ docker network create app-network

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
```

## Add users into OpenLDAP

```sh
$ docker exec -it <ContainerID> bash

# Add users
$ ldapadd -H ldap://localhost:1389 -x -W -D "cn=admin,dc=example,dc=com" -f users.ldif

# Search
ldapsearch -H ldap://localhost:1389 -x -D "cn=admin,dc=example,dc=com" -w admin -b "dc=example,dc=com"
```

## OpenLDAP keycloak configuration

```sh
-------------------------------------------------------------------------
Open LDAP Keycloak Configuration:-
-----------------------------------
Vendor - Other
Connection URL - ldap://openldap:1389
Bind DN - cn=admin,dc=example,dc=com
Bind credentials - admin

Edit mode - UNSYNCED
Users DN - ou=dev,dc=example,dc=com
Username LDAP attribute - cn
RDN LDAP attribute - cn
UUID LDAP attribute - uid
User object classes - inetOrgPerson
Trust email - Enable
-------------------------------------------------------------------------
```
