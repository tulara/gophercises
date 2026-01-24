## Exercise
Build a rate-limited REST API for tracking and discovering coffee shops with JWT authentication, middleware, and graceful shutdown.

### Authentication
- `POST /auth/register` - Register new user (username, password)
- `POST /auth/login` - Login and receive JWT token

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