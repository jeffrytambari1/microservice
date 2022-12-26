package main

import (
	"strconv"
    "fmt"
    "log"
    "strings"
	"context"
    "github.com/gofiber/fiber/v2"
)


func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", hello)

	// CRUD API
	app.Get("/product", handleGetProduct)
	app.Post("/product", handlePostProduct)
	app.Put("/product", handlePutProduct)
	app.Delete("/product", handleDeleteProduct)
	
	app.Get("/user", handleGetUser)
	app.Post("/user", handlePostUser)
	app.Put("/user", handlePutUser)
	app.Delete("/user", handleDeleteUser)

	app.Get("/order", handleGetOrder)
	app.Post("/order", handlePostOrder)
	app.Put("/order", handlePutOrder)
	app.Delete("/order", handleDeleteOrder)
	

	app.Get("/order_detail", handleGetOrderDetail)
	app.Post("/order_detail", handlePostOrderDetail)
	app.Put("/order_detail", handlePutOrderDetail)
	app.Delete("/order_detail", handleDeleteOrderDetail)
	

	app.Get("/order_template", handleGetOrderTemplate)
	app.Post("/order_template", handlePostOrderTemplate)
	app.Put("/order_template", handlePutOrderTemplate)
	app.Delete("/order_template", handleDeleteOrderTemplate)

	app.Get("/stock", handleGetStock)
	app.Post("/stock", handlePostStock)
	app.Put("/stock", handlePutStock)
	app.Delete("/stock", handleDeleteStock)

	// LIST
	app.Get("/products", handleSearchProducts)
	app.Get("/users", handleSearchUsers)
	app.Get("/order_details", handleSearchOrderDetails)
	app.Get("/order_templates", handleSearchOrderTemplates)


	// BUSINESS RULES API
	app.Post("/create_cart_item", handleCreateCartItem) // add to chart // create_order_draft
	app.Get("/get_cart_items", handleGetCartItems)
	app.Post("/update_cart_item", handleUpdateCartItem)
	app.Post("/update_cart_item_quantity", handleUpdateCartItemQuantity)
	app.Delete("/delete_cart_item", handleDeleteCartItem)
	app.Post("/checkout", handleCheckout) // create_order_final
	app.Post("/cancel_cart_order", handleCancelCartOrder) // cancel all cart item in order_template which belong to the user_id
	app.Get("/get_order_items", handleGetOrderItems)


	// AUTH
	app.Post("/login", handleLogin)
	app.Post("/logout", handleLogout)


	log.Fatal(app.Listen(":3000"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func initResponse(cc *fiber.Ctx) (isSessionExpired bool) {
	loggingRequest(Logging{Url: cc.OriginalURL(), RemoteAddress: cc.IP()})

    isSessionExpired, _ = checkSession(cc)
    return 
}

func handleLogin(cc *fiber.Ctx) error {
	initResponse(cc)
	login_data := new(LoginData)
	if err := cc.BodyParser(login_data); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Login Failed"}))
	}

	isUserValid, err := login_data.login()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Login Failed"}))
	}

	if !isUserValid {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Username or Password Not Found"}))
	}

	// // return session_code that will be used entire app
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Login Success", SessionCode: login_data.SessionCode, SessionCreatedDatetime: login_data.SessionCreatedDatetime}))
}


func handleLogout(cc *fiber.Ctx) error {
	initResponse(cc)
	login_data := new(LoginData)
	if err := cc.BodyParser(login_data); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Logout Failed"}))
	}
	isUserValidLogout, err := login_data.logout()

	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Logout Failed"}))
	}

	if !isUserValidLogout {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Logout Failed"}))
	}

	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Logout Success"}))
}


