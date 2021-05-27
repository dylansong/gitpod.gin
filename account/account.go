package account

type Account struct {
	Name string `json:"name" form:"name" bind:"required"`
	Age  int    `json:"age" form:"age" bind:"required,min=6,max=100"`
}
