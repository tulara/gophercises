package queue

var orderChan chan Order

func Init() {
	orderChan = make(chan Order)
}

func SendOrder(order Order) {

}

func RecieveOrder() *Order {
	return nil
}
