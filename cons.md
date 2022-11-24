### cmd - entry point for our application

### internal - private application and library code (handler, service, etc.)

### pkg - library code that's ok to use by external applications (client to db etc, utils, loggers)


### REST - (cache, stateless, http)


##### 200 0k, 204 - no content, 400 - bad request, 404 - not found, 500 - internal server error
GET /api/users - list of users -- 200, 404, 500 
GET /api/users/1 - user with id 1 -- 200, 404, 500
POST /api/users/:id - create user with id 1 -- 204, 4xx, Header Local: url
PUT /api/users/:id - update user with id 1 (fully) -- 204, 4xx, Header Local: url
PATCH /api/users/:id - update user with id 1 (partially) -- 204, 4xx, Header Local: url
DELETE /api/users/:id - delete user with id 1 -- 204, 4xx, Header Local: url