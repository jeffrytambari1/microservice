## Description

Tech Test - Advance Level Senior Level

## Repositories


## Installation

1. Create database go_microservice
2. Restore database
```sql
mysql -u mysql_username -ppassword go_microservice < install.sql
```
3. Change mysql username 'your_user' & password 'your_password' in ./data.go
4. Download library needed
```bash
$ go get "github.com/gofiber/fiber/v2"
$ go get "github.com/go-sql-driver/mysql"
$ go get "github.com/google/uuid"
$ go get "golang.org/x/crypto/bcrypt"

```
5. Tidy up
```bash
$ go mod tidy

```
6. Run
```bash
$ go run .

```
7. Open web browser, go to http://127.0.0.1:3000/
8. Import [Postman collection](https://learning.postman.com/docs/getting-started/importing-and-exporting-data/)
	nityo_microservice.postman_collection_221226_122034.json


## Running the app

```bash
$ go run .

```
http://127.0.0.1:3000/


## Read me
implementation:
- Golang
- Framework Fiber
- REST API
- Product Service
- Order Service (cart)
- Manual Auth Service
- CRUD
- Search
- Simple Cart System
- Mysql connection
- Manual Logging

### Session / Auth Service
- Only get product(retrieve) http://127.0.0.1:3000/product?id=7 that has authorization/session validation security, 
this is because not yet get verification from google to get ClientId to implement Oauth2
- Notice in postman this request has Authorization with type Bearer & Token
	example Token = 533be427-9800-48be-8c72-d9efaffd6d2f
- If expired, will show "Session has Expired"
- Session setting is in table "sessions"

### Login/Logout
- login to fetch new session_code, if already 30 minute, session_code will be renew
- logout to change is_expired
- login http://127.0.0.1:3000/login
- logout http://127.0.0.1:3000/logout

### CRUD
- Product http://127.0.0.1:3000/product
- User http://127.0.0.1:3000/user
- Order http://127.0.0.1:3000/order
- Order Detail http://127.0.0.1:3000/order_detail
- Order Template (or cart) http://127.0.0.1:3000/order_template
- Stock http://127.0.0.1:3000/stock

### Cart System
- Implementation of cart using table order_templates, 
which in the future will be lot of junk/unused data, 
because customer can add product to cart as many as he wants.
- If customer checkout, his cart data/items will copied to order_details.
- Not yet implement session, to get user (user_id) data
- Create Cart Item http://127.0.0.1:3000/create_cart_item
	it's like customer press 'add to cart this product'
- Cancel Cart Order http://127.0.0.1:3000/cancel_cart_order
	I think Cancel order is better in cart level, not after checkout.
	I think after checkout the order is immutable (can not be changed).
- Get Cart Items http://127.0.0.1:3000/get_cart_items?user_id=10
- Get Order Items http://127.0.0.1:3000/get_order_items?user_id=10
- Update Cart Item http://127.0.0.1:3000/update_cart_item
- Update Cart Item Quantity http://127.0.0.1:3000/update_cart_item_quantity
- Delete Cart Item http://127.0.0.1:3000/delete_cart_item?order_template_id=14
- Checkout http://127.0.0.1:3000/checkout

### Drawback
- there's no separation between Admin User & Ordinary User, because not yet implemented.

## Author

- [Jeffry Tambari](https://github.com/jeffrytambari1)
