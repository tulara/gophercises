## Exercise
Build a rate-limited REST API for tracking and discovering coffee shops with JWT authentication, middleware, and graceful shutdown.

### Authentication
- `POST /register` - Register new user (username, password)
- `POST /login` - Login and receive JWT token

### Cafes (All require authentication)
- `GET /cafes` - List cafes with pagination
- `GET /cafes/{id}` - Get single cafe
- `PUT /cafes/{id}` - Create or update cafe (only owner or admin)

ids are treated as idemponent; PUT updates with new details to the same id will overwrite pre-existing records. 

### How to Run
```
cd coffee-shop 
go run .
```

```
go clean -testcache
```

### To do

[ ] implement graceful shutdown e.g. https://dev.to/mokiat/proper-http-shutdown-in-go-3fji
-> remind myself how channels work.

[ ] logger

[ ] stub `time` in jwt so that expiry on tokens can be tested properly (via unit tests or config so that don't need to wait 24 hours in a test environment.)

[] Tests are failing because you haven't mocked the auth. Because this is a learning project to get back up to speed I have deprioritised tests but there is an exercise in going back through and properly testing. And particularly reminding self of old ways.

[ ] rate limiting

### Bugs and tidys
- Sort out ID generation and assumptions made in memory store. 

- If the server is restarted and a previous unexpired JWT token is used, the requests is authorised despite the user no longer existing. Probably would need to do something like purge/expire all tokens on restart.

- inconsistent json encoding (NewEncoder vs Marshal) and inconsistent error responses.
- input validation (content-type header validation, request body size limits, pagination parameter validation)