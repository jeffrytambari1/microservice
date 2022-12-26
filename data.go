package main

import (
	"fmt"
	"database/sql"
	"time"
	"strings"
	"context"
	_ "github.com/go-sql-driver/mysql"
    "github.com/google/uuid"
)

var Db *sql.DB

const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "your_user"
	password = "your_password"
	dbname   = "go_microservice"
)

type Product struct {
	ProductId 			int64 		`json:"product_id"`
	ProductName 		string 		`json:"product_name"`
	Price 				float64 	`json:"price"`
	CreatedDatetime 	string 		`json:"created_datetime"`
	UpdatedDatetime 	string 		`json:"updated_datetime"`
}

type User struct {
	UserId 				int64 		`json:"user_id"`
	Username 			string 		`json:"username"`
	Email 				string	 	`json:"email"`
	CreatedDatetime 	string 		`json:"created_datetime"`
	UpdatedDatetime 	string 		`json:"updated_datetime"`
	Password 			string	 	`json:"password"`
}

type Order struct {
	OrderId 			int64 		`json:"order_id"`
	OrderDatetime 		string 		`json:"order_datetime"`
	CheckoutDatetime 	string 		`json:"checkout_datetime"`
	UserId 				int64	 	`json:"user_id"`
}

type OrderDetail struct {
	OrderDetailId 		int64 		`json:"order_detail_id"`
	OrderId 			int64 		`json:"order_id"`
	ProductId 			int64 		`json:"product_id"`
	Quantity 			int64 		`json:"quantity"`
	Price 				float64		`json:"price"`
	CreatedDatetime 	string 		`json:"created_datetime"`
	UpdatedDatetime 	string 		`json:"updated_datetime"`
}

type Stock struct {
	StockId 			int64 		`json:"stock_id"`
	ProductId 			int64 		`json:"product_id"`
	Quantity 			int64 		`json:"quantity"`
	CreatedDatetime 	string 		`json:"created_datetime"`
	UpdatedDatetime 	string 		`json:"updated_datetime"`
}

type OrderTemplate struct {
	OrderTemplateId 	int64 		`json:"order_template_id"`
	ProductId 			int64 		`json:"product_id"`
	OrderId 			int64 		`json:"order_id"`
	Quantity 			int64 		`json:"quantity"`
	Price 				float64		`json:"price"`
	UserId 				int64 		`json:"user_id"`
	CartPhaseType		int8 		`json:"cart_phase_type"` // 0 = cart, 1 = checkout, 2 = canceled
	CreatedDatetime 	string 		`json:"created_datetime"`
	UpdatedDatetime 	string 		`json:"updated_datetime"`
}

type CartItem struct {
	OrderTemplateId 	int64 		`json:"order_template_id"`
	ProductId 			int64 		`json:"product_id"`
	OrderId 			int64 		`json:"order_id"`
	Quantity 			int64 		`json:"quantity"`
	Price 				float64		`json:"price"`
	UserId 				int64 		`json:"user_id"`
	CartPhaseType		int8 		`json:"cart_phase_type"` // cart item must 0
	CreatedDatetime 	string 		`json:"created_datetime"`
	UpdatedDatetime 	string 		`json:"updated_datetime"`
	ProductName 		string 		`json:"product_name"`
	ProductPrice 		float64		`json:"product_price"`
	Username 			string 		`json:"username"`
}

type OrderItem struct {
	OrderId 			int64 		`json:"order_id"`
	// Invoice
	UserId 				int64 		`json:"user_id"`
	OrderDetailId 		int64 		`json:"order_detail_id"`
	ProductId 			int64 		`json:"product_id"`
	Quantity 			int64 		`json:"quantity"`
	Price 				float64		`json:"price"`
	// CreatedDatetime 	string 		`json:"created_datetime"`
	// UpdatedDatetime 	string 		`json:"updated_datetime"`
	ProductName 		string 		`json:"product_name"`
	OrderDatetime 		string 		`json:"order_datetime"`
	// CheckoutDatetime 	string 		`json:"checkout_datetime"` // remove later, OrderDatetime is enough
	Username 			string 		`json:"username"`
}

type SearchProduct struct {
	ProductId 				int64 		`json:"product_id"`
	ProductName 			string 		`json:"product_name"`
	Price 					float64 	`json:"price"`
	CreatedDatetime 		string 		`json:"created_datetime"`
	UpdatedDatetime 		string 		`json:"updated_datetime"`
	StockId 				int64 		`json:"stock_id"`
	Quantity 				int64 		`json:"quantity"`
	StockCreatedDatetime 	string 		`json:"stock_created_datetime"`
	StockUpdatedDatetime 	string 		`json:"stock_updated_datetime"`
}

type SearchUser struct {
	UserId 				int64 		`json:"user_id"`
	Username 			string 		`json:"username"`
	Email 				string	 	`json:"email"`
	CreatedDatetime 	string 		`json:"created_datetime"`
	UpdatedDatetime 	string 		`json:"updated_datetime"`
	Password 			string	 	`json:"password"`
}

