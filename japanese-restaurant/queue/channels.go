package queue

import (
	"context"
	"errors"
)

// https://gobyexample.com/select

type channelQueue struct {
	orderChan chan Order
}

func NewChannelQueue() Queue {
	q := make(chan Order, 2)
	return &channelQueue{
		orderChan: q,
	}
}

func (q *channelQueue) SendOrder(ctx context.Context, order Order) error {
	select {
	case q.orderChan <- order:
		return nil
	case <-ctx.Done():
		return nil
	}
}

func (q *channelQueue) ReceiveOrder(ctx context.Context) (Order, error) {
	select {
	case order, ok := <-q.orderChan:
		if !ok {
			return Order{}, errors.New("queue closed")
		}
		return order, nil
	case <-ctx.Done():
		return Order{}, ctx.Err()
	}
}

func (q *channelQueue) Close() {
	close(q.orderChan)
}

func (q *channelQueue) Drain(ctx context.Context) ([]Order, error) {
	panic("unimplemented")
}
