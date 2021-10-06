package main

import (
	"encoding/json"
	"fmt"

	"github.com/mochganjarn/json-manipulation/contract"
)

func main() {

	byt := []byte(json_data())

	var dat []contract.Inventory
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	menu(dat)
}

func menu(dat []contract.Inventory) {
	var i int
	for i != 6 {
		fmt.Println("Menu")
		fmt.Println("1. Find items in the Meeting Room.")
		fmt.Println("2. Find all electronic devices.")
		fmt.Println("3. Find all the furniture.")
		fmt.Println("4. Find all items were purchased on 16 Januari 2020.")
		fmt.Println("5. Find all items with brown color.")
		fmt.Println("6. Exit")

		fmt.Print("Enter Number: ")
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			contract.FindByMeetRoom(dat)
			fmt.Print("Enter to continue...")
			fmt.Scanf("%d", &i)
		case 2:
			contract.FindByType(dat, "electronic")
			fmt.Print("Enter to continue...")
			fmt.Scanf("%d", &i)
		case 3:
			contract.FindByType(dat, "furniture")
			fmt.Print("Enter to continue...")
			fmt.Scanf("%d", &i)
		case 4:
			contract.FindByPurchase(dat)
			fmt.Print("Enter to continue...")
			fmt.Scanf("%d", &i)
		case 5:
			contract.FindByBrown(dat)
			fmt.Print("Enter to continue...")
			fmt.Scanf("%d", &i)
		}
	}
}

func json_data() (data string) {
	data = `[
		{
		  "inventory_id": 9382,
		  "name": "Brown Chair",
		  "type": "furniture",
		  "tags": [
			"chair",
			"furniture",
			"brown"
		  ],
		  "purchased_at": 1579190471,
		  "placement": {
			"room_id": 3,
			"name": "Meeting Room"
		  }
		},
		{
		  "inventory_id": 9380,
		  "name": "Big Desk",
		  "type": "furniture",
		  "tags": [
			"desk",
			"furniture",
			"brown"
		  ],
		  "purchased_at": 1579190642,
		  "placement": {
			"room_id": 3,
			"name": "Meeting Room"
		  }
		},
		{
		  "inventory_id": 2932,
		  "name": "LG Monitor 50 inch",
		  "type": "electronic",
		  "tags": [
			"monitor"
		  ],
		  "purchased_at": 1579017842,
		  "placement": {
			"room_id": 3,
			"name": "Lavender"
		  }
		},
		{
		  "inventory_id": 232,
		  "name": "Sharp Pendingin Ruangan 2PK",
		  "type": "electronic",
		  "tags": [
			"ac"
		  ],
		  "purchased_at": 1578931442,
		  "placement": {
			"room_id": 5,
			"name": "Dhanapala"
		  }
		},
		{
		  "inventory_id": 9382,
		  "name": "Alat Makan",
		  "type": "tableware",
		  "tags": [
			"spoon",
			"fork",
			"tableware"
		  ],
		  "purchased_at": 1578672242,
		  "placement": {
			"room_id": 10,
			"name": "Rajawali"
		  }
		}
	  ]`
	return data
}
