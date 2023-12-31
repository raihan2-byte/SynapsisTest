package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"synapsisid/auth"
	"synapsisid/cart"
	"synapsisid/category"
	"synapsisid/handler"
	"synapsisid/helper"
	"synapsisid/product"
	"synapsisid/transaction"
	transactiondetail "synapsisid/transactionDetail"
	"synapsisid/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	dbUsername := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbName := os.Getenv("MYSQLDATABASE")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(
		&user.User{},
		&category.Categorys{},
		&product.Products{},
		&transaction.Transaction{},
		&cart.Carts{},
		&transactiondetail.TransactionDetails{},
	)

	//user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	//category
	categoryRepository := category.NewRepositoryCategory(db)
	categoryService := category.NewServiceCategory(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	//product
	productRepository := product.NewRepositoryProduct(db)
	productService := product.NewServiceProduct(productRepository, categoryRepository)
	productHandler := handler.NewProductHandler(productService)

	//transactionDetails
	transactionDetailsRepository := transactiondetail.NewRepositoryTransactionDetails(db)
	transactionDetailsService := transactiondetail.NewService(transactionDetailsRepository)
	transactionDetailsHandler := handler.NewtransactionDetailsHandler(transactionDetailsService)


	//transaction
	transactionRepository := transaction.NewRepositoryTransaction(db)
	transactionService := transaction.NewService(transactionRepository, productRepository, userRepository, transactionDetailsRepository)
	transactionHandler := handler.NewtransactionHandler(transactionService)

	//addtocart
	addToCartRepository := cart.NewRepositoryCart(db)
	addToCartService := cart.NewServiceCart(addToCartRepository, categoryRepository, productRepository, userRepository, transactionDetailsRepository, transactionRepository)
	cartHandler := handler.NewAddToCartHandler(addToCartService)



	router := gin.Default()
	api := router.Group("/users")
	api2 := router.Group("/categories")
	api3 := router.Group("/products")
	api4 := router.Group("/transactions")
	api5 := router.Group("/cart")
	

	//user
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.PATCH("/topup", authMiddleware(authService, userService), userHandler.UpdatedUser)
	
	//category
	api2.POST("/", authMiddleware(authService, userService), authRole(authService, userService), categoryHandler.CreateCategory)
	api2.GET("/", categoryHandler.GetAllCategory)
	api2.GET("/:id", categoryHandler.GetCategory)
	api2.PATCH("/:id", authMiddleware(authService, userService), authRole(authService, userService), categoryHandler.UpdatedCategory)
	api2.DELETE("/:id", authMiddleware(authService, userService), authRole(authService, userService), categoryHandler.DeletedCategory)

	//product
	api3.POST("/", authMiddleware(authService, userService), authRole(authService, userService), productHandler.CreateProduct)
	api3.GET("/:id", authMiddleware(authService, userService), productHandler.GetProduct)
	api3.GET("/", authMiddleware(authService, userService), productHandler.GetAllProduct)
	api3.GET("/category/:categoryID", authMiddleware(authService, userService), productHandler.GetAllProductByCategory)
	api3.PUT("/:id", authMiddleware(authService, userService), authRole(authService, userService), productHandler.UpdateProduct)
	api3.DELETE("/:id", authMiddleware(authService, userService), authRole(authService, userService), productHandler.DeleteProduct)


	//transaction
	api4.POST("/:id", authMiddleware(authService, userService),  transactionHandler.CreateTransaction)
	api4.GET("/", authMiddleware(authService, userService), transactionHandler.GetTransaction)

	//cart
	api5.POST("/:id", authMiddleware(authService, userService), cartHandler.AddToCart)
	api5.PUT("/:cartID", authMiddleware(authService, userService), cartHandler.UpdateCart)
	api5.GET("/", authMiddleware(authService, userService), cartHandler.GetAllChart)
	api5.POST("/PayFromChart", authMiddleware(authService, userService), cartHandler.PayByUserID)
	api5.POST("/PayByCartID/:id", authMiddleware(authService, userService), cartHandler.PayByCartID)
	api5.DELETE("/DeleteCart/:id", authMiddleware(authService, userService), cartHandler.DeleteCart)

	router.GET("/transaction-details", authMiddleware(authService, userService), authRole(authService, userService), transactionDetailsHandler.GetAllTransactionDetails )

	

	//tambah request body aja di postman gitu
	router.Run(":8000")

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// fmt.Println(authHeader)
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			//nah ini kalau emang ada dua key nya dan sesuai, maka tokenString tadi masuk ke arrtoken index ke1
			tokenString = arrToken[1]
		}
		token, err := authService.ValidasiToken(tokenString)
		// fmt.Println(token, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		// fmt.Println(claim, ok)
		if !ok || !token.Valid {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByid(userID)
		// fmt.Println(user, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}


func authRole(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// fmt.Println(authHeader)
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			//nah ini kalau emang ada dua key nya dan sesuai, maka tokenString tadi masuk ke arrtoken index ke1
			tokenString = arrToken[1]
		}
		token, err := authService.ValidasiToken(tokenString)
		// fmt.Println(token, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		// fmt.Println(claim, ok)
		if !ok || !token.Valid {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		if int(claim["role"].(float64)) != 1 {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user, err := userService.GetUserByid(userID)
		// fmt.Println(user, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
