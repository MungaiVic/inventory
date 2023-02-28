# inventory

- A simple CRUD API for inventory

## Libraries used

- These include:
  - gorm <https://gorm.io/>
  - fiber(V2) <https://docs.gofiber.io/>
  - validate <https://pkg.go.dev/github.com/go-playground/validator>

## Usage

- The API is structured as follows:
  - `GET /api/v1/items/` - Get all inventory
  - `GET /api/v1/items/get_item/{id}` - Get inventory by id
  - `POST /api/v1/items/create_item/` - Create new inventory
  - `PATCH /api/v1/items/update_item/{id}` - Update inventory by id
  - `DELETE /api/v1/items/delete_item/{id}` - Delete inventory by id

## Running the project

- Clone the repository
- Run `go mod tidy` to install dependencies
- cd into the project directory `cd cmd`
- Run `go run main.go` to start the server

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
