## Exercise
Build a rate-limited REST API for tracking and discovering coffee shops with JWT authentication, middleware, and graceful shutdown.

### Authentication
- `POST /auth/register` - Register new user (username, password)
- `POST /auth/login` - Login and receive JWT token

### Cafes (All require authentication)
- `GET /cafes` - List cafes with filtering and pagination
- `GET /cafes/{id}` - Get single cafe
- `PUT /cafes/{id}` - Create or update cafe (only owner or admin)
- `DELETE /cafes/{id}` - Delete cafe (only owner or admin)
- `POST /cafes/{id}/visit` - Record a visit (increments visit_count, updates last_visited)

## To Do

[ ] Make sure Memory store has cafes sorted by ID.