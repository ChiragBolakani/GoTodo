# Todo API
TODO API using Golang and ScyllaDB

Supports:
- CRUD Operations
- filtering : pending/completed
- Time Sorted
- Pagination

## Getting Started

To get a local copy up and running follow these simple example steps.

### Prerequisites

This project requires the following to be installed
* Docker
* Go

### Installation

1. pull ScyllaDB docker image
   ```sh
   docker pull scylladb/scylla
   ```
2. Use the docker run command to start a container
   ```sh
   docker run --name some-scylla -p 9042:9042 --hostname some-scylla -d scylladb/scylla --smp 1
   ```
   **_NOTE:_**  This command will start a Scylla single-node cluster in developer mode limited by a single CPU core.
3. Clone this repository
   ```sh
   git clone https://github.com/ChiragBolakani/GoTodo.git
   ```
4. Create keyspace and tables
   Enter into the container cqlsh
   ```sh
   docker exec -it <container_tag> cqlsh
   ```

   Run the .cql file
   ```sh
   SOURCE /path/to/pkg/db/schema/schema.cql
   ```

   OR copy paste the commands from schema.cql file inside your docker container
   
4. Configure .env

   ```env
   SERVER_PORT=
   
   SCYLLA_DB_HOSTS=
   SCYLLA_DB_DATABASE=
   SCYLLA_DB_USERNAME=
   SCYLLA_DB_PASSWORD=
   ```
5. Run the project
   ```sh
   go run .\cmd\server\main.go
   ```

## Schema 

```cql
CREATE KEYSPACE todo WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;

CREATE TABLE todo.users (
    user_id timeuuid PRIMARY KEY,
    first_name text,
    last_name text
);

CREATE TABLE todo.items (
    user_id timeuuid,
    id timeuuid,
    created_at timestamp,
    description text,
    status boolean,
    title text,
    updated_at timestamp,
    PRIMARY KEY ((user_id), id)
);
```

## API Documentation
The API provides the following routes:
| Method   | URL                                      | Description                              |
| -------- | ---------------------------------------- | ---------------------------------------- |
| `GET`    | `/api/v1/users/{user_id}/items`                             |get todo items.                      |
| `GET`    | `/api/v1/users/{user_id}/items/{item_id}`                             |get single todo item.                      |
| `POST`    | `/api/v1/items`                          | create a new todo item.                       |
| `PUT`    | `/api/v1/items/{itme_id}`                          | update todo item.                       |
| `DELETE`    | `/api/v1/users/{user_id}/items/{item_id}`                          | delete todo item.                       |

## Usage
### Create new todo item
`POST /api/v1/items`

request : 
```http
POST /api/v1/items HTTP/1.1
Content-Type: application/json
User-Agent: PostmanRuntime/7.39.0
Accept: */*
Cache-Control: no-cache
Postman-Token: 115eaad4-6e96-4e9b-9a09-370ab9ef554d
Host: localhost:8000
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Content-Length: 231
 
{
  "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
  "title": "Dishes",
  "description": "Do the dishes by 7pm",
  "status": false,
  "created_at": "2024-07-02T00:53:55.632Z",
  "updated_at": "2024-07-02T00:53:55.632Z"
}
```
response :
```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 03 Jul 2024 11:38:31 GMT
Content-Length: 17
 
{
  "success":true
}
```

### Get single todo item
`GET /api/v1/users/0bf7a0b1-371a-11ef-a6ea-ff343fe64132/items/1888ee91-38e8-11ef-8cf3-ff303fe64132`

request : 
```http
GET /api/v1/users/0bf7a0b1-371a-11ef-a6ea-ff343fe64132/items/c5d37dc1-3930-11ef-8cf3-ff303fe64132 HTTP/1.1
User-Agent: PostmanRuntime/7.39.0
Accept: */*
Cache-Control: no-cache
Postman-Token: 6101389e-a825-431b-8e51-21520eaff907
Host: localhost:8000
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
```

response :
```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 03 Jul 2024 12:06:21 GMT
Content-Length: 268
 
{
    "data": {
        "id": "c5d37dc1-3930-11ef-8cf3-ff303fe64132",
        "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
        "title": "Dishes",
        "description": "Do the dishes by 7pm",
        "status": false,
        "created_at": "2024-07-03T11:38:31.165Z",
        "updated_at": "2024-07-03T11:38:31.165Z"
    },
    "success": true
}
```

### Get all todo items of a user
`GET /api/v1/users/0bf7a0b1-371a-11ef-a6ea-ff343fe64132/items`

- Pagination Supported : The url allows a query parameter `page_state` that holds the pageStateToken required to retrieve the next page of records.
- Example :
  ```http
  GET /api/v1/users/0bf7a0b1-371a-11ef-a6ea-ff343fe64132/items?page_state=00000000b20000001c00000001000000100000000bf7a0b1371a11efa6eaff343fe64132011c000000010000001000000080dc2271380f11ef8cf3ff303fe64132f5ffffff444199d86ac03ecb6d69b66408bd8793010000003900000001190000001400000001000000080000000532e00ef808febd0101190000001400000001000000080000000532e00ef808febd010101000000c341dbc08068c344a1dbabcc5534fdbe010000000000ffffffff000000000002
  ```

request : 
```http
GET /api/v1/users/0bf7a0b1-371a-11ef-a6ea-ff343fe64132/items HTTP/1.1
User-Agent: PostmanRuntime/7.39.0
Accept: */*
Cache-Control: no-cache
Postman-Token: acb8b719-224d-4fd7-adbe-167655888784
Host: localhost:8000
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
```

