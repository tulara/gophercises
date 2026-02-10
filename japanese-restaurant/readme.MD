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
[ ] refresh queue/messaging concepts
 - consumer groups for multiple "kitchen workers"
[ ] customers place orders in English
[ ] list inventory is nondeterministic


## To Run

Start redis container
```
docker run -d --name redis -p 6379:6379 redis:latest
```