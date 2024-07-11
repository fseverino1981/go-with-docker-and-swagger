package response

type UserRequest struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