// implementation of Session only one function, handleGetProduct, 
// because according to the requirement, Auth Service must use OAuth2.0 of google
// due to asking for OAuth 2.0 Client IDs still under Google Verification process (OAuth consent screen)
func handleGetProduct(cc *fiber.Ctx) error {
	// initResponse(cc)
	if isSessionExpired := initResponse(cc); isSessionExpired {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Message: "Session Expired"}))
	}

    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	product, err := retrieveProduct(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Product Not Found"}))
	}
	return cc.JSON(product)
}
func handlePostProduct(cc *fiber.Ctx) error {
	initResponse(cc)
	product := new(Product)
	if err := cc.BodyParser(product); err != nil {
		// return cc.Status(400).SendString(err.Error())
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Product Failed"}))
	}
	err := product.createProduct()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Product Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Create Product Success"}))
}
func handlePutProduct(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	product, err := retrieveProduct(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Product Not Found"}))
	}
	product2 := new(Product)
	if err := cc.BodyParser(product2); err != nil {
		// return cc.Status(400).SendString(err.Error())
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Product Failed"}))
	}
	product2.ProductId = product.ProductId
	err = product2.updateProduct()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Product Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Update Product Success"}))
}
func handleDeleteProduct(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	product, err := retrieveProduct(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Product Not Found"}))
	}
	err = product.deleteProduct()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Delete Product Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Delete Product Success"}))
}
func handleSearchProducts(cc *fiber.Ctx) error {
	initResponse(cc)
	search_product := new(SearchProduct)
	if err := cc.BodyParser(search_product); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Search Product Failed"}))
	}
	search_products, err := searchProducts(search_product)
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Search Product Failed"}))
	}
	return cc.JSON(search_products)
}





func handleGetUser(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	user, err := retrieveUser(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "User Not Found"}))
	}
	return cc.JSON(user)
}
func handlePostUser(cc *fiber.Ctx) error {
	initResponse(cc)
	user := new(User)
	if err := cc.BodyParser(user); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create User Failed"}))
	}
	search_user := User{
		Username: user.Username,
		Email: user.Email,
	}
	isUsernameOrEmailExist := checkExistingUsernameOrEmail(&search_user)
	if isUsernameOrEmailExist {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Message: "Username or Email Already Exists."}))
	}
	err := user.createUser()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create User Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Create User Success"}))
}
func handlePutUser(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	user, err := retrieveUser(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "User Not Found"}))
	}
	user2 := new(User)
	if err := cc.BodyParser(user2); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update User Failed"}))
	}
	user2.UserId = user.UserId
	err = user2.updateUser()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update User Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Update User Success"}))
}
func handleDeleteUser(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	user, err := retrieveUser(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "User Not Found"}))
	}
	err = user.deleteUser()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Delete User Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Delete User Success"}))
}
func handleSearchUsers(cc *fiber.Ctx) error {
	initResponse(cc)
	search_user := new(SearchUser)
	if err := cc.BodyParser(search_user); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Search User Failed"}))
	}
	search_users, err := searchUsers(search_user)
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Search User Failed"}))
	}
	return cc.JSON(search_users)
}




func handleGetOrder(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	order, err := retrieveOrder(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Order Not Found"}))
	}
	return cc.JSON(order)
}
func handlePostOrder(cc *fiber.Ctx) error {
	initResponse(cc)
	order := new(Order)
	if err := cc.BodyParser(order); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Order Failed"}))
	}
	err := order.createOrder()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Order Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Create Order Success"}))
}
func handlePutOrder(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	order, err := retrieveOrder(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Order Not Found"}))
	}
	order2 := new(Order)
	if err := cc.BodyParser(order2); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Order Failed"}))
	}
	order2.OrderId = order.OrderId
	err = order2.updateOrder()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Order Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Update Order Success"}))
}
func handleDeleteOrder(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	order, err := retrieveOrder(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Order Not Found"}))
	}
	err = order.deleteOrder()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Delete Order Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Delete Order Success"}))
}
func handleSearchOrderDetails(cc *fiber.Ctx) error {
	initResponse(cc)
	search_order_detail := new(SearchOrderDetail)
	if err := cc.BodyParser(search_order_detail); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Search Order Detail Failed"}))
	}
	search_order_details, err := searchOrderDetails(search_order_detail)
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Search Order Detail Failed"}))
	}
	return cc.JSON(search_order_details)
}




