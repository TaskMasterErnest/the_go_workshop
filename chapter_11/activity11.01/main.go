/***
Mimic a customer order
An online e-commerce portal needs to accept customer orders over its web application
Customer will add items to their order.
Web application will need to take the JSON and add orders to the JSON
***/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// address struct
type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipcode"`
}

// item struct
type item struct {
	Name        string  `json:"itemname"`
	Description string  `json:"desc,omitempty"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

// order struct
type order struct {
	TotalPrice  float64 `json:"totalprice"`
	IsPaid      bool    `json:"paid"`
	Fragile     bool    `json:"fragile,omitempty"`
	OrderDetail []item  `json:"items"`
}

// customer struct
type customer struct {
	UserName      string  `json:"username"`
	Password      string  `json:"-"`
	Token         string  `json:"-"`
	ShipTo        address `json:"shipto"`
	PurchaseOrder order   `json:"order"`
}

// calculate total amount
func (c *customer) Total() {
	var price float64
	for _, item := range c.PurchaseOrder.OrderDetail {
		price += item.Price * float64(item.Quantity)
	}
	c.PurchaseOrder.TotalPrice = price
}

func main() {
	data := []byte(`
	{
		"username": "TaskMasterErnest",
		"shipto": {
			"street": "Mango Avenue",
			"city": "Mango City",
			"state": "Philadelphia",
			"zipcode": "PH-234"
		},
		"order": {
			"paid": false,
			"items": [
				{
					"itemname": "Math Rock",
					"desc": "A vinyl of complex rock songs",
					"price": 19.99,
					"quantity": 1
				}]
		}
	}
	`)

	// check JSON validity
	if !json.Valid(data) {
		fmt.Println("JSON is not valid:", data)
		os.Exit(1)
	}

	var c customer

	err := json.Unmarshal(data, &c)
	if err != nil {
		fmt.Println(err)
	}

	// adding two new items
	puppet := item{}
	puppet.Name = "Woody"
	puppet.Description = "A wooden cowboy doll from the Toy Story franchise"
	puppet.Price = 49.5
	puppet.Quantity = 2

	pc := item{}
	pc.Name = "Apple MacBook M3 Pro"
	pc.Price = 4199.00
	pc.Quantity = 1

	// add the two new items to the customer orders
	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, puppet)
	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, pc)

	// calculate total
	c.Total()

	// assert whether payment has been made
	c.PurchaseOrder.Fragile = true
	c.PurchaseOrder.IsPaid = true

	// marshal the customer order back to JSON
	customerOrder, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(customerOrder))
}