type SearchOrderDetail struct {
	OrderDetailId 		int64 		`json:"order_detail_id"`
	OrderId 			int64 		`json:"order_id"`
	ProductId 			int64 		`json:"product_id"`
	Quantity 			int64 		`json:"quantity"`
	Price 				float64		`json:"price"`
	CreatedDatetime 	string 		`json:"created_datetime"`
	UpdatedDatetime 	string 		`json:"updated_datetime"`
	OrderDatetime 		string 		`json:"order_datetime"`
	CheckoutDatetime 	string 		`json:"checkout_datetime"`
	UserId 				int64 		`json:"user_id"`
	Username 			string 		`json:"username"`
	// Email 				string 		`json:"email"`
}

type SearchOrderTemplate struct {
	OrderTemplateId 	int64 		`json:"order_template_id"`
	OrderId 			int64 		`json:"order_id"`
	ProductId 			int64 		`json:"product_id"`
	Quantity 			int64 		`json:"quantity"`
	Price 				float64		`json:"price"`
	CreatedDatetime 	string 		`json:"created_datetime"`
	UpdatedDatetime 	string 		`json:"updated_datetime"`
	UserId 				int64 		`json:"user_id"`
	Username 			string 		`json:"username"`
	// Email 			string 		`json:"email"`
	CartPhaseType		string 		`json:"cart_phase_type"`
}


type Products struct {
	Products []Product `json:"products"`
}

type Users struct {
	Users []User `json:"users"`
}

type SearchProducts struct {
	SearchProducts []SearchProduct `json:"products"`
}

type SearchUsers struct {
	SearchUsers []SearchUser `json:"users"`
}

type SearchOrderDetails struct {
	SearchOrderDetails []SearchOrderDetail `json:"order_details"`
}

type SearchOrderTemplates struct {
	SearchOrderTemplates []SearchOrderTemplate `json:"order_templates"`
}


type CartItems struct {
	CartItems []CartItem `json:"cart_items"`
}

type OrderTemplates struct {
	OrderTemplates []OrderTemplate `json:"order_templates"`
}

type OrderItems struct {
	OrderItems []OrderItem `json:"order_items"`
}


type Logging struct {
	Url 			string
	RemoteAddress 	string
}

type LoginData struct {
	UserId					int64 		`json:"user_id"`
	Username				string 		`json:"username"`
	Email 					string 		`json:"email"`
	Password 				string 		`json:"password"`
	PasswordCrypted			string 		`json:"password_crypted"`
	SessionCode				string 		`json:"session_code"`
	SessionCreatedDatetime	string 		`json:"session_created_datetime"`
}

type SessionData struct {
	UserId					int64 		`json:"user_id"`
	SessionCode				string 		`json:"session_code"`
	SessionCreatedDatetime	string 		`json:"session_created_datetime"`
}

// type CartPhaseType1 int8

const (
	Carted 			int8 	= 0 // CartPhaseType1
	Ordered					= 1
	Canceled 				= 2
)

func Connect() error {
	// fmt.Println("Connect run")
	var err error
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return err
	}
	if err = Db.Ping(); err != nil {
		return err
	}
	return nil
}



// CRUD MODEL FUNCTIONS
func retrieveProduct(id int64) (product Product, err error) {
	product = Product{}
	err = Db.QueryRow("select product_id, product_name, price, created_datetime, updated_datetime from products where product_id = ?", id).Scan(&product.ProductId, &product.ProductName, &product.Price, &product.CreatedDatetime, &product.UpdatedDatetime)
	return

}
func (product *Product) createProduct() (err error) {
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)
 	res, err := Db.Exec("insert into products (product_name, price, created_datetime, updated_datetime) values (?, ?, ?, ?)", product.ProductName, product.Price, dt, dt)
    if err != nil {
    	logger(err)
        return
    }
	product.ProductId, err = res.LastInsertId()
	if err != nil {
		logger(err)
	    return
	}
	return
}
func (product *Product) updateProduct() (err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
	_, err = Db.Exec("update products set product_name = ?, price = ?, updated_datetime = ? where product_id = ?", product.ProductName, product.Price, dt, product.ProductId)
	return
}
func (product *Product) deleteProduct() (err error) {
	_, err = Db.Exec("delete from products where product_id = ?", product.ProductId)
	return
}

func searchProducts(search_product *SearchProduct) (search_products SearchProducts, err error) {
	w := map[string]interface{}{"pr.product_name": search_product.ProductName}
	var whereVal []interface{}
	var whereCol []string
	for k, v := range w {
	    // whereVal = append(whereVal, v)
	    // // whereCol = append(whereCol, fmt.Sprintf("%s = ?", k))
	    // whereCol = append(whereCol, fmt.Sprintf("%s like ?",  "%" + k + "%"))
	    // whereVal = append(whereVal, v)
	    // whereVal = append(whereVal, fmt.Sprintf("%%" + "%s" + "%%", v))
	    whereVal = append(whereVal, fmt.Sprintf("%%%s%%", v))
	    whereCol = append(whereCol, fmt.Sprintf("%s like ?", k))
	}
	rows, err := Db.Query(`
		SELECT pr.product_id, pr.product_name, pr.price, pr.created_datetime, pr.updated_datetime,
			st.stock_id, st.quantity, st.created_datetime stock_created_datetime, st.updated_datetime stock_updated_datetime
		FROM products pr
		LEFT JOIN stocks st ON pr.product_id = st.product_id
		WHERE ` + strings.Join(whereCol, " AND "), whereVal...)

	if err != nil {
		logger(err)
	    return
	}
	defer rows.Close()

	for rows.Next() {
		search_product := SearchProduct{}
		if err := rows.Scan(&search_product.ProductId, &search_product.ProductName, &search_product.Price, &search_product.CreatedDatetime, &search_product.UpdatedDatetime, &search_product.StockId, &search_product.Quantity, &search_product.StockCreatedDatetime, &search_product.StockUpdatedDatetime); err != nil {
			logger(err)
			return search_products, err
		}
		search_products.SearchProducts = append(search_products.SearchProducts, search_product)
	}
	return search_products, err
}



