package cart

type InputCart struct {
	Quantity int `json:"quantiy" binding:"required"`
}

type PayProductFromCart struct {
	CartsID []int `json:"product_id"`
}

type GetIdProduct struct {
	ID int `uri:"id" binding:"required"`
}

type GetIDCart struct {
	ID int `uri:"id" binding:"required"`
}
