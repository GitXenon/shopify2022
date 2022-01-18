# Shopify Backend Developer Intern Challenge - Summer 2022
A simple CRUD app built with Go and Gin. This let you update track your inventory, add new items, remove items, change items info and export to CSV.

## Run the application
```
go run main.go
```
This launches the application on port `:8080`.

Go in your browser to http://localhost:8080/item to receive a JSON with the items already present in the storage.
The prefered method is to use CURL or Postman to interact with the API.

### Postman
You can import `postman_collection.json` into Postman to have all the calls already in your workspace.

### Endpoints
`GET localhost:8080/item` Returns a json with the list of the items and a 200 code.

`POST localhost:8080/item` Post with a json body containing an id, name, quantity and unit_price. Returns the item added with a 201 code.

`GET localhost:8080/item/1` Returns a json with the item wanted, and a 200 code.

`DELETE localhost:8080/item/1` Delete an item by its ID. Returns 200 if item was removed.

`PATCH localhost:8080/item/2` Body needs to contain a json with name, quantity and unit_price. The info of this item will be changed. Returns 200.

`GET localhost:8080/item/csv` Returns a CSV file.