func retrieveUser(id int64) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select user_id, username, email, created_datetime, updated_datetime from users where user_id = ?", id).Scan(&user.UserId, &user.Username, &user.Email, &user.CreatedDatetime, &user.UpdatedDatetime)
	fmt.Println("retrieveUser")
	return
}
func checkExistingUsernameOrEmail(user *User) (isUsernameOrEmailExist bool) {
	Db.QueryRow("select user_id from users where username = ? or email = ?", user.Username, user.Email).Scan(&user.UserId)
	isUsernameOrEmailExist = user.UserId > 0
	return
}
func (user *User) createUser() (err error) {
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)
    password_hash, _ := HashPassword(user.Password)
 	res, err := Db.Exec("insert into users (username, email, created_datetime, updated_datetime, password) values (?, ?, ?, ?, ?)", user.Username, user.Email, dt, dt, password_hash)
    if err != nil {
    	logger(err)
        return
    }
	user.UserId, err = res.LastInsertId()
	if err != nil {
		logger(err)
	    return
	}
	return
}
func (user *User) updateUser() (err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
	_, err = Db.Exec("update users set username = ?, email = ?, updated_datetime = ? where user_id = ?", user.Username, user.Email, dt, user.UserId)
	return
}
func (user *User) deleteUser() (err error) {
	_, err = Db.Exec("delete from users where user_id = ?", user.UserId)
	return
}
func (user *User) changePassword() (err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
    password_hash, _ := HashPassword(user.Password)
	_, err = Db.Exec("update users set password = ?, updated_datetime = ? where user_id = ?", password_hash, dt, user.UserId)
	return
}
func searchUsers(search_user *SearchUser) (search_users SearchUsers, err error) {
	w := map[string]interface{}{"username": search_user.Username, "email": search_user.Email}
	var whereVal []interface{}
	var whereCol []string
	for k, v := range w {
	    whereVal = append(whereVal, fmt.Sprintf("%%%s%%", v))
	    whereCol = append(whereCol, fmt.Sprintf("%s like ?", k))
	}
	rows, err := Db.Query(`
		SELECT user_id, username, email, created_datetime, updated_datetime -- password
		FROM users WHERE ` + strings.Join(whereCol, " OR "), whereVal...)
	if err != nil {
		logger(err)
	    return
	}
	defer rows.Close()

	for rows.Next() {
		search_user := SearchUser{}
		if err := rows.Scan(&search_user.UserId, &search_user.Username, &search_user.Email, &search_user.CreatedDatetime, &search_user.UpdatedDatetime); err != nil {
			logger(err)
			return search_users, err
		}
		search_users.SearchUsers = append(search_users.SearchUsers, search_user)
	}
	return search_users, err
}





func retrieveOrder(id int64) (order Order, err error) {
	order = Order{}
	err = Db.QueryRow("select order_id, order_datetime, checkout_datetime, user_id from orders where order_id = ?", id).Scan(&order.OrderId, &order.OrderDatetime, &order.CheckoutDatetime, &order.UserId)
	fmt.Println("retrieveOrder")
	return
}
func (order *Order) createOrder() (err error) {
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)
 	res, err := Db.Exec("insert into orders (order_datetime, checkout_datetime, user_id) values (?, ?, ?)", dt, dt, order.UserId)
    if err != nil {
    	logger(err)
        return
    }
	order.OrderId, err = res.LastInsertId()
	if err != nil {
		logger(err)
	    return
	}
	return
}
func (order *Order) updateOrder() (err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
	odt := dt
	cdt := dt
	if order.OrderDatetime != "" {
		odt = order.OrderDatetime
	}
	if order.CheckoutDatetime != "" {
		cdt = order.CheckoutDatetime
	}
	_, err = Db.Exec("update orders set order_datetime = ?, checkout_datetime = ?, user_id = ? where order_id = ?", odt, cdt, order.UserId, order.OrderId)
	return
}
func (order *Order) checkoutOrder() (err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
	_, err = Db.Exec("update orders set checkout_datetime = ? where order_id = ?", dt)
	return
}
func (order *Order) deleteOrder() (err error) {
	_, err = Db.Exec("delete from orders where order_id = ?", order.OrderId)
	return
}





