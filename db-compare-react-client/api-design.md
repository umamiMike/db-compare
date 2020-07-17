## NOTES
<!-- Look for nouns and verbs in the requirement -->
This document assumes the use of HTTP.

## list of resources

### datasource

a node of queryable information.I am using datasource as an abstract concept. The mvp will contain only a postgres db and a csv file


### POST /query

request for (usually) filtered data from a [datasource](#datasource)


TODO: best simple way to create a user, I suppose a log in screen, but I would like to keep it simple

### GET /user?uuid

### PUT /user: Create a new user

```
  {

  }
```

create a new user

# endpoints 

### GET /user?name=
returns a single datasource 

### GET /users?filter=

types of filters
```
{
"data": {
value: 
}

}

```

### DELETE /user/{user-id}/{[datasources](#datasource)}

### request

{user-id}

### response
```

{
data: "success"
}

```
a datasource, which is a file
db-datasource

----

### POST /datasource

when a user makes a query they need to provide appropriate datasource connection information to the server in order to complete the query.

### request

```json
{ "datasource": {
    "username": "user", 
    "hostname": "localhost", 
    "password": "passwd", 
    "port": "5432"
    "db-name": "example"
  }
  "name": "local-postgres"
}

```

### response:

```javascript

{
"data": {
  "status"
  "id": "3"
  }
}
```

## HTTP POST /datasource/[query](#query)/
Question: would this maybe be  a post or a put since it is a potentially NEW piece of data being returned.
I think it will be post because the same query at different times could be return different responses.  (NOT IDEMPOTENT)


### request

### response


