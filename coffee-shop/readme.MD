## Exercise
Build a rate-limited REST API for tracking and discovering coffee shops with JWT authentication, middleware, and graceful shutdown.

### Authentication
- `POST /auth/register` - Register new user (username, password)
- `POST /auth/login` - Login and receive JWT token

### Cafes (All require authentication)
- `GET /cafes` - List cafes with pagination
- `GET /cafes/{id}` - Get single cafe
- `PUT /cafes/{id}` - Create or update cafe (only owner or admin)
- `DELETE /cafes/{id}` - Delete cafe (only owner or admin)
- `POST /cafes/{id}/visit` - Record a visit (increments visit_count, updates last_visited)

ids are treated as idemponent; PUT updates with new details to the same id will overwrite pre-existing records. 

### How to Run
```
cd coffee-shop 
go run .
```

```
go clean -testcache
```

### BUGS

[ ] not returning full set of cafes or default page size