func retrieveOrderDetail(id int64) (order_detail OrderDetail, err error) {
	order_detail = OrderDetail{}
	err = Db.QueryRow("select order_detail_id, order_id, product_id, quantity, price, created_datetime, updated_datetime from order_details where order_detail_id = ?", id).Scan(&order_detail.OrderDetailId, &order_detail.OrderId, &order_detail.ProductId, &order_detail.Quantity, &order_detail.Price, &order_detail.CreatedDatetime, &order_detail.UpdatedDatetime)
	fmt.Println("retrieveOrderDetail")
	return
}
// normally order detail created after create order
func (order_detail *OrderDetail) createOrderDetail() (err error) {
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)
 	res, err := Db.Exec("insert into order_details (order_id, product_id, quantity, price, created_datetime, updated_datetime) values (?, ?, ?, ?, ?)", order_detail.OrderId, order_detail.ProductId, order_detail.Quantity, order_detail.Price, dt, dt)
    if err != nil {
    	logger(err)
        return
    }
	order_detail.OrderDetailId, err = res.LastInsertId()
	if err != nil {
		logger(err)
	    return
	}
	return
}
// normally updateOrderDetail is enough with updateQuantityOrderDetail (only quantity & price can be changed), but for wholeness of CRUD, this function used
func (order_detail *OrderDetail) updateOrderDetail() (err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)

	total_price := order_detail.Price
	if total_price == 0 {
		product := Product{}
		err = Db.QueryRow("select product_id, price from products where product_id = ?", order_detail.ProductId).Scan(&product.ProductId, &product.Price)
	    if err != nil {
	    	logger(err)
	        return
	    }
	    total_price = float64(order_detail.Quantity) * product.Price
	}

	_, err = Db.Exec("update order_details set order_id = ?, product_id = ?, quantity = ?, price = ?, updated_datetime = ? where order_detail_id = ?", order_detail.OrderId, order_detail.ProductId, order_detail.Quantity, total_price, dt, order_detail.OrderDetailId)
	return
}
func (order_detail *OrderDetail) updateQuantityOrderDetail() (err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
	total_price := order_detail.Price
	if total_price == 0 {
		// it's better to save current price when customer at the moment agree the price that day
		product := Product{}
		err = Db.QueryRow("select product_id, price from products where product_id = ?", order_detail.ProductId).Scan(&product.ProductId, &product.Price)
	    if err != nil {
	    	logger(err)
	        return
	    }
	    total_price = float64(order_detail.Quantity) * product.Price
	}
	_, err = Db.Exec("update order_details set quantity = ?, price = ?, updated_datetime = ? where order_detail_id = ?", order_detail.Quantity, total_price, dt)
	return
}
func (order_detail *OrderDetail) deleteOrderDetail() (err error) {
	_, err = Db.Exec("delete from order_details where order_detail_id = ?", order_detail.OrderDetailId)
	return
}
func searchOrderDetails(search_order_detail *SearchOrderDetail) (search_order_details SearchOrderDetails, err error) {
	w := map[string]interface{}{"ord.user_id": search_order_detail.UserId}
	var whereVal []interface{}
	var whereCol []string
	for k, v := range w {
	    // whereVal = append(whereVal, fmt.Sprintf("%%%s%%", v))
	    // whereCol = append(whereCol, fmt.Sprintf("%s like ?", k))
	    whereVal = append(whereVal, v)
	    whereCol = append(whereCol, fmt.Sprintf("%s = ?", k))
	}
	rows, err := Db.Query(`
		SELECT od.order_detail_id, od.order_id, od.product_id, od.quantity, od.price, 
			od.created_datetime, od.updated_datetime, 
			ord.order_datetime, ord.checkout_datetime, ord.user_id, us.username
		FROM order_details od
		INNER JOIN orders ord ON od.order_id = ord.order_id
		INNER JOIN users us ON ord.user_id = us.user_id
		WHERE ` + strings.Join(whereCol, " OR "), whereVal...)
	if err != nil {
		logger(err)
	    return
	}
	defer rows.Close()

	for rows.Next() {
		search_order_detail := SearchOrderDetail{}
		if err := rows.Scan(&search_order_detail.OrderDetailId, &search_order_detail.OrderId, &search_order_detail.ProductId, &search_order_detail.Quantity, &search_order_detail.Price, &search_order_detail.CreatedDatetime, &search_order_detail.UpdatedDatetime, &search_order_detail.OrderDatetime, &search_order_detail.CheckoutDatetime, &search_order_detail.UserId, &search_order_detail.Username); err != nil {
			logger(err)
			return search_order_details, err
		}
		search_order_details.SearchOrderDetails = append(search_order_details.SearchOrderDetails, search_order_detail)
	}
	return search_order_details, err
}



func retrieveStock(id int64) (stock Stock, err error) {
	stock = Stock{}
	err = Db.QueryRow("select stock_id, product_id, quantity, created_datetime, updated_datetime from stocks where stock_id = ?", id).Scan(&stock.StockId, &stock.ProductId, &stock.Quantity, &stock.CreatedDatetime, &stock.UpdatedDatetime)
	fmt.Println("retrieveStock")
	return
}
func (stock *Stock) createStock() (err error) {
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)
 	res, err := Db.Exec("insert into stocks (product_id, quantity, created_datetime, updated_datetime) values (?, ?, ?, ?)", stock.ProductId, stock.Quantity, dt, dt)
    if err != nil {
    	logger(err)
        return
    }
	stock.StockId, err = res.LastInsertId()
	if err != nil {
		logger(err)
	    return
	}
	return
}
func (stock *Stock) updateStock() (err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
	_, err = Db.Exec("update stocks set product_id = ?, quantity = ?, updated_datetime = ? where stock_id = ?", stock.ProductId, stock.Quantity, dt, stock.StockId)
	return
}
func (stock *Stock) deleteStock() (err error) {
	_, err = Db.Exec("delete from stocks where stock_id = ?", stock.StockId)
	return
}



