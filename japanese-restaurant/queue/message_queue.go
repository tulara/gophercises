package queue

type Queue interface {
	SendOrder(order Order)
	RecieveOrder() Order
}
