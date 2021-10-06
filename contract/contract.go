package contract

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Inventory struct {
	Inventory_id int
	Name         string
	Type         string
	Tags         []string
	Purchased_at int
	Placement    Placement
}

type Placement struct {
	Room_id int
	Name    string
}

func FindByPurchase(dat []Inventory) {
	start, _ := time.Parse(time.RFC822, "16 Jan 20 00:00 WIB")
	end, _ := time.Parse(time.RFC822, "16 Jan 20 23:59 WIB")
	fmt.Println("All items were purchased on 16 Januari 2020.")
	for _, v := range dat {
		tm := time.Unix(int64(v.Purchased_at), 0)
		if tm.After(start) && tm.Before(end) {
			r, _ := json.Marshal(v)
			fmt.Println(string(r))
		}
	}
}

func FindByMeetRoom(dat []Inventory) {
	fmt.Println("All in the Meeting Room.")
	for _, v := range dat {
		if v.Placement.Name == "Meeting Room" {
			r, _ := json.Marshal(v)
			fmt.Println(string(r))
		}
	}
}

func FindByType(dat []Inventory, inv_type string) {
	fmt.Println("All Electronic Devices.")
	for _, v := range dat {
		if v.Type == inv_type {
			r, _ := json.Marshal(v)
			fmt.Println(string(r))
		}
	}
}

func FindByBrown(dat []Inventory) {
	fmt.Println("All Electronic Devices.")
	for _, v := range dat {
		if find_brown(v.Name) == "Brown" {
			r, _ := json.Marshal(v)
			fmt.Println(string(r))
		}
	}
}

func find_brown(val string) string {
	res := strings.Split(val, " ")
	for _, v := range res {
		if v == "Brown" {
			return v
		}
	}
	return "Not Found"
}
