package dylan

type Account struct {
	Name string `json:"name,omitempty" form:"name" binding:"required"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty" form:"age" binding:"required,gte=6,lte=100"`
}
