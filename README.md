# inventory

- A simple CRUD API for inventory

## Libraries used

- These include:
  - gorm <https://gorm.io/>
  - fiber(V2) <https://docs.gofiber.io/>

## Usage

- The API is structured as follows:
  - `GET /api/v1/items/` - Get all inventory
  - `GET /api/v1/items/get_item/{id}` - Get inventory by id
  - `POST /api/v1/items/inventory/` - Create new inventory
  - `PATCH /api/v1/items/inventory/` - Update inventory by id
  - `DELETE /api/v1/items/inventory/` - Delete inventory by id
  - `GET /api/v1/users/` - Get all users
  - `GET /api/v1/items/get_user/` - Get user by id
  - `POST /api/v1/items/create_user/` - Create new user

## Running the project

- Clone the repository
- Run `go mod tidy` to install dependencies
- cd into the project directory `cd cmd`
- Run `go run main.go` to start the server
- The server will run on `localhost:5000`

<!-- Highlight block -->

<!-- > To migrate the database, run `go run main.go migrate` -->

## Sample request

- `POST /api/v1/items/create_item/`

```json
{
 "name": "test",
 "reorderlvl": 1,
 "quantity": 1,
 "price": 1
}
```

- `PATCH /api/v1/items/update_item/`

```json
{
 "ID": 1,
 "name": "test",
 "reorderlvl": 1,
 "quantity": 1,
 "price": 1
}
```

- `DELETE /api/v1/items/delete_item/{id}`
