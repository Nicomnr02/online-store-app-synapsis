
# Synapsis Software Engineer Challenge
This challenge is to make a service of **Online Store Application** using Go or Java, exactly i use Go to land off this project.
_Synapsis challenge of internship recruitment._ 
(2081015@unai.edu)

This API is a pretty basic implementation of an online(e-commerce) store.

- Perform basic CRUD operations (based on challenge requiring).
- Search on a predefined database of products. 
- Only Authenticated users can Add, Update and Delete products from database.
- Authentication is based on JWT(JSON web Tokens) Tokens.
- API is backed by a predefined PostgreSQL.
- Both of Service - application and database (PostgreSQL) is containerized in Docker.



## Directory Structure

Basicly, i was using factory pattern to implement the Dependency injection code. I arrange the folder like the folder raw structure below.
```
api/                                   *Contains all handler incoming request - make by injected services.
    |- users.go           
    |- products.go 
    |- carts.go
    |- etc.go
middleware/                            *Contains all handler to validate incoming request  
    |- auth.go (session validator)
    |- method.go (method validator)
model/                                 *Contains all entity of project needs
    |- cart.go           
    |- category.go 
    |- credential.go
    |- etc.go                          
repositories/                          *Contains all entity represent table in db.
    |- products.go           
    |- carts.go 
    |- etc.go
services/                              *Contains all service functions (interface of repositories).
    |- users.go  (login,register, etc)          
    |- carts.go 
    |- etc.go      
utils/                                 *Contains setup functions (db or other third party services).
    |- users.go           
    |- carts.go 
    |- etc.go    
main.go                                *Entry point of the API
  
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DATABASE_URL` postgres://postgres:secret@synapsis-db:5432/postgres 




## API Reference
I put some basic endpoints for the application's use of the service/API.

#### Register

```http
  POST /v1/user/register
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `username` | `string` |  **JSON** |
| `password` | `string` |  **JSON** |



#### Login

```http
  POST /v1/user/login
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **JSON** |
| `password` | `string` | **JSON** |



#### Show All Products

```http
  GET /v1/user/show/products
```


#### Show Product by ID

```http
  GET /v1/user/show/product/on [?product_id=]
```

| Parameter | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `product_id` | `number` | **URL Param** |



#### Show Product by Category ID

```http
  GET /v1/user/show/productsByCategory/on [?category_id=]
```

| Parameter | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `category_id` | `number` | **URL Param** |



#### Show Categories

```http
  GET /v1/user/show/categories
```


#### Show Categories with products

```http
  GET /v1/user/show/categoriesWithProducts
```


#### Add item to cart

```http
  POST /v1/user/send/cart
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `product_id` | `int` | **JSON** |
| `total` | `int` | **JSON** |



#### Delete item from cart

```http
  DELETE /v1/user/delete/cart/on [?cart_id=]
```

| Parameter | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `cart_id` | `number` | **URL Param** |


#### Show all item (cart)

```http
  GET /v1/user/show/categoriesWithProducts
```


#### Bring cart to a transaction (unpaid) 

```http
  POST /v1/user/send/transaction
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `cart_id` | `int` | **JSON** |



#### Bring carts to transactions (unpaid)

```http
  POST /v1/user/send/transactions
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `list of cart_id` | `int` | **JSON** |



#### Pay selected transaction

```http
  PUT /v1/user/update/transaction
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `transaction_id`| `int` | **JSON** |



#### Pay selected transactions

```http
  PUT /v1/user/update/transaction
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `list of transaction_id`| `int` | **JSON** |



#### Cancel selected transactions

```http
  DELETE  /v1/user/delete/transaction
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `transaction_id`| `int` | **JSON** |



#### Show all transactions

```http
  GET /v1/user/show/transactions
```


