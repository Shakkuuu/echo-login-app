package entity

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
