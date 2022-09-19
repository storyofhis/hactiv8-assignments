package params

type UserInput struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"reqired"`
}