func handleGetOrderDetail(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	order_detail, err := retrieveOrderDetail(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Order Detail Not Found"}))
	}
	return cc.JSON(order_detail)
}
func handlePostOrderDetail(cc *fiber.Ctx) error {
	initResponse(cc)
	order_detail := new(OrderDetail)
	if err := cc.BodyParser(order_detail); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Order Detail Failed"}))
	}
	err := order_detail.createOrderDetail()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Order Detail Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Create Order Detail Success"}))
}
func handlePutOrderDetail(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	order_detail, err := retrieveOrderDetail(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Order Detail Not Found"}))
	}
	order_detail2 := new(OrderDetail)
	if err := cc.BodyParser(order_detail2); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Order Detail Failed"}))
	}
	order_detail2.OrderDetailId = order_detail.OrderDetailId
	err = order_detail2.updateOrderDetail()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Order Detail Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Update Order Detail Success"}))
}
func handleDeleteOrderDetail(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	order_detail, err := retrieveOrderDetail(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Order Detail Not Found"}))
	}
	err = order_detail.deleteOrderDetail()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Delete Order Detail Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Delete Order Detail Success"}))
}




func handleGetStock(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	stock, err := retrieveStock(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Stock Not Found"}))
	}
	return cc.JSON(stock)
}
func handlePostStock(cc *fiber.Ctx) error {
	initResponse(cc)
	stock := new(Stock)
	if err := cc.BodyParser(stock); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Stock Failed"}))
	}
	err := stock.createStock()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Stock Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Create Stock Success"}))
}
func handlePutStock(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	stock, err := retrieveStock(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Stock Not Found"}))
	}
	stock2 := new(Stock)
	if err := cc.BodyParser(stock2); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Stock Failed"}))
	}
	stock2.StockId = stock.StockId
	err = stock2.updateStock()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Stock Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Update Stock Success"}))
}
func handleDeleteStock(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	stock, err := retrieveStock(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Stock Not Found"}))
	}
	err = stock.deleteStock()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Delete Stock Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Delete Stock Success"}))
}




func handleGetOrderTemplate(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	order_template, err := retrieveOrderTemplate(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Order Template Not Found"}))
	}
	return cc.JSON(order_template)
}
func handlePostOrderTemplate(cc *fiber.Ctx) error {
	initResponse(cc)
	order_template := new(OrderTemplate)
	if err := cc.BodyParser(order_template); err != nil {
		// return cc.Status(400).SendString(err.Error())
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Order Template Failed"}))
	}
	err := order_template.createOrderTemplate()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Order Template Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Create Order Template Success"}))
}
func handlePutOrderTemplate(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	order_template, err := retrieveOrderTemplate(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Order Template Not Found"}))
	}
	order_template2 := new(OrderTemplate)
	if err := cc.BodyParser(order_template2); err != nil {
		// return cc.Status(400).SendString(err.Error())
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Order Template Failed"}))
	}
	order_template2.OrderTemplateId = order_template.OrderTemplateId
	err = order_template2.updateOrderTemplate()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Order Template Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Update Order Template Success"}))
}
func handleDeleteOrderTemplate(cc *fiber.Ctx) error {
	initResponse(cc)
    id, err := strconv.Atoi(cc.Query("id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter id"}))
	}
	order_template, err := retrieveOrderTemplate(int64(id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Order Template Not Found"}))
	}
	err = order_template.deleteOrderTemplate()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Delete Order Template Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Delete Order Template Success"}))
}
func handleSearchOrderTemplates(cc *fiber.Ctx) error {
	initResponse(cc)
	search_order_template := new(SearchOrderTemplate)
	if err := cc.BodyParser(search_order_template); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Search Order Template Failed"}))
	}
	search_order_templates, err := searchOrderTemplates(search_order_template)
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Search Order Template Failed"}))
	}
	return cc.JSON(search_order_templates)
}



