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

func (r RedisQueue) Close() {
	r.rdb.Close()
}

// Blocking
func (r RedisQueue) ReceiveOrder(ctx context.Context) (Order, error) {
	results, err := r.rdb.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, "$"}, // start from the last id we saw (new messages only)
		Block:   2000,                  // timeout after 2 seconds
		Count:   1,
	}).Result()

	if err != nil {
		return Order{}, err
	}

	msg := results[0].Messages[0]
	var order Order
	err = order.UnmarshalBinary([]byte(msg.Values["order"].(string)))
	return order, err
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
