package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jszwec/csvutil"
)

type Item struct {
	Name      string  `json:"name"`
	ID        string  `json:"id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

var items = []Item{
	{Name: "jade pendant", ID: "1", Quantity: 17, UnitPrice: 15.40},
	{Name: "casual shoes", ID: "2", Quantity: 2, UnitPrice: 65.58},
	{Name: "fluorescent pen", ID: "3", Quantity: 119, UnitPrice: 21.77},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func postItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, i := range items {
		if i.ID == id {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")
	for index, i := range items {
		if i.ID == id {
			items = append(items[:index], items[index+1:]...)
		}
	}
}

func updateItem(c *gin.Context) {
	var newItem Item
	id := c.Param("id")

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	for x, i := range items {
		if i.ID == id {
			items[x].Name = newItem.Name
			items[x].Quantity = newItem.Quantity
			items[x].UnitPrice = newItem.UnitPrice
			c.IndentedJSON(http.StatusOK, items[x])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func exportCSV(c *gin.Context) {
	csvfile, err := os.Create("output.csv")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error with the CSV file creation"})
	}
	defer csvfile.Close()
	b, err := csvutil.Marshal(items)
	if err != nil {
		return
	}
	csvfile.Write(b)
	c.FileAttachment("output.csv", "itemlist.csv")
}

func main() {
	r := gin.Default()

	r.GET("/item", getItems)
	r.GET("/item/csv", exportCSV)
	r.GET("/item/:id", getItemByID)
	r.POST("/item", postItem)
	r.DELETE("/item/:id", deleteItem)
	r.PATCH("/item/:id", updateItem)

	r.Run() // listen and serve on port :8080
}