func retrieveOrderTemplate(id int64) (order_template OrderTemplate, err error) {
	order_template = OrderTemplate{}
	err = Db.QueryRow("select order_template_id, product_id, order_id, quantity, price, user_id, created_datetime, updated_datetime from order_templates where order_template_id = ?", id).Scan(&order_template.OrderTemplateId, &order_template.ProductId, &order_template.OrderId, &order_template.Quantity, &order_template.Price, &order_template.UserId, &order_template.CreatedDatetime, &order_template.UpdatedDatetime)
	return
}
func (order_template *OrderTemplate) createOrderTemplate() (err error) {
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)
	total_price := order_template.Price
	if total_price == 0 {
		product := Product{}
		err = Db.QueryRow("select product_id, price from products where product_id = ?", order_template.ProductId).Scan(&product.ProductId, &product.Price)
	    if err != nil {
	    	logger(err)
	        return
	    }
	    total_price = float64(order_template.Quantity) * product.Price
	}
	// order_template.CartPhaseType = Carted
 	res, err := Db.Exec("insert into order_templates (product_id, order_id, quantity, price, user_id, cart_phase_type, created_datetime, updated_datetime ) values (?, ?, ?, ?, ?, ?, ?, ?)", order_template.ProductId, order_template.OrderId, order_template.Quantity, total_price, order_template.UserId, order_template.CartPhaseType, dt, dt)
    if err != nil {
    	logger(err)
        return
    }
	order_template.OrderTemplateId, err = res.LastInsertId()
	if err != nil {
		logger(err)
	    return
	}
	return
}
func (order_template *OrderTemplate) updateOrderTemplate() (err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
	total_price := order_template.Price
	if total_price == 0 {
		product := Product{}
		err = Db.QueryRow("select product_id, price from products where product_id = ?", order_template.ProductId).Scan(&product.ProductId, &product.Price)
	    if err != nil {
	    	logger(err)
	        return
	    }
	    total_price = float64(order_template.Quantity) * product.Price
	}
	_, err = Db.Exec("update order_templates set product_id = ?, order_id = ?, quantity = ?, price = ?, user_id = ?, cart_phase_type = ?, updated_datetime = ? where order_template_id = ?", order_template.ProductId, order_template.OrderId, order_template.Quantity, total_price, order_template.UserId, order_template.CartPhaseType, dt, order_template.OrderTemplateId)
	return
}
func (order_template *OrderTemplate) deleteOrderTemplate() (err error) {
	_, err = Db.Exec("delete from order_templates where order_template_id = ?", order_template.OrderTemplateId)
	return
}
func searchOrderTemplates(search_order_template *SearchOrderTemplate) (search_order_templates SearchOrderTemplates, err error) {
	w := map[string]interface{}{"ot.user_id": search_order_template.UserId}
	var whereVal []interface{}
	var whereCol []string
	for k, v := range w {
	    // whereVal = append(whereVal, fmt.Sprintf("%%%s%%", v))
	    // whereCol = append(whereCol, fmt.Sprintf("%s like ?", k))
	    whereVal = append(whereVal, v)
	    whereCol = append(whereCol, fmt.Sprintf("%s = ?", k))
	}
	rows, err := Db.Query(`
		SELECT ot.order_template_id, ot.order_id, ot.product_id, ot.quantity, ot.price, 
			ot.created_datetime, ot.updated_datetime, 
			ot.user_id, us.username, 
			CASE
			    WHEN cart_phase_type = 0 THEN "Carted"
			    WHEN cart_phase_type = 1 THEN "Ordered"
			    WHEN cart_phase_type = 2 THEN "Canceled"
			    ELSE "Unknown"
			END cart_phase_type1
		FROM order_templates ot
		INNER JOIN users us ON ot.user_id = us.user_id
		WHERE ` + strings.Join(whereCol, " OR "), whereVal...)
	if err != nil {
		logger(err)
	    return
	}
	defer rows.Close()

	for rows.Next() {
		search_order_template := SearchOrderTemplate{}
		if err := rows.Scan(&search_order_template.OrderTemplateId, &search_order_template.OrderId, &search_order_template.ProductId, &search_order_template.Quantity, &search_order_template.Price, &search_order_template.CreatedDatetime, &search_order_template.UpdatedDatetime, &search_order_template.UserId, &search_order_template.Username, &search_order_template.CartPhaseType); err != nil {
			logger(err)
			return search_order_templates, err
		}
		search_order_templates.SearchOrderTemplates = append(search_order_templates.SearchOrderTemplates, search_order_template)
	}
	return search_order_templates, err
}