// BUSINESS RULES API
func handleCreateCartItem(cc *fiber.Ctx) error {
	initResponse(cc) // handleCreateCart / handleCreateOrderDraft
	cart_item := new(CartItem)
	if err := cc.BodyParser(cart_item); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Cart Item Failed"}))
	}

	msg, err := cart_item.createCartItem()
	if err != nil {
		// return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Create Cart Item Failed"}))
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: msg }))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Create Cart Item Success"}))
}
func handleGetCartItems(cc *fiber.Ctx) error {
	initResponse(cc)
    user_id, err := strconv.Atoi(cc.Query("user_id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Need Parameter user_id"}))
	}
	cart_item := new(CartItem)
	cart_item.UserId = int64(user_id)

	// cart_items, err := getCartList(cart_item)
	cart_items, err := getCartItems(cart_item)
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Get My Cart Items Failed"}))
	}
	return cc.JSON(cart_items)
}
func handleUpdateCartItem(cc *fiber.Ctx) error {
	initResponse(cc)
	cart_item := new(CartItem)
	if err := cc.BodyParser(cart_item); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Cart Item Failed"}))
	}

	msg, err := cart_item.updateCartItem()
	if err != nil {
		fmt.Println("error after updateCartItem, inside err")
		fmt.Println(msg)
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: msg }))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Update Cart Item Success"}))
}
func handleUpdateCartItemQuantity(cc *fiber.Ctx) error {
	initResponse(cc)
	cart_item := new(CartItem)
	if err := cc.BodyParser(cart_item); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Update Cart Item Quantity Failed"}))
	}

	msg, err := cart_item.updateCartItemQuantity()
	if err != nil {
		// fmt.Println("error after updateCartItem, inside err")
		// fmt.Println(msg)
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: msg }))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Update Cart Item Quantity Success"}))
}
func handleDeleteCartItem(cc *fiber.Ctx) error {
	initResponse(cc)
    order_template_id, err := strconv.Atoi(cc.Query("order_template_id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn( ErrorStruct{ Error: err, Message: "Need Parameter order_template_id"}))
	}
	order_template, err := retrieveOrderTemplate(int64(order_template_id))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Cart Item Not Found"}))
	}
	err = order_template.deleteOrderTemplate()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Delete Cart Item Failed"}))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Delete Cart Item Success"}))
}
func handleCheckout(cc *fiber.Ctx) error {
	initResponse(cc) // handleCreateOrderFinal
	cart_item := new(CartItem)
	if err := cc.BodyParser(cart_item); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Checkout Failed"}))
	}

	ctx := context.Background()
	msg, err := cart_item.checkout(ctx)
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: msg }))
	}
	return cc.JSON(JSONSuccessReturn( SuccessStruct{Message: "Checkout Success"}))
}
func handleCancelCartOrder(cc *fiber.Ctx) error {
	initResponse(cc)
	cart_item := new(CartItem)
	if err := cc.BodyParser(cart_item); err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Cancel Cart Order Failed"}))
	}
	msg, err := cart_item.cancelCartOrder()
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: msg }))
	}
	return cc.JSON(JSONSuccessReturn(SuccessStruct{Message: "Update Cart Item Quantity Success"}))
}
func handleGetOrderItems(cc *fiber.Ctx) error {
	initResponse(cc)
	user_id, err := strconv.Atoi(cc.Query("user_id"))
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Need Parameter user_id"}))
	}
	order_item := new(OrderItem)
	order_item.UserId = int64(user_id)

	order_items, err := getOrderItems(order_item)
	if err != nil {
		return cc.JSON(JSONErrorReturn(ErrorStruct{ Error: err, Message: "Get My Order Items Failed"}))
	}
	return cc.JSON(order_items)	
}

func checkSession(cc *fiber.Ctx) (isSessionExpired bool, err error) {
	session_data := new(SessionData)
	// session_data.SessionCode = cc.Query("session")

	headers := cc.GetReqHeaders()
	auth1 := headers["Authorization"]
	// fmt.Println(headers["Bearer"])
	// fmt.Println(headers["Token"])
	bearerToken := strings.Fields(auth1)
	session_data.SessionCode = ""

	if (isset(bearerToken, 1) && bearerToken[1] != "") {
		session_data.SessionCode = bearerToken[1]
	}

	isSessionExpired, _ = session_data.checkSessionDb()
	return
}

