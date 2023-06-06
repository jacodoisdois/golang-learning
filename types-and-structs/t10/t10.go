package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := product{1, "Smarthphone", 5000.99, []string{"PROMO", "Eletronic"}}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))
	fmt.Println(p1Json)

	var p2 product
	jsonString := `{"id":2, "name":"Pencil", "price": 92.3, "tags": ["School", "Write"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)
}