// BUSINESS RULES MODEL FUNCTIONS
func getCartItems(cart_item *CartItem) (cart_items CartItems, err error) {
	rows, err := Db.Query(`
		SELECT ot.order_template_id, 
			ot.product_id, 
			ot.order_id, 
			ot.quantity, 
			ot.price, 
			ot.user_id, 
			ot.cart_phase_type, 
			ot.created_datetime, 
			ot.updated_datetime,
			pr.product_name,
			pr.price product_price,
			us.username
		FROM order_templates ot
		INNER JOIN products pr ON ot.product_id = pr.product_id
		INNER JOIN users us ON ot.user_id = us.user_id
		WHERE ot.cart_phase_type = ? AND ot.user_id = ?
		`, Carted, cart_item.UserId) // cart_phase_type 0 = cart

	if err != nil {
		logger(err)
	    return
	}
	defer rows.Close()
	for rows.Next() {
		cart_item := CartItem{}
		if err := rows.Scan(&cart_item.OrderTemplateId, &cart_item.ProductId, &cart_item.OrderId, &cart_item.Quantity, &cart_item.Price, &cart_item.UserId, &cart_item.CartPhaseType, &cart_item.CreatedDatetime, &cart_item.UpdatedDatetime, &cart_item.ProductName, &cart_item.ProductPrice , &cart_item.Username ); err != nil {
			logger(err)
			return cart_items, err
		}
		cart_items.CartItems = append(cart_items.CartItems, cart_item)
	}
	return cart_items, err
}


func (cart_item *CartItem) createCartItem() (msg string, err error) {
	msg = "Create Cart Item Failed"
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)

	// check exists
	product, err := retrieveProduct(cart_item.ProductId)
	if err != nil {
		msg = "Product Not Found"
		return
	}
	if product.ProductId == 0 {
		msg = "Product Not Found"
		return
	}
	user, err := retrieveUser(cart_item.UserId)
	if err != nil {
		msg = "User Not Found"
		return
	}
	if user.UserId == 0 {
		msg = "User Not Found"
		return
	}

	total_price := cart_item.Price
	if total_price == 0 {
		product := Product{}
		err = Db.QueryRow("select product_id, price from products where product_id = ?", cart_item.ProductId).Scan(&product.ProductId, &product.Price)
	    if err != nil {
	    	logger(err)
	        return
	    }
	    total_price = float64(cart_item.Quantity) * product.Price
	}

	cart_item.CartPhaseType = Carted
 	res, err := Db.Exec("insert into order_templates (product_id, order_id, quantity, price, user_id, cart_phase_type, created_datetime, updated_datetime ) values (?, ?, ?, ?, ?, ?, ?, ?)", cart_item.ProductId, cart_item.OrderId, cart_item.Quantity, total_price, cart_item.UserId, cart_item.CartPhaseType, dt, dt)
    if err != nil {
    	logger(err)
        return
    }
	cart_item.OrderTemplateId, err = res.LastInsertId()
	if err != nil {
		logger(err)
	    return
	}
	return
}
func (cart_item *CartItem) updateCartItem() (msg string, err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
	if cart_item.OrderTemplateId == 0 {
		msg = "Invalid Order Template Id"
		return
	}

	// check exists
	order_template, err := retrieveOrderTemplate(cart_item.OrderTemplateId)
	if err != nil {
		msg = "Cart is Empty"
		return
	}
	if order_template.OrderTemplateId == 0 {
		msg = "Product Not Found 2"
		return
	}

	// If still in session Cart, order_templates should have order_id = 0
	// after checkout, order_templates copied to orders, orders.checkout_datetime is set
	if order_template.OrderId > 0 {
		msg = "This Item has been Checkout"
		return
	}
	if order_template.CartPhaseType != Carted {
		msg = "Only Cart Items Can be Updated"
		return
	}

	// check exists
	product, err := retrieveProduct(cart_item.ProductId)
	if err != nil {
		msg = "Product Not Found"
		return
	}
	if product.ProductId == 0 {
		msg = "Product Not Found"
		return
	}

	user, err := retrieveUser(cart_item.UserId)
	if err != nil {
		msg = "User Not Found"
		return
	}
	if user.UserId == 0 {
		msg = "User Not Found"
		return
	}

	total_price := cart_item.Price
	if total_price == 0 {
		product := Product{}
		err = Db.QueryRow("select product_id, price from products where product_id = ?", cart_item.ProductId).Scan(&product.ProductId, &product.Price)
	    if err != nil {
	    	logger(err)
	        return
	    }
	    total_price = float64(cart_item.Quantity) * product.Price
	}
	// cart_item.CartPhaseType = Carted
	_, err = Db.Exec("update order_templates set quantity = ?, price = ?, updated_datetime = ? where order_template_id = ? AND cart_phase_type = ?", cart_item.Quantity, total_price, dt, cart_item.OrderTemplateId, Carted)
	return
}
func (cart_item *CartItem) updateCartItemQuantity() (msg string, err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)
	if cart_item.OrderTemplateId == 0 {
		msg = "Invalid Order Template Id"
		return
	}
	if cart_item.Quantity < 0 { // 0 if user want to delete item from cart, alternative using delete
		msg = "Invalid Quantity Value"
		return
	}

	// check exists
	order_template, err := retrieveOrderTemplate(cart_item.OrderTemplateId)
	if err != nil {
		msg = "Cart is Empty"
		return
	}
	if order_template.OrderTemplateId == 0 {
		msg = "Order Template Not Found"
		return
	}

	// If still in session Cart, order_templates should have order_id = 0
	// after checkout, order_templates copied to orders, orders.checkout_datetime is set
	if order_template.OrderId > 0 {
		msg = "This Item has been Checkout, Can not be Changed"
		return
	}
	if order_template.CartPhaseType != Carted {
		msg = "Only Cart Items Can be Updated"
		return
	}

	order_template.Quantity = cart_item.Quantity
	total_price := 0.00 // order_template.Price
	if total_price == 0 {
		product := Product{}
		err = Db.QueryRow("select product_id, price from products where product_id = ?", order_template.ProductId).Scan(&product.ProductId, &product.Price)
	    if err != nil {
	    	logger(err)
	        return
	    }
	    total_price = float64(order_template.Quantity) * product.Price
	}
	// order_template.CartPhaseType = Carted
	_, err = Db.Exec("update order_templates set quantity = ?, price = ?, updated_datetime = ? where order_template_id = ? AND cart_phase_type = ?", order_template.Quantity, total_price, dt, order_template.OrderTemplateId, Carted)
	return
}

