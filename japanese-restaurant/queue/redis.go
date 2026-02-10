package queue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const stream = "orders"

type RedisQueue struct {
	rdb *redis.Client
}

func NewRedisStreamsQueue() Queue {
	return RedisQueue{
		rdb: redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
	}
}

// Close implements [Queue].
func (r RedisQueue) Close() {
	r.rdb.Close()
}

// Blocking
func (r RedisQueue) ReceiveOrder(ctx context.Context) (Order, error) {
	panic("unimplemented")
}

func (r RedisQueue) SendOrder(ctx context.Context, order Order) error {
	// what happens on retry?
	_, err := r.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "orders",
		Values: map[string]interface{}{"order": &order},
	}).Result()

	if err != nil {
		return err
	}
	return nil
}
