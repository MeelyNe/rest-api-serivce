### cmd - entry point for our application

### internal - private application and library code (handler, service, etc.)

### pkg - library code that's ok to use by external applications (client to db etc, utils, loggers)

### REST - (cache, stateless, http)

##### 200 ok, 204 - no content, 400 - bad request, 404 - not found, 500 - internal server error
GET /api/users - list of users -- 200, 404, 500 
GET /api/users/1 - user with id 1 -- 200, 404, 500
POST /api/users/:id - create user with id 1 -- 204, 4xx, Header Local: url
PUT /api/users/:id - update user with id 1 (fully) -- 204, 4xx, Header Local: url
PATCH /api/users/:id - update user with id 1 (partially) -- 204, 4xx, Header Local: url
DELETE /api/users/:id - delete user with id 1 -- 204, 4xx, Header Local: url


### some think about dynamic changing logging level
if on prod stage we see some error, we can change logging level to debug and see what's going on
if we see that everything is ok, we can change logging level to info and don't see debug messages

### running app on socket
it might be useful for some cases, when we need to run app on socket, for example, when we don't have tcp port's available
and just run on socket. But if we need provide some kind of load balancing or proxy (nginx), we can't do it with socket, because we can't
use nginx as proxy for socket. So, we need to use tcp port's. But if we need to run app on socket, we can do it with some kind of wrapper
like systemd or supervisord. And we can run app on socket and on tcp port's. And we can use nginx as proxy for tcp port's and use socket
for internal communication between services.