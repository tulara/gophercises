package queue

import (
	"context"
	"testing"
)

func TestSendAndRecieve(t *testing.T) {
	q := Init()
	order := Order{
		ID: "1",
		Meals: map[int]int{
			1: 1,
		},
	}

	q.SendOrder(order)
	recievedOrder, err := q.RecieveOrder(context.Background())

	if err != nil {
		t.Fatalf("Error where none expected: %v", err)
	}

	if recievedOrder.ID != order.ID {
		t.Fail()
	}
}
