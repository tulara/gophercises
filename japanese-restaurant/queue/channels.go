package queue

import "context"

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

func (q *channelQueue) SendOrder(order Order) bool {
	select {
	case q.orderChan <- order:
		return true
	default:
		// log buffer overflow?
		return false
	}
}

func (q *channelQueue) RecieveOrder(ctx context.Context) (Order, error) {
	select {
	case order := <-q.orderChan:
		return order, nil
	case <-ctx.Done():
		return Order{}, ctx.Err()
	}
}
