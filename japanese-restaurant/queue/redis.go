package queue

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const stream = "orders"

type RedisQueue struct {
	rdb *redis.Client
	// XRead can be read using $, but this will only return messages that arrive after the XRead call starts.
	// Messages may be missed when we stop listening to process the order. So intead, we track the last ID seen.
	lastID string
}

func NewRedisStreamsQueue() Queue {
	return &RedisQueue{
		rdb:    redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
		lastID: "0",
	}
}

func (r RedisQueue) Close() {
	r.rdb.Close()
}

// Blocking
func (r *RedisQueue) ReceiveOrder(ctx context.Context) (Order, error) {
	results, err := r.rdb.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, r.lastID},
		Block:   1000 * time.Millisecond, // timeout after 1 sec
		Count:   1,
	}).Result()

	if err != nil {
		return Order{}, err
	}

	msg := results[0].Messages[0]

	r.lastID = msg.ID

	var order Order
	err = order.UnmarshalBinary([]byte(msg.Values["order"].(string)))
	return order, err
}

func (r *RedisQueue) SendOrder(ctx context.Context, order Order) error {
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
