## Exercise 3: Event-Driven Order Processing System

Goal: Build a simple event-driven microservice using message queues.
Requirements:
* Create a pub/sub system using Redis Streams or NATS (or even channels for simplicity)
* Implement an "Order Service" that publishes order events
* Build an "Inventory Service" that consumes events and updates stock
* Add idempotency handling (process each event exactly once)
* Implement circuit breaker pattern for external calls
* Include retry logic with exponential backoff
Key Concepts: Event-driven architecture, message queues, concurrency patterns, error handling, at-least-once delivery


## TODO
[ ] handle case where amount ordered exceeds inventory stock.
[ ] list inventory is nondeterministic