
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

```http
Response :
{
  "message": "Register Success",
    "data": {
        "ID": 0,
        "CreatedAt": "2023-05-08T07:26:23.759305098Z",
        "UpdatedAt": "2023-05-08T07:26:23.759305098Z",
        "DeletedAt": null,
        "id": 1,
        "username": "synapsis_user",
        "password": "123",
        "cash": 500000
    }
}
```



#### Login

```http
  POST /v1/user/login
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **JSON** |
| `password` | `string` | **JSON** |


```http
Response :
{
  "message": "Login Success",
    "data": {
        "ID": 0,
        "CreatedAt": "2023-05-08T09:11:00.790439469Z",
        "UpdatedAt": "2023-05-08T09:11:00.790439469Z",
        "DeletedAt": null,
        "id": 1,
        "user_id": 1,
        "session_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.pYhVGwis9I-dQ0Y2lFVzZxgCz_ceesK0fNY7hOsFHZk"
    }
}
```


#### Show All Products

```http
  GET /v1/user/show/products
```

```http
Response :
{
 "message": "Success get all products",
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2023-05-06T22:59:06.078666+07:00",
            "UpdatedAt": "2023-05-07T18:54:09.692321+07:00",
            "DeletedAt": null,
            "id": 44,
            "category_id": 7,
            "product_name": "product 4",
            "product_price": 35000,
            "product_stock": 10
        },
        {
            "ID": 0,
            "CreatedAt": "2023-05-06T22:59:06.078666+07:00",
            "UpdatedAt": "2023-05-06T22:59:06.078666+07:00",
            "DeletedAt": null,
            "id": 50,
            "category_id": 10,
            "product_name": "product 10",
            "product_price": 20000,
            "product_stock": 10
        },
        {
            "ID": 0,
            "CreatedAt": "2023-05-06T22:59:06.078666+07:00",
            "UpdatedAt": "2023-05-06T22:59:06.078666+07:00",
            "DeletedAt": null,
            "id": 45,
            "category_id": 8,
            "product_name": "product 5",
            "product_price": 30000,
            "product_stock": 10
        }...
}
```


#### Show Product by ID

```http
  GET /v1/user/show/product/on [?product_id=]
```

| Parameter | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `product_id` | `number` | **URL Param** |


```http
Response :
{
    "message": "Success get product by id",
    "data": {
        "ID": 0,
        "CreatedAt": "2023-05-08T09:12:40.929865Z",
        "UpdatedAt": "2023-05-08T09:12:40.929865Z",
        "DeletedAt": null,
        "id": 1,
        "category_id": 6,
        "product_name": "product 1",
        "product_price": 50000,
        "product_stock": 60
    }
}
```


#### Show Product by Category ID

```http
  GET /v1/user/show/productsByCategory/on [?category_id=]
```

| Parameter | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `category_id` | `number` | **URL Param** |


```http
Response :
{
    "message": "Success get products by category id",
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2023-05-08T09:12:40.929865Z",
            "UpdatedAt": "2023-05-08T09:12:40.929865Z",
            "DeletedAt": null,
            "id": 1,
            "category_id": 6,
            "product_name": "product 1",
            "product_price": 50000,
            "product_stock": 60
        },
        {
            "ID": 0,
            "CreatedAt": "2023-05-08T09:12:40.929865Z",
            "UpdatedAt": "2023-05-08T09:12:40.929865Z",
            "DeletedAt": null,
            "id": 2,
            "category_id": 6,
            "product_name": "product 2",
            "product_price": 45000,
            "product_stock": 50
        }
    ]
``` 


#### Show Categories

```http
  GET /v1/user/show/categories
```

```http
Response :
{
    "message": "Success get all categories",
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2023-05-06T22:56:37.318174+07:00",
            "UpdatedAt": "2023-05-06T22:56:37.318174+07:00",
            "DeletedAt": null,
            "id": 6,
            "category_type": "cat 1"
        },
        {
            "ID": 0,
            "CreatedAt": "2023-05-06T22:56:37.318174+07:00",
            "UpdatedAt": "2023-05-06T22:56:37.318174+07:00",
            "DeletedAt": null,
            "id": 7,
            "category_type": "cat 2"
        },
        {
            "ID": 0,
            "CreatedAt": "2023-05-06T22:56:37.318174+07:00",
            "UpdatedAt": "2023-05-06T22:56:37.318174+07:00",
            "DeletedAt": null,
            "id": 8,
            "category_type": "cat 3"
        }...
```


#### Show Categories with products

```http
  GET /v1/user/show/categoriesWithProducts