response :
```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 03 Jul 2024 03:12:24 GMT
Transfer-Encoding: chunked

{
  "data": [
    {
      "id": "0bf7a0b1-371a-11ff-a6ea-ff343fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "",
      "description": "updated",
      "status": true,
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "2024-07-03T00:26:30.908Z"
    },
    {
      "id": "1888ee91-38e8-11ef-8cf3-ff303fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "Dishes",
      "description": "Do the dishes by 7pm",
      "status": false,
      "created_at": "2024-07-03T02:58:16.683Z",
      "updated_at": "2024-07-03T02:58:16.683Z"
    },
    {
      "id": "f7075b31-38e7-11ef-8cf3-ff303fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "",
      "description": "tesdadadadat",
      "status": true,
      "created_at": "2024-07-03T02:57:20.471Z",
      "updated_at": "2024-07-03T02:57:20.471Z"
    },
    {
      "id": "e5853d51-38e7-11ef-8cf3-ff303fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "",
      "description": "tesdadadadat",
      "status": true,
      "created_at": "2024-07-03T02:56:51.09Z",
      "updated_at": "2024-07-03T02:56:51.09Z"
    },
    {
      "id": "978c8421-38d6-11ef-8cf3-ff303fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "",
      "description": "tesdadadadat",
      "status": true,
      "created_at": "2024-07-03T00:52:58.836Z",
      "updated_at": "2024-07-03T00:52:58.836Z"
    },
    {
      "id": "537b6ca1-38d2-11ef-8cf3-ff303fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "",
      "description": "tesdadadadat",
      "status": true,
      "created_at": "2024-07-03T00:22:26.637Z",
      "updated_at": "2024-07-03T00:22:26.637Z"
    },
    {
      "id": "44db0961-38bb-11ef-8cf3-ff303fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "task ",
      "description": "tesdadadadat",
      "status": true,
      "created_at": "2024-07-02T21:37:23.698Z",
      "updated_at": "2024-07-02T21:37:23.698Z"
    },
    {
      "id": "43b8a971-38bb-11ef-8cf3-ff303fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "task ",
      "description": "tesdadadadat",
      "status": true,
      "created_at": "2024-07-02T21:37:21.788Z",
      "updated_at": "2024-07-02T21:37:21.788Z"
    },
    {
      "id": "32a775d1-38bb-11ef-8cf3-ff303fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "task ",
      "description": "tesdadadadat",
      "status": true,
      "created_at": "2024-07-02T21:36:53.132Z",
      "updated_at": "2024-07-02T21:36:53.132Z"
    },
    {
      "id": "80dc2271-380f-11ef-8cf3-ff303fe64132",
      "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
      "title": "task ",
      "description": "tesdadadadat",
      "status": true,
      "created_at": "2024-07-02T01:07:50.915Z",
      "updated_at": "2024-07-02T01:07:50.915Z"
    }
  ],
  "next_page_state": "00000000b20000001c00000001000000100000000bf7a0b1371a11efa6eaff343fe64132011c000000010000001000000080dc2271380f11ef8cf3ff303fe64132f5ffffff444199d86ac03ecb6d69b66408bd8793010000003900000001190000001400000001000000080000000532e00ef808febd0101190000001400000001000000080000000532e00ef808febd010101000000c341dbc08068c344a1dbabcc5534fdbe010000000000ffffffff000000000002",
  "success": true
}
```

### Update todo item
`PUT /api/v1/items/0bf7a0b1-371a-11ef-a6ea-ff343fe64132/items/`

request :
```http
PUT /api/v1/items/c5d37dc1-3930-11ef-8cf3-ff303fe64132 HTTP/1.1
Content-Type: application/json
User-Agent: PostmanRuntime/7.39.0
Accept: */*
Cache-Control: no-cache
Postman-Token: d2e4c030-f30d-43d5-9479-d4868c761726
Host: localhost:8000
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Content-Length: 294
 
{
    "id": "c5d37dc1-3930-11ef-8cf3-ff303fe64132",
    "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
    "title": "Dishes",
    "description": "Do the dishes by 7pm",
    "status": false,
    "created_at": "2024-07-03T11:38:31.165Z",
    "updated_at": "2024-07-03T11:38:31.165Z"
}
```

response : 
```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 03 Jul 2024 12:17:18 GMT
Content-Length: 44

{
    "data": {
        "id": "c5d37dc1-3930-11ef-8cf3-ff303fe64132",
        "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
        "title": "Dishes",
        "description": "Do the dishes by 7pm",
        "status": false,
        "created_at": "2024-07-03T11:38:31.165Z",
        "updated_at": "2024-07-03T11:38:31.165Z"
    },
    "success": true
}
```

### Delete todo item
`DELETE /api/v1/users/0bf7a0b1-371a-11ef-a6ea-ff343fe64132/items/c5d37dc1-3930-11ef-8cf3-ff303fe64132`

request :
```http
DELETE /api/v1/users/0bf7a0b1-371a-11ef-a6ea-ff343fe64132/items/c5d37dc1-3930-11ef-8cf3-ff303fe64132 HTTP/1.1
User-Agent: PostmanRuntime/7.39.0
Accept: */*
Cache-Control: no-cache
Postman-Token: 0cd86295-d63d-4d13-ab8c-94582f235d29
Host: localhost:8000
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
```

response : 
```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 03 Jul 2024 12:21:27 GMT
Content-Length: 268
 
{
    "data": {
        "id": "c5d37dc1-3930-11ef-8cf3-ff303fe64132",
        "user_id": "0bf7a0b1-371a-11ef-a6ea-ff343fe64132",
        "title": "Dishes",
        "description": "Do the dishes by 7pm",
        "status": false,
        "created_at": "2024-07-03T11:38:31.165Z",
        "updated_at": "2024-07-03T11:38:31.165Z"
    },
    "success": true
}
```
