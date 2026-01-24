package domain

type Cafe struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