```

```http
Response :
 "message": "Success get categories with products",
    "data": [
        {
            "id": 6,
            "category_type": "cat 1",
            "products": [
                {
                    "ID": 0,
                    "CreatedAt": "2023-05-06T22:59:06.078666+07:00",
                    "UpdatedAt": "2023-05-07T19:25:42.748583+07:00",
                    "DeletedAt": null,
                    "id": 41,
                    "category_id": 6,
                    "product_name": "product 1",
                    "product_price": 50000,
                    "product_stock": 8
                },
                {
                    "ID": 0,
                    "CreatedAt": "2023-05-06T22:59:06.078666+07:00",
                    "UpdatedAt": "2023-05-07T19:25:42.750177+07:00",
                    "DeletedAt": null,
                    "id": 42,
                    "category_id": 6,
                    "product_name": "product 2",
                    "product_price": 45000,
                    "product_stock": 8
                }
            ]
        }...
```


#### Add item to cart

```http
  POST /v1/user/send/cart
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `product_id` | `int` | **JSON** |
| `total` | `int` | **JSON** |

```http
Response :
{
  "message": "Success add a cart",
    "data": {
        "ID": 0,
        "CreatedAt": "2023-05-08T08:58:15.6825263+07:00",
        "UpdatedAt": "2023-05-08T08:58:15.6825263+07:00",
        "DeletedAt": null,
        "id": 100,
        "user_id": 6,
        "product_id": 43,
        "total": 2,
        "price": 100000
    }
}
```


#### Delete item from cart

```http
  DELETE /v1/user/delete/cart/on [?cart_id=]
```

| Parameter | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `cart_id` | `number` | **URL Param** |


```http
Response :
{
    "message": "Success delete a cart",
    "data": {
        "ID": 0,
        "CreatedAt": "2023-05-08T09:34:57.503716956Z",
        "UpdatedAt": "2023-05-08T09:34:57.503716956Z",
        "DeletedAt": null,
        "id": 1,
        "user_id": 1,
        "product_id": 1,
        "total": 2,
        "price": 100000
    }
}
```


#### Show all item (cart)

```http
  GET /v1/user/show/categoriesWithProducts
```

```http
Response :
{
    "message": "Success show all carts by user ID",
    "data": {
        "id": 8,
        "user_id": 1,
        "carts": [
            {
                "ID": 0,
                "CreatedAt": "2023-05-08T09:34:57.503716Z",
                "UpdatedAt": "2023-05-08T09:34:57.503716Z",
                "DeletedAt": null,
                "id": 1,
                "user_id": 1,
                "product_id": 1,
                "total": 2,
                "price": 100000
            }...
        ]
    }
}
```


#### Bring cart to a transaction (unpaid) 

```http
  POST /v1/user/send/transaction
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `cart_id` | `int` | **JSON** |

```http
Response :
{
    "message": "Success create a transaction ",
    "data": {
        "ID": 0,
        "CreatedAt": "2023-05-08T09:36:48.649253069Z",
        "UpdatedAt": "2023-05-08T09:36:48.649253069Z",
        "DeletedAt": null,
        "id": 1,
        "user_id": 1,
        "cart_id": 1,
        "transaction_status": "not_paid"
    }
}
```


#### Bring carts to transactions (unpaid)

```http
  POST /v1/user/send/transactions
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `list of cart_id` | `int` | **JSON** |

```http
Response :
"message": "Success create transactions ",
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2023-05-08T08:59:09.8646638+07:00",
            "UpdatedAt": "2023-05-08T08:59:09.8646638+07:00",
            "DeletedAt": null,
            "id": 37,
            "user_id": 6,
            "cart_id": 99,
            "transaction_status": "not_paid"
        },
        {
            "ID": 0,
            "CreatedAt": "2023-05-08T08:59:09.8744386+07:00",
            "UpdatedAt": "2023-05-08T08:59:09.8744386+07:00",
            "DeletedAt": null,
            "id": 38,
            "user_id": 6,
            "cart_id": 98,
            "transaction_status": "not_paid"
        },
        {
            "ID": 0,
            "CreatedAt": "2023-05-08T08:59:09.8757307+07:00",
            "UpdatedAt": "2023-05-08T08:59:09.8757307+07:00",
            "DeletedAt": null,
            "id": 39,
            "user_id": 6,
            "cart_id": 100,
            "transaction_status": "not_paid"
        }
    ]
}
```


#### Pay selected transaction

```http
  PUT /v1/user/update/transaction
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `transaction_id`| `int` | **JSON** |

```http
Response :
{
    "message": "Success update transaction",
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2023-05-08T08:56:07.3602749+07:00",
        "DeletedAt": null,
        "id": 36,
        "user_id": 6,
        "cart_id": 0,
        "transaction_status": "paid"
    }
}
```


#### Pay selected transactions

```http
  PUT /v1/user/update/transaction
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `list of transaction_id`| `int` | **JSON** |

```http
Response :
{
    "message": "Success update choosen transaction",
    "data": [
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "2023-05-08T09:00:02.3101505+07:00",
            "DeletedAt": null,
            "id": 37,
            "user_id": 6,
            "cart_id": 0,
            "transaction_status": "paid"
        },
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "2023-05-08T09:00:02.3108695+07:00",
            "DeletedAt": null,
            "id": 38,
            "user_id": 6,
            "cart_id": 0,
            "transaction_status": "paid"
        },
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "2023-05-08T09:00:02.311616+07:00",
            "DeletedAt": null,
            "id": 39,
            "user_id": 6,
            "cart_id": 0,
            "transaction_status": "paid"
        }
    ]
}
```


#### Cancel selected transactions

```http
  DELETE  /v1/user/delete/transaction
```

| Parameter | Type     | Header Type                |
| :-------- | :------- | :------------------------- |
| `transaction_id`| `int` | **JSON** |

```http
Response :
  {
      "message": "Success delete transaction by user id",
      "data": null
  }
```


#### Show all transactions

```http
  GET /v1/user/show/transactions
```

```http
Response :
{
    "message": "Success get all transactions by user id",
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2023-05-08T07:59:08.608273+07:00",
            "UpdatedAt": "2023-05-08T07:59:08.608273+07:00",
            "DeletedAt": null,
            "id": 28,
            "user_id": 5,
            "cart_id": 89,
            "transaction_status": "not_paid"
        },...
    ]
}
```




