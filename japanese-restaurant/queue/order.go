package queue

import (
	"encoding/json"
	"sync/atomic"
)

var nextOrderID atomic.Int64

type Order struct {
	ID int `json:"id"`
	//meal id and number ordered
	Meals map[int]int `json:"meals"`
}

func NewOrder(meals map[int]int) Order {
	id := nextOrderID.Add(1)
	return Order{
		ID:    int(id),
		Meals: meals,
	}
}

func (o *Order) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}

func (o *Order) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, o)
}
