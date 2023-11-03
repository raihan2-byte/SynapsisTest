package cart

type CartsResponse struct {
	ID         int
	Quantity   int
	TotalPrice int
	ProductID  int
	CategoryID int
	UserID     int
	Product    ProductCartResponse
	User       UserCartResponse
}

type ProductCartResponse struct {
	ID    int
	Title string
	Price int
}

type UserCartResponse struct {
	ID       int
	FullName string
}

type CartsResponsePost struct {
	ID         int
	Quantity   int
	TotalPrice int
	ProductID  int
	CategoryID int
	UserID     int
}

func FormatterCartPost(cart Carts) CartsResponsePost {
	formatterCart := CartsResponsePost{
		ID:         cart.CategoryID,
		Quantity:   cart.Quantity,
		TotalPrice: cart.TotalPrice,
		ProductID:  cart.ProductID,
		CategoryID: cart.CategoryID,
		UserID:     cart.UserID,
	}
	return formatterCart
}

func FormatterCart(cart Carts) CartsResponse {
	formatterCart := CartsResponse{
		ID:         cart.CategoryID,
		Quantity:   cart.Quantity,
		TotalPrice: cart.TotalPrice,
		ProductID:  cart.ProductID,
		CategoryID: cart.CategoryID,
		UserID:     cart.UserID,
	}

	formatterProduct := cart.Product

	productFormatter := ProductCartResponse{}
	productFormatter.ID = formatterProduct.ID
	productFormatter.Price = formatterProduct.Price

	formatterCart.Product = productFormatter

	formatterUser := cart.User
	userFormatter := UserCartResponse{}
	userFormatter.ID = formatterUser.ID
	userFormatter.FullName = formatterUser.FullName

	formatterCart.User = userFormatter

	return formatterCart

}

func FormatterGetCart(cart []Carts) []CartsResponse {
	cartGetFormatter := []CartsResponse{}

	for _, carts := range cart {
		beritaFormatter := FormatterCart(carts)
		cartGetFormatter = append(cartGetFormatter, beritaFormatter)
	}

	return cartGetFormatter
}