func (cart_item *CartItem) checkout(ctx context.Context) (msg string, err error) {
	msg = "Checkout Failed"
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)

	// // check exists
	// product, err := retrieveProduct(cart_item.ProductId)
	// if err != nil {
	// 	msg = "Product Not Found"
	// 	return
	// }
	// if product.ProductId == 0 {
	// 	msg = "Product Not Found"
	// 	return
	// }
	user, err := retrieveUser(cart_item.UserId)
	if err != nil {
		msg = "User Not Found"
		return
	}
	if user.UserId == 0 {
		msg = "User Not Found"
		return
	}

    // Create a helper function for preparing failure results.
    fail := func(err error) (string, error) {
    	msg = fmt.Sprintf("Checkout: %v", err)
        return msg, fmt.Errorf("Checkout: %v", err)
    }
    tx, err := Db.BeginTx(ctx, nil)
    if err != nil {
        return fail(err)
    }
    defer tx.Rollback()


    // validate stock first
	rows, err := Db.Query(`
		SELECT ot.order_template_id, ot.product_id, ot.quantity, ot.price, ot.user_id, 
			pr.product_name, pr.price product_price
		FROM order_templates ot
		INNER JOIN products pr ON ot.product_id = pr.product_id
		WHERE order_id = 0 AND cart_phase_type = ? AND user_id = ? `, Carted, cart_item.UserId)
	if err != nil {
		// logger(err)
		return fail(err)
	}
	defer rows.Close()

	fmt.Println("checkout")

	cart_items := CartItems{}
	for rows.Next() {
		ci := CartItem{}
		if err := rows.Scan(&ci.OrderTemplateId, &ci.ProductId, &ci.Quantity, &ci.Price, &ci.UserId, 
			&ci.ProductName, &ci.ProductPrice); err != nil {
			// logger(err)
            return fail(fmt.Errorf("Problem in Database"))
		}
		cart_items.CartItems = append(cart_items.CartItems, ci)

	    // Confirm that product stock is enough for every item in cart/order_template.
	    // using group by to ensure 1 row only
	    var enough bool
		if err = tx.QueryRowContext(ctx, `
				SELECT (st.qty >= ?) is_enough
				FROM products pr
				INNER JOIN (
					SELECT product_id, SUM(quantity) qty
					FROM stocks
					WHERE product_id = ?
					GROUP BY product_id
				) st ON pr.product_id = st.product_id
				WHERE pr.product_id = ?
			`, ci.Quantity, ci.ProductId, ci.ProductId).Scan(&enough); err != nil {
	        if err == sql.ErrNoRows {
	            return fail(fmt.Errorf("Product Not Found"))
	        }
	        return fail(err)
	    }
	    if !enough {
	        return fail(fmt.Errorf("Not Enough Product Stock - %v", ci.ProductName))
	    }

	    fmt.Println(enough)
	}

	// create 1 order first
    // Create a new row in the order table.
    result, err := tx.ExecContext(ctx, "INSERT INTO orders (order_datetime, checkout_datetime, user_id) VALUES (?, ?, ?)",
        dt, dt, cart_item.UserId)
    if err != nil {
        return fail(err)
    }
    orderID, err := result.LastInsertId()
    if err != nil {
        return fail(err)
    }

    fmt.Println("orderId")
    fmt.Println(orderID)

    fmt.Println("cart items")
    fmt.Println(cart_items)

    // create order_details
    // update stock
    // // update order_template // better to update all cart items in one go, select where user_id, cart_phase_id = Carted, order_id = 0
    // TODO: if possible, implement additional checking if ot.price/ci.price = 0 then use product.price
	for ii, ci := range cart_items.CartItems {
		fmt.Println(ii)
		fmt.Println(ci)
	 	_, err := Db.Exec(`
	 		INSERT INTO order_details (
	 			order_id, product_id, quantity, price, created_datetime, updated_datetime
	 		) VALUES (?, ?, ?, ?, ?, ?)`, orderID, ci.ProductId, ci.Quantity, ci.Price, dt, dt)
	    if err != nil {
	    	// logger(err)
	        return fail(err)
	    }

		// Update the product stocks to reduce the quantity
	    _, err = tx.ExecContext(ctx, "UPDATE stocks SET quantity = quantity - ? WHERE product_id = ?",
	        ci.Quantity, ci.ProductId)
	    if err != nil {
	        return fail(err)
	    }
	}

	// Update the all order_template/cart rows to change cart_phase_type
    _, err = tx.ExecContext(ctx, `
    	UPDATE order_templates SET cart_phase_type = ?, order_id = ?
    	WHERE user_id = ? AND cart_phase_type = ?
    	`, Ordered, orderID, cart_item.UserId, Carted)
    if err != nil {
        return fail(err)
    }

    // Commit the transaction.
    if err = tx.Commit(); err != nil {
        return fail(err)
    }

	return
}
func (cart_item *CartItem) cancelCartOrder() (msg string, err error) {
	var updated_datetime = time.Now()
	dt := updated_datetime.Format(time.RFC3339)

	if cart_item.UserId == 0 {
		msg = "Invalid User Id"
		return
	}

	_, err = Db.Exec("update order_templates set cart_phase_type = ?, updated_datetime = ? where user_id = ? AND cart_phase_type = ?", Canceled, dt, cart_item.UserId, Carted)
	return
}


