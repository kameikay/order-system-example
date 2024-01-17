# Order System Example

This repository provides a straightforward example of an order system that supports REST, gRPC, and GraphQL on the same server. The system allows you to create orders and retrieve a list of existing orders.

## How to Run

1. **Clone this Repository:**
   ```bash
   git clone https://github.com/kameikay/order-system-example.git
   cd order-system-example
   ```

2. **Start Docker Containers:**
   ```bash
   docker-compose up -d
   ```

3. **Navigate to the 'cmd/ordersystem' Directory:**
   ```bash
   cd cmd/ordersystem
   ```

4. **Run the Application:**
   ```bash
   go run main.go wire_gen.go
   ```

## How to Use

### REST

#### Create Order

1. Open `api/create_order.http`
2. Click on the `Send Request` button (or modify the request fields)

#### Get Orders

1. Open `api/list_order.http`
2. Click on the `Send Request` button

### gRPC

#### Create Order

1. Use Evans CLI
   ```bash
   evans -r repl
   ```
   - Set the package:
     ```evans
     $ package pb
     ```
   - Set the service:
     ```evans
     $ service OrderService
     ```
   - Call the service:
     ```evans
     $ call CreateOrder
     ```

#### Get Orders

1. Use Evans CLI
   ```bash
   evans -r repl
   ```
   - Set the package:
     ```evans
     $ package pb
     ```
   - Set the service:
     ```evans
     $ service OrderService
     ```
   - Call the service:
     ```evans
     $ call ListOrders
     ```

### GraphQL

#### Create Order

1. Navigate to [http://localhost:8080](http://localhost:8080)
2. Run the following mutation query:
   ```graphql
   mutation createOrder {
   	createOrder(input: { id: "cccc", Price: 123, Tax: 23 }) {
   		id
   		Price
   		Tax
   		FinalPrice
   	}
   }
   ```

#### Get Orders

1. Navigate to [http://localhost:8080](http://localhost:8080)
2. Run the following query:
   ```graphql
   query listOrders {
   	listOrders {
   		id
   		Price
   		Tax
   		FinalPrice
   	}
   }
   ```