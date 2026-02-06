package queue

import (
	"context"
	"testing"
)

func TestSendAndRecieve(t *testing.T) {
	q := NewChannelQueue()
	order := Order{
		Meals: map[int]int{
			1: 1,
		},
	}

	q.SendOrder(context.Background(), order)
	recievedOrder, err := q.RecieveOrder(context.Background())

	if err != nil {
		t.Fatalf("Error where none expected: %v", err)
	}

	if recievedOrder.ID() != order.ID() {
		t.Fail()
	}
}