func getOrderItems(order_item *OrderItem) (order_items OrderItems, err error) {
	rows, err := Db.Query(`
		SELECT ord.order_id, 
			ord.user_id,
			ord.order_datetime,
			od.order_detail_id,
			od.product_id, 
			od.quantity, 
			od.price,
			pr.product_name,
			us.username
		FROM orders ord
		INNER JOIN order_details od ON ord.order_id = od.order_id
		INNER JOIN products pr ON od.product_id = pr.product_id
		INNER JOIN users us ON ord.user_id = us.user_id
		WHERE ord.user_id = ?
		`, order_item.UserId) // cart_phase_type 0 = cart

	if err != nil {
		logger(err)
	    return
	}
	defer rows.Close()
	for rows.Next() {
		order_item := OrderItem{}
		if err := rows.Scan(&order_item.OrderId, &order_item.UserId, &order_item.OrderDatetime, &order_item.OrderDetailId, &order_item.ProductId, &order_item.Quantity, &order_item.Price, &order_item.ProductName, &order_item.Username); err != nil {
			logger(err)
			return order_items, err
		}
		order_items.OrderItems = append(order_items.OrderItems, order_item)
	}
	return order_items, err
}


func loggingRequest(log Logging) {
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)
 	_, err := Db.Exec(`
 		INSERT INTO logs (url, remote_address, created_datetime) 
 		VALUES (?, ?, ?)
 	`, log.Url, log.RemoteAddress, dt)
    if err != nil {
    	logger(err)
        return
    }
	return
}


func (login_data *LoginData) login() (isUserValid bool, err error) {
	var created_datetime = time.Now()
	dt := created_datetime.Format(time.RFC3339)
	isUserValid = false
	user := new(User)
	err = Db.QueryRow(`
		SELECT us.user_id, us.username, us.email, us.password
		FROM users us
		WHERE (us.username = ? OR us.email = ?)
		ORDER BY us.created_datetime DESC
		`, login_data.Username, login_data.Email).Scan(&user.UserId, &user.Username, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return		
	}
    if !CheckPasswordHash(login_data.Password, user.Password) {
    	return
    }

    // session valid for 30 minutes
	err = Db.QueryRow(`
		SELECT ss.session_code, ss.created_datetime
		FROM users us
		INNER JOIN sessions ss ON us.user_id = ss.user_id
		WHERE us.user_id = ? AND DATE_ADD(ss.created_datetime, INTERVAL 30 MINUTE) > NOW()
		ORDER BY ss.created_datetime DESC
		`, user.UserId).Scan(&login_data.SessionCode, &login_data.SessionCreatedDatetime)
	if err == sql.ErrNoRows {
		// create new session
		login_data.SessionCreatedDatetime = dt
		login_data.SessionCode = uuid.New().String()
	 	_, err = Db.Exec("INSERT INTO sessions (session_code, user_id, created_datetime) values (?, ?, ?)", login_data.SessionCode, user.UserId, dt)
	    if err != nil {
	    	logger(err)
	        return
	    }
	}

    isUserValid = true
	return
}


func (login_data *LoginData) logout() (isUserValidLogout bool, err error) {
	isUserValidLogout = false
	_, err = Db.Exec("UPDATE sessions SET is_expired = 1 WHERE session_code = ?", login_data.SessionCode)
    if err != nil {
    	logger(err)
        return
    }
    isUserValidLogout = true
	return
}


func (session_data *SessionData) checkSessionDb() (isSessionExpired bool, err error) {
	isSessionExpired = true

	err = Db.QueryRow(`
		SELECT us.user_id, ss.created_datetime
		FROM users us
		INNER JOIN sessions ss ON us.user_id = ss.user_id
		WHERE ss.session_code = ? AND DATE_ADD(ss.created_datetime, INTERVAL 30 MINUTE) > NOW()
		ORDER BY ss.created_datetime DESC
		`, session_data.SessionCode).Scan(&session_data.UserId, &session_data.SessionCreatedDatetime)

	if err == sql.ErrNoRows {
		return
	}
    isSessionExpired = false
	return
}




