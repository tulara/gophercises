package queue

type Order struct {
	ID    string
	Meals map[int]int //meal id and number ordered
}
